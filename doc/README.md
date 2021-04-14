# 北分CMS管理后台 - 接口文档
## 接口说明
- 除
    (1)登录接口
外,都需要在header中传入AccessToken值，用以校验用户权限
- 除
    (1)文件上传接口
外,其他接口都需要生成签名（sign）字段，放入header，用以校验数据完整性，以及防篡改  
规则如下： md5(method+host+path+body+timestamp+salt)

## 我的
- [校验用户是否存在(不验token)](mine_exist.md)
- [登录(不验token)](mine_login.md)
- [发送验证码到邮箱(不验token)](mine_emailcode.md)
- [校验验证码(不验token)](mine_confirmcode.md)
- [重置密码(不验token)](mine_changepassword.md)

## 视频
- [视频文件上传](video_upload.md)
- [视频文件取消上传](video_cancel.md)
- [根据id查询视频信息](video_getbyid.md)
- [根据关键词查询视频信息](video_getbykeywords.md)
- [更新视频关键字](video_keywords.md)
- [设置启用/停用](video_enable.md)
- [视频列表](video_videolist.md)

## APP端
- [文件上传](app_file_upload.md)
- [文件信息获取](app_file_info.md)
