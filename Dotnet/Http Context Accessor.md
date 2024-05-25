# :point_down: Http Context Accessor

คือ เครื่องมือที่ช่วยทำให้เข้าถึง Http Context จาก request ได้โดยตรง ซึ่งปกติจะเข้าถึงได้โดยตรงที่ controller เท่านั้น ทำให้เกิดข้อจำกัดในการที่จะอ่านอะไรบางอย่าง เช่น identity ว่า login อยู่รึเปล่า หรืออ่าน claim values ได้ยากขึ้น

## :mag: การใช้งาน

### :bulb: Create HttpContextAccessor ที่ต้องการในที่นี้ข้อสมมุติ  CurrentAccessor คือ user ที่ login อยู่

####  ICurrentUserAccessor.cs
```c#
public interface ICurrentUserAccessor
{
    bool IsCurrentUserAsAdmin();
    Guid GetCurrentUserId();
}
```

####  CurrentUserAccessor.cs
```c#
public class CurrentUserAccessor : ICurrentUserAccessor
{
    private readonly IHttpContextAccessor httpContextAccessor;

    public CurrentUserAccessor(IHttpContextAccessor httpContextAccessor)
    {
        this.httpContextAccessor = httpContextAccessor;
    }

    public Guid GetCurrentUserId()
    {
        var userId = httpContextAccessor.HttpContext.User?.Claims?.FirstOrDefault(x => x.Type == "UserId")?.Value;
        return Guid.Parse(userId);
    }

    public bool IsCurrentUserAsAdmin()
    {
         var isAdmin = httpContextAccessor.HttpContext.User?.Claims?.FirstOrDefault(x => x.Type == "IsAdmin")?.Value;
         return isAdmin != null;
    }
}
```

### :bulb: Startup.cs

```c#
public void ConfigureServices(IServiceCollection services)
{
    // ...
    services.AddHttpContextAccessor();
    services.AddScoped<ICurrentUserAccessor,CurrentUserAccessor>();
}
```

### :bulb: เรียกใช้งานผ่าน DI เช่น

```c#
public TestController(ICurrentUserAccessor currentAccessor)
{
    // do something
}
```
