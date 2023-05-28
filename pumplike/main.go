package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"bytes"
	"io/ioutil"
	"strings"
	"encoding/json"
	"net/http"
)

var botbas = `
    ____  ____  __________  ___   ______  __
   / __ )/ __ \/_  __/ __ )/   | / ___/ |/ /
  / __  / / / / / / / __  / /| | \__ \|   / 
 / /_/ / /_/ / / / / /_/ / ___ |___/ /   |  
/_____/\____/ /_/ /_____/_/  |_/____/_/|_|  

        https://github.com/botbasx 
`
var welcome = `
───▄▀▀▀▄▄▄▄▄▄▄▀▀▀▄───
───█▒▒░░░░░░░░░▒▒█───
────█░░█░░░░░█░░█────
─▄▄──█░░░▀█▀░░░█──▄▄─
█░░█─▀▄░░░░░░░▄▀─█░░█
█▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█
█░░╦─╦╔╗╦─╔╗╔╗╔╦╗╔╗░░█
█░░║║║╠─║─║─║║║║║╠─░░█
█░░╚╩╝╚╝╚╝╚╝╚╝╩─╩╚╝░░█
█▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄█
`
func main() {
    fmt.Println(welcome)
	fmt.Println("ยินดีต้อนรับสู่โปรแกรมปั้มไลค์!")
	fmt.Println("ฟังชั่นที่มีตอนนี้")
	fmt.Println("1.เช็คเครดิต")
	fmt.Println("2.ปั้มไลค์")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("ใส่ตัวเลือก: ")
	scanner.Scan()
	clearConsole()
	fmt.Println(botbas)
	choice := scanner.Text()
	if choice == "1" {
		creditPump()
	} else if choice == "2" {
		fmt.Print("ใส่ลิงก์: ")
		scanner.Scan()
		link := scanner.Text()
		fmt.Print("เลือกอีโมชั่น\n1 = 👍 ไลค์\n2 = ❤️ หัวใจ\n3 = 😮 ว้าว\n4 = 😆 ขำ\n7 = 😥 เศร้า\n8 = 😡 โกรธ\nเลือกเลขที่ต้องการ:>>")
		scanner.Scan()
		idValue := scanner.Text()
		fmt.Print("เลือกจำนวน: ")
		scanner.Scan()
		amount := scanner.Text()
		makePump(link, idValue, amount)
	} else {
		fmt.Println("พบข้อผิดพลาด....")
	}
}
func creditPump() {
	file, err := ioutil.ReadFile("api.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	apiKey := strings.TrimSpace(string(file))
	url := "https://v2.sc24shop.store/api/v2"
	requestBody := map[string]interface{}{
		"key":    apiKey,
		"action": "balance",
	}
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	response, err := http.Post(url, "application/json", strings.NewReader(string(requestJSON)))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	responseData := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.StatusCode == 200 && responseData["balance"] != nil {
		balance := responseData["balance"]
		currency := responseData["currency"]
		message := fmt.Sprintf("เครดิตคงเหลือ: %v %v", balance, currency)
		fmt.Println(message)
	}
}
func makePump(link string, idValue string, amount string) interface{} {
	apiKey, err := ioutil.ReadFile("api.txt")
	if err != nil {
		fmt.Println("Error reading API key file:", err)
		return nil
	}
	apiKeyStr := string(apiKey)
	url := "https://v2.sc24shop.store/api/v3/emojithai/normal"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	requestBody := map[string]interface{}{
		"key":    apiKeyStr,
		"link":   link,
		"id":     idValue,
		"amount": amount,
	}
	fmt.Printf("ลิงก์: %s\nอีโมชั่นID: %s\nจำนวน: %s\n", link, idValue, amount)
	jsonStr, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		return nil
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return nil
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return nil
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}
	if resp.StatusCode == 200 {
		var responseData map[string]interface{}
		err := json.Unmarshal(respBody, &responseData)
		if err != nil {
			fmt.Println("Error parsing response JSON:", err)
			return nil
		}
		if success, ok := responseData["success"].(bool); ok && success {
			return responseData
		} else if message, ok := responseData["message"].(string); ok {
			fmt.Printf("สถานะคำสั่งซื้อ: %s\n", message)
		} else {
			fmt.Println("ไม่สามารถรับข้อมูลการสั่งซื้อได้")
		}
	} else {
		fmt.Printf("คำขอผิดพลาดสถานะโค๊ด: %d\n", resp.StatusCode)
	}
	return nil
}
func clearConsole() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
