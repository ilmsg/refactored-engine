# refactored-engine
golang การใช้ gorm กับ relationship, one to one, one to many, many to many


### many-to-many
relation ที่ gorm สร้างขึ้นให้ใน database

![![Alt text](2023-10-17_065412.png)](https://github.com/ilmsg/refactored-engine/blob/main/images/2023-10-17_065412.png?raw=true)


ผลลัพธ์การรันโค้ด

![![Alt text](2023-10-24_061018.png)](https://github.com/ilmsg/refactored-engine/blob/main/images/2023-10-24_161018.png?raw=true)


### ony-to-many

- 🐛 uuid ที่ใช้เป็น primary key ไม่สร้าง relation ให้เลย
