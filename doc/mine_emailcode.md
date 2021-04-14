## 接口 - 发送验证码到邮箱
 - 说明<br>发送验证码到邮箱<br>

 - 地址<br>
     /api/v1/mine/emailcode

 - 请求方式<br>
     POST   

 - 请求参数<br>
     **MineEmailCodeRequest**
     
     header (示例)
     
     | 字段         | 类型                             | 说明 |
     | ------------ | -------------------------------- | ---- |
     | Content-Type | application/json                 |      |
     | timestamp    | 1234567890                       |      |
     | sign         | f7bb1121e2f2f2048ccc83d8a6b37643 |      |
     
     body
     
     | 字段  |  类型  | 说明 |
     | :---: | :----: | :--- |
     | email | string | 邮箱 |
     
 - 返回值说明<br>
     **MineEmailCodeResponse**

     | 字段 |   类型    | 说明 |
     | :--: | :-------: | :--- |
     | N/A  | BaseModel |      |
     
     **BaseModel**
     
     | 字段 |  类型  | 说明              |
     | :--: | :----: | :---------------- |
     | Code | number | 返回码            |
     | Msg  | string | 返回说明/错误说明 |
     
示例

```json
{
    "email":"18813141886@163.com"
}
```

```json
{
    "code": 0,
    "msg": "SUCCESS"
}
```

```json
{
    "code": 20103,
    "msg": "验证码还在有效期内; "
}
```
