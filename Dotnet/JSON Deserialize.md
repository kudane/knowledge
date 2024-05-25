# :point_down: JSON Deserialize

Deserialize คือ การแกะ Json String ให้อยู่ในรูปแบบของ Object ได้หลายประเภท เช่น Model, List, Arrray
ซึ่งข้อดีของการใช้ Deserialize คือสามารถ Deserialize ลงลึกลงไปได้ลึกเลยที่เดียวนั้นเอง

## อะไรคือ Class/Method/Field Property Attribute ?

Class/Field Attribute คือ การเพิ่มความสามารถรูปแบบหนึ่ง โดยจะกำหนดไว้บน Field หรือ Method หรือ Class ได้

```c#
[ApiController]
// กำหนด Class HttpRequest เป็น REST.API
public class HttpRequest
{
    [HttpGet]
    // กำหนด method Get() เป็น HTTP.GET
    public void Get()
    {
        // get action
    }
}
```

## :mag: การใช้งาน

### :bulb: System.Text.Json

#### ***[JsonPropertyName]***

คือ Field Attribute มีหน้าที่กำหนดว่าต้องการ Mapping Field จาก Json String Field ใดนั้นเอง
โดยต้องกำหนดทุก Field 

#### ***[JsonIgnore]***

คือ Class Attribute กำหนดว่าไม่ต้องการ Mapping Field นี้จาก Json String

```json
[
  {
    "id": 1,
    "fullname": "John Walker",
    "gender": "Male"
},
  {
    "id": 2,
    "fullname": "Peter Jackson",
    "gender": "Male"
  }
]
```

```c#
// import
using System.Text.Json;
using System.Text.Json.Serialization;

// model
public class Customer
{
    [JsonPropertyName("ชื่อ_Field_Json_ที่ต้องการ")]
    public int Id { get; set; }
    [JsonPropertyName("fullname")]
    public string FullName { get; set; }
    [JsonPropertyName("gender")]
    public string Gender { get; set; }
}

// code
// เนื่องจากตัวอย่าง Json String เป็น Array จึง mapping ด้วย IEnumerable/List ให้ประเภทตรงกัน
JsonSerializer.Deserialize<IEnumerable<Customer>>("ตัวอย่าง Json จากด้านบนเด้อ");
```

```c#
// การ setting format ของ datetime ในjson กรณี local timezone คลาดเคลื่อน ใน System.Text.Json
services
  .AddControllers()
  .AddJsonOptions(options =>
  {
    options.JsonSerializerOptions.Converters.Add(new DateTimeConverter());
  });

public class DateTimeConverter : JsonConverter<DateTime>
{
  public override DateTime Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
  {
    return DateTime.Parse(reader.GetString());
  }
  public override void Write(Utf8JsonWriter writer, DateTime value, JsonSerializerOptions options)
  {
    writer.WriteStringValue(value.ToUniversalTime().AddHours(7).ToString("yyyy'-'MM'-'dd'T'HH':'mm':'ssZ"));
  }
}
```

### :bulb: Newtonsoft.Json

Newtonsoft Json มีหลายความสามารถที่ใช้ได้ง่ายกว่าตัวด้านบน แต่โดยหลักๆแล้ว ไม่ต่างกันมากมาย

#### ***[JsonProperty]***

คือ Field Attribute กำหนดว่าต้องการ Mapping Field จาก Json String Field ใดนั้นเอง โดยมีเงื่อนไขดังนี้
1. ไม่ต้องกำหนด หากชื่อ Field ตรงกัน
2. ต้องกำหนดหากชื่อ Field มีสัญลักษณ์ "_" หรือ "-" และอื่นๆ คั่นกลาง

#### ***[JsonIgnore] เหมือนด้านบน***

```json
[
  {
    "_id": 1,
    "fullname": "John Walker",
    "gender": "Male"
},
  {
    "_id": 2,
    "fullname": "Peter Jackson",
    "gender": "Male"
  }
]
```

```c#
// import
using Newtonsoft.Json;

// model
public class Customer
{
    [JsonProperty("_id")]
    public int Id { get; set; }

    public string FullName { get; set; }
    public string Gender { get; set; }
}

// code
// เนื่องจากตัวอย่าง Json String เป็น Array จึง mapping ด้วย IEnumerable/List ให้ประเภทตรงกัน
JsonConvert.DeserializeObject<IEnumerable<Customer>>("ตัวอย่าง Json จากด้านบนเด้อ");
```

```c#
// การ setting format ของ datetime ในjson กรณี local timezone คลาดเคลื่อน ใน Newtonsoft.Json
services.AddControllers().AddNewtonsoftJson(options =>
{
  options.SerializerSettings.DateTimeZoneHandling = Newtonsoft.Json.DateTimeZoneHandling.Local;
  options.SerializerSettings.ReferenceLoopHandling = Newtonsoft.Json.ReferenceLoopHandling.Ignore;
});

```
