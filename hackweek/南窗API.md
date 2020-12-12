# 南窗API 接口文档



## 1.1 API V1 接口说明

- 接口基准地址：`47.103.200.147:8080`
- 服务端已开启 CORS 跨域支持
- API V1 认证统一使用 Token 认证
- 需要授权的 API ，必须在请求头中使用 `Authorization` 字段提供 `token` 令牌
- 使用 HTTP Status Code 标识状态
- 数据返回格式统一使用 JSON

### 1.1.1  支持的请求方法

- GET（SELECT）：从服务器取出资源（一项或多项）。
- POST（CREATE）：在服务器新建一个资源。
- PUT（UPDATE）：在服务器更新资源（客户端提供改变后的完整资源）。
- DELETE（DELETE）：从服务器删除资源。

  

  

### 1.1.2  通用返回状态说明

| *状态码* | *含义*                 | *说明*                                              |
| -------- | :--------------------- | --------------------------------------------------- |
| 200      | OK                     | 请求成功                                            |
| 201      | CREATED                | 创建成功                                            |
| 204      | DELETED                | 删除成功                                            |
| 400      | BAD REQUEST            | 请求的地址不存在或者包含不支持的参数                |
| 401      | UNAUTHORIZED           | 未授权                                              |
| 403      | FORBIDDEN              | 被禁止访问                                          |
| 404      | NOT FOUND              | 请求的资源不存在                                    |
| 422      | Unprocesable entity    | [POST/PUT/PATCH] 当创建一个对象时，发生一个验证错误 |
| 500      | INTERNAL SERVER ERROR  | 内部错误                                            |
| 1001     | ERROR_USERNAME_USED    | 用户名已存在                                        |
| 1002     | ERROR_PASSWORD_WRONG   | 密码错误                                            |
| 1003     | ERROR_USER_NOT_EXIST   | 用户不存在                                          |
| 1004     | ERROR_TOKEN_EXIST      | TOKEN不存在                                         |
| 1005     | ERROR_TOKEN_RUNTIME    | TOKEN已过期                                         |
| 1006     | ERROR_TOKEN_WRONG      | TOKEN错误                                           |
| 1007     | ERROR_TOKEN_TYPE_WRONG | TOKEN格式错误                                       |
| 1008     | ERROR_USER_NO_RIGHT    | 用户无权限                                          |
| 1009     | ERROR_PASSWORD_EDIT    | 密码修改成功                                        |







------

## 1.2. 登录

### 1.2.1. 登录验证接口

- 请求路径：/login
- 请求方法：POST
- 请求参数

| 参数名   | 参数说明 | 备注       |
| -------- | -------- | ---------- |
| username | 用户名   | 6~18个字符 |
| password | 密码     | 8~18个字符 |

- 响应参数

| 参数名 | 参数说明 | 备注            |
| ------ | -------- | --------------- |
| token  | 令牌     | 基于 jwt 的令牌 |
| msg    | 信息     |                 |
| status | 状态码   |                 |
| id     | 用户ID   |                 |

- 响应数据

```json
{
    "id": 2,
    "msg": "OK",
    "status": 200,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InF3ZSIsImV4cCI6MTYwNzc4MTY2NiwiaXNzIjoiaGFja3dlZWsyIn0.5M-FcjnCOCqk6QCbXyQZ4u12ifWppHxGbA1XTSNdgAg"
}
```

## 1.3. 用户管理

### 1.3.1. 用户数据列表

- 请求路径：/user
- 请求方法：GET
- 请求参数

| 参数名   | 参数说明     | 备注     |
| -------- | ------------ | -------- |
| pagenum  | 当前页码     | 可以为空 |
| pagesize | 每页显示条数 | 可以为空 |

- 响应参数

| 参数名   | 参数说明     | 备注 |
| -------- | ------------ | ---- |
| pagenum  | 当前页码     |      |
| pagesize | 每页显示条数 |      |

- 响应数据

```json
{
    "data": [
        {
            "ID": 10,
            "CreatedAt": "2020-12-09T10:17:33+08:00",
            "UpdatedAt": "2020-12-11T20:02:46+08:00",
            "DeletedAt": null,
            "username": "lijianhao",
            "password": "SdYD5qUcRiBBtQ==",
            "sex": "",
            "info": "健豪",
            "birth_day": "",
            "location": "",
            "role": 0
        },
        {
            "ID": 11,
            "CreatedAt": "2020-12-11T17:12:13+08:00",
            "UpdatedAt": "2020-12-11T17:12:13+08:00",
            "DeletedAt": null,
            "username": "lijianhao123",
            "password": "SdYD5qUcRiBBtQ==",
            "sex": "",
            "info": "",
            "birth_day": "",
            "location": "",
            "role": 0
        },
        {
            "ID": 12,
            "CreatedAt": "2020-12-11T20:31:01+08:00",
            "UpdatedAt": "2020-12-11T20:31:01+08:00",
            "DeletedAt": null,
            "username": "taicaile",
            "password": "SdYD5qUcRiBBtQ==",
            "sex": "",
            "info": "",
            "birth_day": "",
            "location": "",
            "role": 0
        }
    ],
    "msg": "OK",
    "status": "ok"
}
```

### 1.3.2. 注册用户

- 请求路径：/user/add
- 请求方法：POST
- 请求参数

- 响应参数（JSON传,前端没有相应填入框，目前只填两个参数）

| 参数名   | 参数说明    | 备注 |
| -------- | ----------- | ---- |
| username | 用户名      |      |
| password | 用户角色 ID |      |

- 响应数据

```json
{
    "msg": "OK",
    "status": 200
}
```



### 1.3.3 修改密码

- 请求路径：/users/:id
- 请求方法：PUT
- 请求参数

| 参数名   | 参数说明 | 备注       |
| -------- | -------- | ---------- |
| old_pwd  | 旧密码   |            |
| new_pwd  | 新密码   | 6~18个字符 |
| username | 用户名   | 8~18个字符 |

- 响应参数

- 响应数据

```json
{
    	"status" : 200,
		"msg" :  "密码修改成功"
}
```

# 功能模块 #

### 1.4.1 发布动态及查看

- 请求路径：/makeActivity        发布动态
- 请求路径: /checkActivity           查看动态
- 请求方法：POST
- 请求参数

| 参数名    | 参数说明 | 备注                 |
| --------- | -------- | -------------------- |
| tittle    | 标题     | 最大18字             |
| username  | 用户名   |                      |
| info      | 具体内容 | 最大150字            |
| image_url | 图片链接 | 以链接形式传入数据库 |
| likes     | 点赞数   | 假的                 |

- 响应参数

| 参数名    | 参数说明 | 备注 |
| --------- | -------- | ---- |
| tittle    | 标题     |      |
| info      | 具体内容 |      |
| image_url | 图片链接 |      |
| likes     | 点赞数   |      |

- 响应数据

```json
发布动态
{
    "msg": "success",
    "status": "Ok"
}
查看动态
{
    	Username: Username,
		Cate:     Cate,
		Tittle:   Tittle,
		Info:     Info,
		ImageUrl: ImageUrl,
		Likes:    999,
	}
    "msg": "success",
    "status": "Ok"
}
```

# 

### 1.4.2 搜索

- 请求路径：/search
- 请求方法： POST
- 请求参数

| 参数名 | 参数说明           | 备注     |
| ------ | ------------------ | -------- |
| tittle | 你要搜索的商品简介 | JSON格式 |

- 响应参数
- 响应数据

```json
{
    "data": [

        {
            "owner": "健豪",
            "avatar": "剑豪",
            "tittle": "健豪",
            "likes": 999,
            "image_url": "健豪",
            "cate": 0
        }
    ],
    "msg": 200,
    "status": "ok"
}
```



### 1.4.3  5种分类的商品

- 请求路径：/figure/:id	
- 请求方法：GET
- 请求参数

| 参数名 | 参数说明 | 备注                          |
| ------ | -------- | ----------------------------- |
| id     | 选取分类 | 分0，1，2，3，4对应不同的商品 |

- 响应参数

| 参数名    | 参数说明   | 备注                         |
| --------- | ---------- | ---------------------------- |
| owner     | 商家名     |                              |
| avatar    | 商家头像   |                              |
| tittle    | 标题简介   |                              |
| likes     | 点赞数     | 假的                         |
| image_url | 商品图链接 |                              |
| cate      | 分类       | 0,1,2,3,4分别代表5种商品数据 |

- 响应数据

- ```js
   {
              "owner": "",
              "avatar": "",
              "tittle": "",
              "likes": 0,
              "image_url": "",
              "cate": 0
          }
  ```

# 1.4.4 编辑用户提交（被前端砍了，不做了，端口写了，🐎的）
