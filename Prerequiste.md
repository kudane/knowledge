prerequisite : system XXX
source code
-- frontend develop ด้วยอะไร (ถ้าเป็น mobile develop ด้วยอะไร,  เอาไว้ใช้สำหรับลง tools เตรียมพัฒนา)
-- backend develop ด้วยอะไร

config server
-- deploy ไว้ที่ไหนเช่น deploy ไว้บน aws, ใช้ service อะไรเช่น k8s
-- ip ของ server คืออะไร
-- deploy ด้วยอะไรเช่นถ้า deploy ด้วย container เก็บ container registry ไว้ที่ไหน
-- มีการเรียกใช้ tools อื่นๆเช่น firebase messaging ไหม config ของ firebase คืออะไร

Third party system ไปเชื่อมต่อ
-- ระบบที่ไปเชื่อมต่อคืออะไรเช่นระบบดูรอบรถ
-- ระบบเราคุยกับระบบนั้นด้วย protocall อะไรเช่น web service, api 
-- ระบบเรายืนยันตัวตนยังไงในการเข้าไปขอ data จาก Third party ยังไงเช่น api key, token

ข้อ concern
-- ระบบมีการใช้ library abcd ซึ่งไม่ได้มีการอัพเดทเป็น version ใหม่อาจจะมีช่องโหว่
-- ระบบมีการใช้ library ที่ vendor เป็นผู้พัฒนาขึ้นมาใช้เองหากต้องการ upgrade version ในอนาคตอาจจะไม่สามารถทำได้

credit: nanthapol.t