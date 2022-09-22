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

### `GET` /summary
> 获取用户数量、订单数量的统计
```json
"usernumber": "2"
"ordernumber": "1"
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

### `GET` /deluser/:id
#### ！！！！删除某用户会导致其下的订单、帖子和评论也被删除
> 删除某用户
```json
"status": "succeess"
```

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

### `GET` /delproduct/:id
#### !!!删除一个商品会导致关联的订单被全部删除
> 删除某商品
```json
"status": "succeess"
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
> 预期返回
```json
"status": "success"
```

### `GET` /orderlist
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

### `GET` /delorder/:id
> 删除某订单
```json
"status": "succeess"
```

### `GET` /findorder/:userid
> 根据用户id查找订单
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

### `GET` /announcementlist
> 获取公告列表
```json
[
    {
        "id": 1,
        "title": "关于停服维护的通知",
        "content": "本站将于9月20号进行为期2天的停服维护",
        "date": "2022/9/15"
        "img": ""
    }
    // ...
]
```

### `POST` /addannouncement
> 添加一条公告
```json
"title": "关于停服维护时间延长的通知",
"content": "由于技术原因，停服时间延长至9月24号",
"date": "2022/9/21",
"img": ""  // 可为空
```
> 预期返回
```json
"status": "success"
```

### `GET` /delannouncement/:id
> 删除一条公告
```json
"status":"succeess"
```

### `GET` /postlist
> 获取帖子列表
```json
[
    {
        "id":1,
        "userid":1,
        "username":"pphui8",
        "title":"这是一条帖子",
        "content":"恭喜本站开始运营",
        "browseNumber":100,
        "date":"2022-09-15",
        "legal":1,
        "elite":1,
        "img":"https://img1.imgtp.com/2022/09/22/QreqdzeV.jpg",
        "tag":"公告"
    }
    // ...
]
```

### `GET` /postcomment/:postid
> 获取帖子评论
```json

```

### `GET` /delpost/:id
> 删除一条帖子

### `GET` /delcomment/:id
> 删除一条评论