Business Core คือ logic ที่ระบบจะต้องมี ที่นี่ Business Core ของ clean คืออะไรละ

เอาแบบย่อๆ สั้นๆ

Step 1: กำหนด interface ของ Business Logic แต่ละตัวให้ก่อน เช่น
```c#
public interface IPingHandler
{
    void Handle();
} 
```

Step 2: กำหนด parameter หรือ output ที่จะมี โดยเติมเข้าไปใน interface
```c#
public class PingCommand
{
    public int ID {get; set;}
}

public class PingOutput
{
    public bool Success {get; set;}
}

public interface IPingHandler
{
    PingOutput Handle(PingCommand command);
} 
```

Step 3: สร้าง imprement ของ `IPingHandler` และกำหนด logic ของ business นี้ไปได้เลย
```c#
public class PingHandler : IPingHandler
{
    public PingOutput Handle(PingCommand command)
    {
        Console.WriteLine(command.ID);

        // do something logic

        return new PingOutput(){ Success: true };
    }
} 
```

Step 4: นำไปใช้ตาม framework ที่ใช้งาน อาจจะมีความต่างในด้าน syntax อยู่บ้าง ทุกภาษาย่อมมีจุดที่มีเหมือนกันเสมอๆ

>> have fun
