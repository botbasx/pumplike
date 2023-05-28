from pystyle import Colors, Colorate
import requests
import json
import os
botbas = """
    ____  ____  __________  ___   ______  __
   / __ )/ __ \/_  __/ __ )/   | / ___/ |/ /
  / __  / / / / / / / __  / /| | \__ \|   / 
 / /_/ / /_/ / / / / /_/ / ___ |___/ /   |  
/_____/\____/ /_/ /_____/_/  |_/____/_/|_|  

        https://github.com/botbasx 
"""
welcome = """
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

"""
def main():
    print(Colorate.Horizontal(Colors.rainbow, welcome))
    print(Colorate.Horizontal(Colors.rainbow, "ยินดีต้อนรับสู่โปรแกรมปั้มไลค์!"))
    print(Colorate.Horizontal(Colors.rainbow, "ฟังชั่นที่มีตอนนี้"))
    print(Colorate.Horizontal(Colors.rainbow, "1.เช็คเครดิต"))
    print(Colorate.Horizontal(Colors.rainbow, "2.ปั้มไลค์"))
    choice = input(Colorate.Horizontal(Colors.rainbow, "ใส่ตัวเลือก: "))
    os.system("clear")
    print(Colorate.Horizontal(Colors.rainbow, botbas))
    if choice == '1':
        creditpump()
    elif choice == '2':
        link = input(Colorate.Horizontal(Colors.rainbow, "ใส่ลิงก์: "))
        id_value = input(Colorate.Horizontal(Colors.rainbow, "เลือกอีโมชั่น\n1 = 👍 ไลค์\n2 = ❤️ หัวใจ\n3 = 😮 ว้าว\n4 = 😆 ขำ\n7 = 😥 เศร้า\n8 = 😡 โกรธ\nเลือกเลขที่ต้องการ:>>"))
        amount = input(Colorate.Horizontal(Colors.rainbow, "เลือกจำนวน: "))
        make_pump(link, id_value, amount)
    else:
        print(Colorate.Horizontal(Colors.rainbow, "พบข้อผิดพลาด...."))
def creditpump():
    with open("api.txt", "r") as f:
        api_key = f.read().strip()
    url = "https://v2.sc24shop.store/api/v2"
    request_body = {
        "key": api_key,
        "action": "balance"
    }
    response = requests.post(url, json=request_body)
    response_data = json.loads(response.text)
    if response.status_code == 200 and 'balance' in response_data:
        balance = response_data['balance']
        currency = response_data['currency']
        message = "เครดิตคงเหลือ: {} {}".format(balance, currency)
        print_rainbow(message)
    else:
        error_message = "Error occurred: {}".format(response_data)
        print_rainbow(error_message)
def make_pump(link, id_value, amount):
    with open("api.txt", "r") as f:
        api_key = f.read().strip()    
    url = "https://v2.sc24shop.store/api/v3/emojithai/normal"
    headers = {"Content-Type": "application/json"}
    request_body = {
        "key": api_key,
        "link": link,
        "id": id_value,
        "amount": amount
    }
    os.system("clear")
    print(Colorate.Horizontal(Colors.rainbow, botbas))
    print(Colorate.Horizontal(Colors.rainbow, "ลิงก์: {}\nอีโมชั่นID: {}\nจำนวน: {}".format(link,id_value,amount)))
    response = requests.post(url, headers=headers, data=json.dumps(request_body))
    if response.status_code == 200:
        response_data = response.json()
        if response_data.get('success'):
            return response_data
        else:
            print(Colorate.Horizontal(Colors.rainbow, "สถานะคำสั่งซื้อ: {}".format(response_data.get('message'))))
    else:
        print(Colorate.Horizontal(Colors.rainbow, "คำขอผิดพลาดสถานะโค๊ด: {}".format(response.status_code)))
    return None
def print_rainbow(message):
    colors = ['\033[91m', '\033[93m', '\033[92m', '\033[94m', '\033[95m', '\033[96m']
    reset_color = '\033[0m'
    rainbow_message = ''.join([f"{colors[i % len(colors)]}{char}" for i, char in enumerate(message)])
    print(rainbow_message + reset_color)
main()