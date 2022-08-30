# Khaos_backend
the backend of Khaos server written in golang

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
        "publickey": "3D75AD4DB3E952BC206E2DAFED2D91DC"
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
"publickey": "3D75AD4DB3E952BC206E2DAFED2D91DC"
```

### `POST` /login
> 管理员登录
```json
"status": "successed",
"id": 1,
"publickey": "3D75AD4DB3E952BC206E2DAFED2D91DC"
```