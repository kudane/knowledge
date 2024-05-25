# :point_down: String Interpolation 

ใช้ดำเนินการระหว่าง ***ตัวแปร*** และ ***ข้อความ*** เช่น ต้องการ replace ข้อความในตำแหน่งที่ต้องการ หรือ การต่อข้อความด้วยตัวแปรใดๆ
[อ่านเพิ่มคลิก](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/tokens/interpolated)

## :mag: การใช้งาน

### string.format method

เหมาะกับงาน replace ข้อความในตำแหน่งที่ต้องการ

```c#
var name = "John";
var lname = "Walker";
var message = string.Format("Hello {0} {1}", name, lname);

// "Hello John Walker"
Console.Write(message);
```

### เครื่องหมาย "$" 

"$" หากนำมาใส่ก่อนเปิด " จะสามารถ replace ข้อความด้วยตัวแปรได้ง่ายขึ้น โดยต้องเปิด {ชื่อตัวแปร} 

```c#
var name = "John";
var message = $"Hello {name}";

// "Hello John"
Console.Write(message);
```

### เครื่องหมาย "@"

"@" ช่วยให้แสดงข้อความได้หลายบรรทัด ไม่ต้องติดกันยาวๆ อีกแล้ว

```c#
var message = @"
	Hello Everyone
	How are you today?
";

/* 
	Hello Everyone
	How are you today?
*/
Console.Write(message);

// html text
var html = $@"
	<div class='mail-template'>
		<h1>Hello Everyone</h1>
		<span>How are you today?<span>
	</div>
";
```

### เครื่องหมาย "$" ร่วมกับ "@"

"$@" ช่วยให้แสดงข้อความได้หลายบรรทัดและ replace ข้อความด้วยตัวแปร พร้อมๆกัน

```c#
var name = "John";
var message = $@"
	Hello {name}
	How are you today?
";

/* 
	Hello John
	How are you today?
*/
Console.Write(message);
```