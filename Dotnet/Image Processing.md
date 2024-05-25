# :point_down: Image Processing

## :mag: การใช้งาน

### Path Method
```c#
var fullPath = "https://localhost:8080/images/simple.jpg";
```

```c#
// "https://localhost:8080/images"
Path.Combine(environment.WebRootPath, "images");

// "https://localhost:8080/images/simple.jpg"
Path.Combine(environment.WebRootPath, "images", "simple.jpg");

// ".jpg"
Path.GetExtension(fullPath);

// "simple.jpg"
Path.GetFileName(fullPath);

// "simple"
Path.GetFileNameWithoutExtension(fullPath);

// "https://localhost:8080/images/simple.gif"
Path.ChangeExtension(fullPath, "gif");

// "fcijve0j.png"
// แล้วเพิ่ม file type ภายหลัง
Path.ChangeExtension(Path.GetRandomFileName(), "png");
```
```c#
HttpUtility.UrlEncode("https://mywebsite.com/api/get me this file.jpg")
//Output: "https%3a%2f%2fmywebsite.com%2fapi%2fget+me+this+file.jpg"

Uri.EscapeUriString("https://mywebsite.com/api/get me this file.jpg");
//Output: "https://mywebsite.com/api/get%20me%20this%20file.jpg"
Uri.EscapeDataString("https://mywebsite.com/api/get me this file.jpg");
//Output: "https%3A%2F%2Fmywebsite.com%2Fapi%2Fget%20me%20this%20file.jpg"

//When your url has a query string:
Uri.EscapeUriString("https://mywebsite.com/api/get?id=123&name=get me this file.jpg");
//Output: "https://mywebsite.com/api/get?id=123&name=get%20me%20this%20file.jpg"
Uri.EscapeDataString("https://mywebsite.com/api/get?id=123&name=get me this file.jpg");    
//Output: "https%3A%2F%2Fmywebsite.com%2Fapi%2Fget%3Fid%3D123%26name%3Dget%20me%20this%20file.jpg"
```

### :bulb: Upload => Normal Image

```c#
public IActionResult Index(IFormFile image) 
{
    var fullPath = Path.Combine(environment.WebRootPath, "images", image.FileName);

    using (var stream = new FileStream(fullPath, FileMode.Create))
    {
        image.CopyTo(stream);
    }
}
```

### :bulb: Upload => Webp
```c#
using System.Drawing.Common;
using ImageProcessor;
using ImageProcessor.Plugins.WebP;
```

```c#
public IActionResult Index(IFormFile image) 
{
   var filenameWithExtension = Path.ChangeExtension(image.FileName, "webp"); 
   var fullPath = Path.Combine(environment.WebRootPath, "images", filenameWithExtension);

    using (var stream = new FileStream(fullPath, FileMode.Create))
    {
        using (var imageFactory = new ImageFactory(preserveExifData: false))
        {
            imageFactory.Load(image.OpenReadStream())
                        .Format(new WebPFormat())
                        .Quality(80)
                        .Save(stream);
        }
    }
}
```

### :bulb: Upload => QR Code
```c#
using QRCoder;
```

```c#
var url = "https://google.co.th";

using (var qrGenerator = new QRCodeGenerator())
{
    using (var qrData = qrGenerator.CreateQrCode(url, QRCodeGenerator.ECCLevel.Q))
    {
        using (var qrBytes = new BitmapByteQRCode(qrData))
        {
            var content = qrBytes.GetGraphic(20);

            //#region write file
            var wwwroot = Path.Combine(environment.WebRootPath, "images");
            var fileName = Guid.NewGuid().ToString();
            var fileExtenstion = ".jpg";
            var fullFullPath = Path.Combine(wwwroot, fileName, fileExtenstion);

            if (!Directory.Exists(wwwroot))
            {
                Directory.CreateDirectory(wwwroot);
            }

            // วิธีที่ 1 
            File.WriteAllBytes(fullFullPath, content);

            // วิธีที่ 2 
            using (var stream = new FileStream(fullFullPath, FileMode.Create))
            {
                stream.Write(content);
            }
            //#endregion
        }
    }
}
```

### :bulb: Delete File
```c#
var fullPath = Path.Combine(environment.WebRootPath, "images", "simple.jpg");

if (File.Exists(fullPath))
{
    File.Delete(fullPath);
}
```
