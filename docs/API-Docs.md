# Web API Docs

## API 主要调用返回状态码及其信息

> 采用 restful 风格 + 状态码与 msg 信息的方式实现

HTTP Code

- 200 调用成功
- 404 NotFound 页面不存在
- 405 method not allowed 请求方式有误

Return code

> 其他逻辑上的问题一律返回200并通过接口的code来判断问题，为了与http code做区分，非成功的返回值都为负数

- 1 成功
- -301 token丢失/权限缺失(未能进入HandelFunc)
- -403 参数请求异常
- -501 RPC调用错误
- -502 服务器处理error

## API 调用说明

>测试环境部署在 base 127.0.0.1:8888/api/ 调用时使用 form-data 进行参数编码并以 json 格式返回

### 前置response wrapper

| 参数 | 类型      | 是否必须 | 说明                                 |
| :--- | --------- | -------- | ------------------------------------ |
| code | int       | 是       | 自定状态码，1表示成功，小于0表示异常 |
| msg  | string    | 是       | 状态message                          |
| data | interface | 否       | 数据实体，当状态不为1时不返回该字段  |

示例

```json
{
    "code": 1,
    "msg": "成功",
    "data": {}
}

```



### 登陆

路由: /api/login

请求方式: POST

请求参数：

| 参数名   | 类型   | 是否必须 | 说明                                       |
| -------- | ------ | -------- | ------------------------------------------ |
| name     | string | 是       | 登陆用户名 pattern:"^[a-zA-Z0-9_-]{6,15}$" |
| password | string | 是       | 登陆密码 pattern:"^().{6,20}$"             |

请求示例：

```shell
curl --location --request POST 'localhost:8888/api/login'\
--form 'name=cerbur'\
--form 'password=123456'
```

响应参数：

| 参数名 | 类型   | 是否必须 | 说明  |
| ------ | ------ | -------- | ----- |
| data   | string | 是       | Token |

响应示例：

```json
{
    "code": 1,
    "data": "c58edc366dfb11eca07bacde48001122",
    "msg": "成功"
}
```

### 获取profile

路由: /api/profile

请求方式: GET

请求参数：

| 参数名        | 类型   | 是否必须 | 说明                               |
| ------------- | ------ | -------- | ---------------------------------- |
| Authorization | Header | 是       | Authorization为登录成功返回的token |

请求示例：

```shell
curl --header "Authorization:10c66960738011ecbc5dacde48001122"\
--request GET 'http://localhost:8888/api/profile'
```

响应参数：

| 参数名          | 类型   | 是否必须 | 说明         |
| --------------- | ------ | -------- | ------------ |
| username        | string | 是       | 用户名       |
| nickname        | string | 是       | 昵称         |
| profile_picture | string | 是       | 头像的文件名 |

响应示例：

```json
{
    "code": 1,
    "msg": "成功",
    "data": {
        "username": "cerbur12345",
        "nickname": "奶盖🐶⬇️ahahah",
        "profile_picture": "742631749dc5b808c28885811bc0298f7a8a1008a990d0ce4dda4835f6559886.png"
    }
}
```



### 修改nickname

路由: /api/profile/nickname

请求方式: PUT

请求参数：

| 参数名        | 类型   | 是否必须 | 说明                               |
| ------------- | ------ | -------- | ---------------------------------- |
| Authorization | Header | 是       | Authorization为登录成功返回的token |
| nickname      | string | 是       | 昵称 pattern:"^().{2,15}$"         |

请求示例：

```shell
curl --header "Authorization:10c66960738011ecbc5dacde48001122" \
--form "nickname=cerbur" \
--request PUT 'http://localhost:8888/api/profile/nickname'
```

响应参数：

| 参数名   | 类型   | 是否必须 | 说明 |
| -------- | ------ | -------- | ---- |
| nickname | string | 是       | 昵称 |

响应示例：

```json
{
    "code": 1,
    "msg": "成功",
    "data": {
        "nickname": "奶盖圈😊",
        "profile_picture": ""
    }
}
```

### 修改头像

路由: /api/profile/picture

请求方式: PUT

请求参数：

| 参数名        | 类型   | 是否必须 | 说明                               |
| ------------- | ------ | -------- | ---------------------------------- |
| Authorization | Header | 是       | Authorization为登录成功返回的token |
| nickname      | string | 是       | 昵称 pattern:"^().{2,15}$"         |

请求示例：

```shell
curl --header "Authorization:10c66960738011ecbc5dacde48001122" \
--form 'file=@/Users/test/Downloads/test.JPG' \
--request PUT 'http://localhost:8888/api/profile/picture'
```

响应参数：

| 参数名 | 类型                             | 是否必须 | 说明 |
| ------ | -------------------------------- | -------- | ---- |
| file   | Multipart/File只允许JPG和PNG类型 | 是       | 头像 |

响应示例：

```json
{
    "code": 1,
    "msg": "成功",
    "data": {
        "nickname": "",
        "profile_picture": "742631749dc5b808c28885811bc0298f7a8a1008a990d0ce4dda4835f6559886.png"
    }
}
```

### 获取图片

路由: /api/img

请求方式: GET

请求参数：

| 参数名 | 类型   | 是否必须 | 说明   |
| ------ | ------ | -------- | ------ |
| img    | string | 是       | 文件名 |

请求示例：

```shell
curl --location --request POST 'localhost:8888/api/img'\
--form 'img=742631749dc5b808c28885811bc0298f7a8a1008a990d0ce4dda4835f6559886.png'
```

响应:

二进制文件