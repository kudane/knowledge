กำหนด environment ของ [NestJS](https://docs.nestjs.com/techniques/configuration)

Step 1: กำหนด NODE_ENV ใน script

`package.json`
```
{
    ...
    script: {
        "start:{dev, production, more...}": "set NODE_ENV={dev, production, more...} && nest ...."
    }
}
```

<br>

Step 2: สร้างไฟล์ .env.{dev, production, more...} ใน folder ที่ต้องการ

for example
```
src
|_environment
    |_.env.dev
    |_.env.production

```

<br>

Step 3: ติดตั้ง @nestjs/config
```
npm i --save @nestjs/config

```

<br>

Step 4: กำหนดการใช้ .env

`app.module.ts`
```ts
import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';

@Module({
  imports: [ConfigModule.forRoot(
    {
        // กำหนดใช้อยู่ใน root provider ใช้งานได้ทุกที่ไม่ต้อง import ใหม่
        isGlobal: true,

        // เลือกใช้ไฟล์ .env ตามคำสั่ง script ใน Step 1
        envFilePath: "./src/environment/.env.${process.env.NODE_ENV.trim()}",

        // Optional: หากต้องทำ custom จาก env เป็น object ให้เรียกใช้งานง่ายๆ ก็ทำ custom ที่ load ได้
        load: [custom1, custom2, ...]
    }
  )],
})
export class AppModule {}
```

<br>

Step 5: การนำไปใช้งาน

`db.service.ts`
```ts
// DI ผ่าน constructor ใน class ที่จะใช้งาน
constructor(private configService: ConfigService) {}

// เรียกใช้คำสั่ง `this.configService.get` และใส่ key ของ env เป็น parameter
const dbUser = this.configService.get<string>('DATABASE');
```
