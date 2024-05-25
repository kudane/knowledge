# :mag: ปัญหา

### :bulb: clone git ไม่ได้

หากเจอว่าเกิด sign in ไม่ได้ให้ใช้คำสั่งด้านล่าง

```
git config --system --unset credential.helper

git clone your_project
```

```
git config --global http.sslVerify false

git clone your_project
```
