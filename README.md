## embedded-album

> 嵌入式电子相册



### 特性

- 图像轮播
- 后台管理



### API文档

> 交换数据格式为json交换

#### 基本响应格式

```json
{
    "success": true,
    "error": "",
    "data": null
}
```



#### 用户认证

接口前面带`*`表示需要验证；

使用JWT进行认证， 将token放入http头的Authorization字段，并使用Bearer开头；



#### 登录 POST /login

请求示例

```json
{
    "username":"admin",
    "password":"316316"
}
```

响应示例

```json
{
    "success": false,
    "error": "用户名或者密码不正确",
    "data": null
}
```

```json
{
    "success": true,
    "error": "",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODc3MjAwMTksImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU4NzQ2MDgxOX0.PPYhsbiA_3xjC_zl8nXNrZrHRD7ofl5pTLB2g2uyVw8"
}
```



#### *图片上传 POST /pictures

请求示例

```http
Content-Disposition: form-data; 
name="file";
filename="12.jpg"
Content-Type: image/jpeg
```

响应示例

```json
{
    "success": true,
    "error": "",
    "data": null
}
```



#### *图片获取 GET  /pictures

请求示例 （无

响应示例

```json
{
    "success": true,
    "error": "",
    "data": [
        {
            "id": 2,
            "created_at": 1587456698,
            "origin_name": "2020-04-18 22-30-33 的屏幕截图.png",
            "upload_name": "dd9ea5a85bad85029bdd9fc3f07fe41a.png"
        },
        {
            "id": 3,
            "created_at": 1587460834,
            "origin_name": "12.jpg",
            "upload_name": "be7639c0052505fb66feb885f5cb864c.jpg"
        }
    ]
}
```





#### *图片删除 DELETE /pictures/:upload_name

请求示例 DELETE /pictures/dd9ea5a85bad85029bdd9fc3f07fe41a.png

响应示例

```json
{
    "success": true,
    "error": "",
    "data": null
}
```

