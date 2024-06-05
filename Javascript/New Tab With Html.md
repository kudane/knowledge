เปิด new tab ด้วย html [Ref](https://stackoverflow.com/questions/76223596/opening-a-computed-document-using-window-open-and-data-scheme)

```js
const html = '<html><body><p>Hello, world!</p></body></html>';
const blob = new Blob([html], { type: 'text/html'});
const url = URL.createObjectURL(blob);
window.open(url, '_blank');
```