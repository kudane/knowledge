# :point_down: Request Exception

Request Exception ในที่นี้คือการ Custom Exception ในการ handle ปัญหาต่างๆที่เกิดขึ้นเมื่อมีการขอ request เข้ามา ซึ่งการทำ Custom Exception นั้น
จะช่วยให้สามารถกำหนด Exception ไว้เพียงจุดเดียว เพื่อใช้ประยุกต์ในการทำอะไรบางอย่าง

## :mag: การใช้งาน

### :bulb: Create CustomException.cs
```c#
public class CustomException : Exception
{
    public object Errors { get; set; }
    public HttpStatusCode Code { get; set; } = HttpStatusCode.BadRequest;

    public CustomException(object errors = null)
    {
        Errors = errors;
    }

    public CustomException(HttpStatusCode code, object errors = null)
    {
        if (code == HttpStatusCode.OK) 
        {
            Code = HttpStatusCode.BadRequest;
        }
        Code = code;
        Errors = errors ?? "Error Occurred";
    }
}
```

### :bulb: Create GlobalExceptionHandle.cs
```c#
public class GlobalExceptionHandle : IExceptionFilter
{
    public void OnException(ExceptionContext context)
    {
        switch (context.Exception)
        {
            // custom exception ที่กำหนดไว้
            case CustomException exception:
                // กำหนด code ที่ต้องการ
                context.HttpContext.Response.StatusCode = (int)exception.Code;
                // กำหนด result ใหม่ที่ต้องการ
                context.Result = new JsonResult(exception.Errors);
                break;

            // ...

            // system exception ที่นอกเหนือการ custom 
            case Exception exception:
                // กำหนด code ที่ต้องการ
                context.HttpContext.Response.StatusCode = (int)HttpStatusCode.InternalServerError;
                // กำหนด result ใหม่ที่ต้องการ
                context.Result = new JsonResult(exception.Message);
                break;
        }
    }
}
```

### :bulb: เพิ่มใน pipeline ในที่นี้คือ Filters
```c#
 public void ConfigureServices(IServiceCollection services)
{
    services.AddControllers(options =>
    {
        options.Filters.Add<GlobalExceptionHandle>();
    });
    
    // ...
}
```

### :bulb: การเรียกใช้งาน
```c#
if (username.invalid)
{
    throw new CustomException("Username is invalid ....");
}
```


