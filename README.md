# Khaos_backend
the backend of Khaos server written in golang

## data flow
main -> routers -> api ( -> database ) -> api -> return

## APIs

### `GET` /
> 测试服务器状态
```json
"status": "200"
```

### `GET` /usersnumber
> 获取用户数量
```json
"usernumber": "2"
```

### `GET` /userslist
> 获取用户列表
```json
[
    {
        "id": 1,
        "username": "pphui8",
        "registerdate": "2022/8/30",
        "phone": "19861550668",
        "privilege": "manager",
    },
    // .....
]
```

### `GET` /user/:id
> 获取某用户的信息
```json
"id": 1,
"username": "pphui8",
"descript": "Ciallo～(∠・ω< )⌒★",
"registerdate": "2022/8/30",
"phone": "19861550668",
"privilege": "manager",
```

<!-- ### `GET` /deluser/:id
> 删除某用户
```json
"status": "succeess"
``` -->

### `POST` /login
> 管理员登录
```json
"publickey": "3D75AD4DB3E952BC206E2DAFED2D91DC"
```
返回：
```json
"status": "successed",
"id": 1,
"username": "pphui8",
```

### `GET` /productlist
> 获取商品列表
```json
[
    {
        "id": 1,
        "productname": "伸腿瞪眼丸",
        "price": 999,
        "descipt": "药到病除",
        "stock": 993,
        "sale": 7,
        "type": "保健品",
        "status": "在售",
    }
    // ...
]
```

### `GET` /product/:id
> 获取某商品详情
```json
[
    {
        "id": 1,
        "productname": "伸腿瞪眼丸",
        "descript": "药到病除",
        "price": 997,
        "img": "base64: xxxxxxx",
        "stock": 999,
        "sale": 7,
        "type": "保健品",
        "status": "在售",
    }
    // ...
]
```

### `POST` /addproduct
> 添加某商品
```json
"productname": "伸腿瞪眼丸",
"descript": "药到病除",
"price": 997,
"img": "base64: xxxxxxx",
"stock": 999,
"sale": 7,
"type": "保健品",
"status": "在售",
```
> 返回
```json
"status": "success"
```

### `GET` /order
> 获取全部订单列表
```json
[
    {
        "id": 1,
        "userid": 1,
        "username": "pphui8",
        "productid": 1,
        "productname": "伸腿瞪眼丸",
        "price": 999,
        "number": 1,
        "date": "2022/8/31",
        "location": "上海",
        "status": "未发货"
    }
    // ...
]
```