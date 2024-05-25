# :point_down: Mat Input

รวบรวมปัญหาของ Mat-Input

## :mag: ปัญหา

### Type Number ทำงานไม่ถูกต้อง
โดยการทำงานจะไม่ได้ตัวเลขาเสมอไป

```ts
numericOnly(event): boolean {
    // noinspection JSDeprecatedSymbols
    const charCode = (event.which) ? event.which : event.key || event.keyCode;  
    // keyCode is deprecated but needed for some browsers
    return !(charCode === 101 || charCode === 69 || charCode === 45 || charCode === 43);
}
```

```html
 <input matInput type="number" (keypress)="numericOnly($event)" />
```


