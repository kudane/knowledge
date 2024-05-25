# :point_down: If Else

การเขียน If Else นั้นส่งผลต่อ flow การทำงานของโค้ดอย่างมาก หากไม่ได้กำหนด condition path ให้ดีแล้วละก็จะทำให้โค้ดมีความซับซ้อนมากขึ้นโดยไม่ได้ตั้งใจ
เพราะฉะนั้นมาขบคิดกันดีกว่าว่าจะใช้หลักอะไรได้บ้างในการกำหนด condtion path ให้น้อยลง

***Dont't***
> ห้ามใส่เงื่อนไขเยอะเกินไป, แยกได้แยก

> ห้ามใส่เงื่อนไขซ่อนในเงื่อนไข, ถ้าไม่จำเป็น

***Do***
> ควรใส่ 1 เงื่อนไขใน 1 condition path

> ใช้ Negative Pattern คือ เช็คในทิสทางตรงกันข้าม เพื่อตัดข้อผิดพลาดออกไป

> ใช้ Classify Case Pattern กำหนดกลุ่มของโค้ดในแต่ละเงื่อนไข

## :mag: การใช้งาน

### :bulb: Positive Pattern

คือ การเช็คด้วยเงื่อนไขที่ต้องการ

```c#
public User FindUser(string id)
{
    if(id != null)
    {
        var user = UserQuery(id);

        if(user != null)
        {
            return user;
        }
        else
        {
            thorw new Excetion("User Notfound");
        }
    }
    else
    {
        thorw new Excetion("Id Invalid");
    }
}
```

### :bulb: Negative Pattern

คือ การเช็คด้วยเคสที่ผิดพลาด

```c#
public User FindUser(string id)
{
    if(id == null)
    {
        thorw new Excetion("Id Invalid");
    }

    var user = UserQuery(id);

    if(user == null)
    {
        thorw new Excetion("User Notfound");
    }
    
    return user;
}
```

### :bulb: Classify Case Pattern

คือ การใช้ตัวแปรเพื่อกำหนดประเภทของเงื่อนไข

```c#
public User DoSomething()
{
    var user = UserQuery(1);
    var admin = (user.IsAdmin == true);
    var approver = (user.IsApprover == true);

    if(admin)
    {
        // do admin something...
    }

    if(approver)
    {
        // do approver something...
    }

    // ......
}
```

```c#
public User DoSomething()
{
    var user = UserQuery(1);
    var userCases = new()
    {
        admin = (user.IsAdmin == true),
        approver = (user.IsApprover == true)
    }

    if(userCases.admin)
    {
        // do admin something...
    }

    if(userCases.approver)
    {
        // do approver something...
    }

    // ......
}
```
