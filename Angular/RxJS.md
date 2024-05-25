# :point_down: RxJS

Angular มีการ Buit-In (RxJS) เข้ามาในโปรเจคอยู่แล้ว โดยความหมายของ RxJS คือ **Reactive Extension Library** และในเชิงเทคนิคการนำไปใช้ คือ การจัดการข้อมูล Asynchronous ซึ่งข้อมูลในที่นี้ต้องเป็น **Obervable** หรือ **Subject ซึ่งใน Angular มีอยู่ 3 ตัวคือ Subject, BehaviorSubject และ ReplaySubject** (FYI)

## :mag: ความรู้พื้นฐาน
การทำงานของ **Obervable หรือ Subject** ให้มองเป็นการทำงานจาก **จุด A ---ไปยัง----> B** <br>
A คือ Observable ตัวแปรที่สามารถทำการ subscribe ได้นั้นเอง <br>
ฺB คือ Observer ฟังก์ชั่นที่เขียนภายใน scope ของ subscribe ที่จะทำงานเมื่อ A ข้อมูลเปลี่ยน <br>

```ts
let initValue = [1, 2, 3, 4, 5];

of(initValue)                   // <------ จุด A
.subscribe(
    val => console.log(val)     // <------ จุด B
);
```

## :mag: Subject ประเภทต่างๆ

### Subject
คือ Obervable ประเภทพื้นฐาน โดยการ new ขึ้นมาไม่จำเป็นต้องกำหนดข้อมูลเริ่มต้น
เมื่อ Observer ทำการ subscribe จะยังไม่ได้ข้อมูลในครั้งแรก
โดย Observer จะได้รับข้อมูลใหม่เมื่อมี assign ข้อมูลใหม่ให้กับ Obervable

```ts
const subject = new Subject<number>();  
```

### BehaviorSubject
คือ Obervable ที่ต้องกำหนดข้อมูลเริ่มต้นเสมอ
เมื่อ Observer ทำการ subscribe จะได้รับข้อมูลเริ่มต้นทันที
และ Observer จะได้รับข้อมูลใหม่อีกครั้งเมื่อมี assign ข้อมูลใหม่ให้กับ Obervable

```ts
const subject = new BehaviorSubject<bool>(true);  
```

### ReplaySubject
คือ Obervable ที่ไม่ต้องกำหนดข้อมูลเริ่มต้น แต่ต้องกำหนดขนาดของข้อมูลที่จะ backup เพื่อทำ replay ไว้
การ replay ในที่นี้ก็คือ หการส่งค่าก่อนหน้าออกมาก่อนจนสุดท้ายปล่อยข้อมูลล่าสุดออกมา ซึ่งการจะปล่อยข้อมูลย้อนหลังจะเท่ากับจำนวณครั้งที่ต้องการ backup
เมื่อ Observer ทำการ subscribe จะไม่ได้รับข้อมูลเริ่มต้นทันที
และ Observer จะได้รับข้อมูลใหม่อีกครั้งเมื่อมี assign ข้อมูลใหม่ให้กับ Obervable

**หากมี Observer ใหม่ที่พึ่งมา subscribe ที่หลัง ก็จะได้ข้อมูลย้อนหลังตามจำนวณที่ backup ไว้จนได้ข้อมูลล่าสุด**

```ts
const subject = new ReplaySubject(3); // backup size คือ 3
```



## :mag: RxJS Operators
การใช้งาน Operators จะอยู่ระหว่าง **จุด A ---[Operators]----> B**
สามารถใช้งานภายในคำสั่ง pipe(.....) สามารถใช้กี่ Operators ก็ได้ เช่น pipe(mapTo(1),mapTo(2))

### Creational Operators
คือ ฟังก์ชั่นในการสร้าง Obervable ออกมานั้นเอง ซึ่งสามารถกำหนดค่าอะไรใส่ไปก็

```ts
interval        คือ  Obervable จะปล่อยออกตามเวลาที่กำหนดไว้    
of              คือ  Obervable ที่กำหนดข้อมูลได้ เช่น of(1), of({a: 1, b: 'b'})   
from            คือ  Obervable ที่กำหนด dataset ได้ เช่น from([1,2,3]) เมื่อ subscribe จะได้รับข้อมูลมาที่ละ set คือ 1, 2 และ 3 เป็นอันสุดท้าย
ajax            คือ  Obervable ที่ได้จาก response ของ request  
empty           คือ  Obervable ที่ไม่กำหนดอะไรและไม่เกิดการ value change

// มีอีกมากอ่านได้ที่นี้ https://rxjs.dev/guide/operators#creation-operators-1
```

### Transformation Operators
คือ ฟังก์ชั่นในการ mapping ข้อมูลนั้น ซึ่งจะสามารถมองเห็นข้อมูลและแก้ไขตามที่ต้องการได้

```ts
map             คือ  ฟังก์ชั่น ที่จะกำหนดข้อมูลใหม่ได้และต้อง return ออกไปเป็น Obervable เหมือนเดิม
mapTo           คือ  ฟังก์ชั่น ที่กำหนดข้อมูลเลย โดยใส่ข้อมูลที่ต้องการไป 
switchMap       คือ  ฟังก์ชั่น ที่ใช้ในการสลับไปทำการ Observer ใหม่ โดยจะ unsubscribe ตัวเก่าก่อน

// มีอีกมากอ่านได้ที่นี้ https://rxjs.dev/guide/operators#transformation-operators
```

### Filtering Operators
คือ ฟังก์ชั่นในการ filter ข้อมูลเลย ไม่ต่างกับ map, fliter, redue เลย

```ts
take            คือ  ฟังก์ชั่น ที่จะเลือกเอามาตามจำนวนที่ต้องการ เมื่อได้ข้อมูลครบจะ unsubscribe อัตโนมัติ take(1)
last            คือ  ฟังก์ชั่น ที่จะเลือกเอามาเฉพาะข้อมูลล่าสุด last()
skip            คือ  ฟังก์ชั่น ที่จะข้ามตามจำนวนที่กำหนด skip(3)
debounceTime    คือ  ฟังก์ชั่น ที่ เมื่อได้ข้อมูลมาจะทำการหน่วงเวลาตามที่กำหนดไว้และ เมื่อครบเวลาจะปล่อยข้อมูลล่าสุดออกไป และเริ่มหน่วงใหม่อีกครั้ง

// มีอีกมากอ่านได้ที่นี้ https://rxjs.dev/guide/operators#filtering-operators
```

### Utility Operators

```ts
tap             คือ  ฟังก์ชั่น ที่เอาไว้ส่องข้อมูลของ Observable
delay           คือ  ฟังก์ชั่น ที่จะบังคับให้ Observer หยุดรอเพื่อทำอะไรบางอย่าง และปล่อยข้อมูลออกไปเหมือนเดิม 

// มีอีกมากอ่านได้ที่นี้ https://rxjs.dev/guide/operators#utility-operators
```


> Operators ของ RxJS นั้นมีเยอะและหลากหลายมาก นำไปประยุกต์ใช้ได้หลายแบบมาก เท่าที่คุณคิดออกเลย
> สามารถอ่านทำความเข้าใจได้จากเว็บเหล่านี้

> https://rxjs.dev/guide/overview, https://indepth.dev/reference/rxjs
