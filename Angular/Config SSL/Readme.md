# :point_down: Angular Config SSL

วิธีที่กำหนดให้ http:localhost:4200 เป็น https

> Credit: https://github.com/RubenVermeulen/generate-trusted-ssl-certificate

## :mag: การใช้งาน

### :bulb: Step 1: สร้าง Cert (มีอยู่แล้วข้ามไป) 

1. เปิด powershell รันไฟล์ generate.sh จะได้ไฟล์มี 2 ไฟล์คือ server.crt และ server.key
2. หรือ ใช้ไฟล์ในโพลเดอร์ก็ได้เช่นกัน

### :bulb: Step 2: ติดตั้ง Cert (มีอยู่แล้วข้ามไป) 

1. คลิกไฟล์ server.crt, กด install cert
2. เลือก local machine,
3. เลือก please all certificates in the following store
4. browse เลือก trusted root certification authorities 
4. เสร็จสิ้น

### :bulb: Step 3: นำไปใช้ใน angular project

1. สร้างโพลเดอร์ ssl
   ```
    -project
    |__src
    |__ssl  <----ตำแหน่งนี้
    |__...
    |__package.json
   ```
2. คัดลอกไฟล์ server.crt และ server.key วางในโพลเดอร์ ssl
3. serve ด้วยคำสั่ง
   ```
   ng serve --ssl true --ssl-cert ssl/server.crt --ssl-key ssl/server.key
   ```
