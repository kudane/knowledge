# :point_down: Angular Inner HTML

Pipe สำหรับ append html ในหน้าเว็บและมีความปลอดภัย

## :mag: การใช้งาน

### Custom Render HTML
```ts
import { Directive, ElementRef, Input } from '@angular/core';
// npm install dompurify --save
import * as DOMPurify from 'dompurify';

@Directive({
    selector: '[innerhtml]',
})
export class InnerHtmlDirective {
    // https://github.com/cure53/DOMPurify#can-i-configure-dompurify
    private readonly CONFIG = {
        ADD_ATTR: [],
        FORBID_ATTR: []
    };

    constructor(
        private el: ElementRef<HTMLElement>,
        private renderer: Renderer2,
    ) { }

    @Input() set innerhtml(html: string) {
        if (!html) throw new Error('[safehtml] error, html is invalid');
        const cleanHtml = DOMPurify.sanitize(html, this.CONFIG);
        this.renderer.setProperty(this.el.nativeElement, 'innerHTML', cleanHtml)
    }
}
```

### พื้นฐาน
```ts
import { Pipe, PipeTransform } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';
// npm install dompurify --save
import * as DOMPurify from 'dompurify';

@Pipe({
  name: 'innerhtml'
})
export class InnerHtmlPipe implements PipeTransform {
  constructor(private sanitizer: DomSanitizer) { }

  // ได้มายังไงแสดงตามนั้น
  transform1(html: string) {
    return this.sanitizer.bypassSecurityTrustHtml(html);
  }

  // remove script & event
  transform2(html: string) {
    return DOMPurify.sanitize(html);
  }

  // remove script & event
  transform3(html: string) {
    return sanitizer.sanitize(SecurityContext.HTML, html);
  }
}
```
