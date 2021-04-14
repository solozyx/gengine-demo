package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// header
	ErrAppEUINotFound    = &Errno{Code: 10003, Message: "The request header LOLA-ACCESS-APPID not found."}
	ErrNonceNotFound     = &Errno{Code: 10004, Message: "The request header LOLA-ACCESS-NONCE not found."}
	ErrSignatureNotFound = &Errno{Code: 10005, Message: "The request header LOLA-ACCESS-SIGNATURE not found."}
	ErrJwtPayload        = &Errno{Code: 10006, Message: "JWT payload decode error."}
	ErrDataVerify        = &Errno{Code: 10007, Message: "数据完整性校验未通过"}
	ErrNonceTypeWrong    = &Errno{Code: 10008, Message: "LOLA-ACCESS-NONCE 数据类型错误"}
	ErrNonceTimeOut      = &Errno{Code: 10009, Message: "LOLA-ACCESS-NONCE 超时"}

	// system
	ErrValidation     = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase       = &Errno{Code: 20002, Message: "Database error."}
	ErrToken          = &Errno{Code: 20003, Message: "令牌验证失败"}
	ErrReply          = &Errno{Code: 20004, Message: "消息队列回复失败"}
	ErrSmsSend        = &Errno{Code: 20005, Message: "短信发送失败"}
	ErrTokenExpire    = &Errno{Code: 20006, Message: "令牌过期失效"}
	ErrResponseParse  = &Errno{Code: 20007, Message: "Response解析失败"}
	ErrJsonParse      = &Errno{Code: 20008, Message: "Json解析失败"}
	ErrProtoParse     = &Errno{Code: 20009, Message: "Proto解析失败"}
	ErrPermAuth       = &Errno{Code: 20010, Message: "权限验证失败"}
	ErrPermNotEnough  = &Errno{Code: 20011, Message: "权限不足"}
	ErrMustParaIsNull = &Errno{Code: 20013, Message: "必填参数不可为空"}
	ErrParaTypeWrong  = &Errno{Code: 20014, Message: "参数类型错误"}

	// user errors
	ErrUserNotFound                   = &Errno{Code: 20101, Message: "账号不存在"}
	ErrVerifyCodeSend                 = &Errno{Code: 20102, Message: "发送验证码错误"}
	ErrVerifyCodeNotExpire            = &Errno{Code: 20103, Message: "验证码有效期5分钟,还在有效期内"}
	ErrVerifyCodeExpired              = &Errno{Code: 20104, Message: "验证码已过期"}
	ErrUserPassword                   = &Errno{Code: 20105, Message: "密码错误"}
	ErrUserEmailUsed                  = &Errno{Code: 20106, Message: "该邮箱已被绑定,请尝试其他邮箱"}
	ErrSendEmail                      = &Errno{Code: 20107, Message: "发送邮件失败"}
	ErrVerifyCodeInput                = &Errno{Code: 20108, Message: "验证码错误,请重新输入"}
	ErrEmailInputAgain                = &Errno{Code: 20109, Message: "填写的邮箱地址,与获取验证码的邮箱不一致"}
	ErrConfirmVerifyCode              = &Errno{Code: 20110, Message: "校验验证码错误"}
	ErrUserEdit                       = &Errno{Code: 20113, Message: "用户信息修改错误"}
	ErrUserDisable                    = &Errno{Code: 20114, Message: "您的账号已被冻结"}
	ErrUserDelete                     = &Errno{Code: 20115, Message: "您的账号已被删除"}
	ErrUserAlreadyDisabled            = &Errno{Code: 20116, Message: "待处理的账号已被冻结"}
	ErrUserAlreadyDeleted             = &Errno{Code: 20117, Message: "待处理的账号已被删除"}
	ErrUserAlreadyNormal              = &Errno{Code: 20118, Message: "待处理的账号是正常状态"}
	ErrUserLoginFail                  = &Errno{Code: 20119, Message: "用户登录错误"}
	ErrUserChangePassword             = &Errno{Code: 20120, Message: "用户修改密码错误"}
	ErrUserCreate                     = &Errno{Code: 20121, Message: "创建成员错误"}
	ErrUserInfo                       = &Errno{Code: 20122, Message: "查看成员错误"}
	ErrUserDisableOperate             = &Errno{Code: 20123, Message: "冻结成员错误"}
	ErrUserEnableOperate              = &Errno{Code: 20124, Message: "解冻成员错误"}
	ErrUserEditOperate                = &Errno{Code: 20125, Message: "编辑成员错误"}
	ErrOrgNotFound                    = &Errno{Code: 20126, Message: "组织不存在"}
	ErrOrgDisable                     = &Errno{Code: 20127, Message: "组织已被冻结"}
	ErrOrgDeleted                     = &Errno{Code: 20128, Message: "组织已被删除"}
	ErrOrgLevelOneNotAllowDelete      = &Errno{Code: 20129, Message: "一级组织不允许删除"}
	ErrDepNotFound                    = &Errno{Code: 20130, Message: "部门不存在"}
	ErrDepDisable                     = &Errno{Code: 20131, Message: "部门已被冻结"}
	ErrDepDelete                      = &Errno{Code: 20132, Message: "部门已被删除"}
	ErrDepDeleteTwoLevelDeleteAnother = &Errno{Code: 20133, Message: "二级组织管理员不能删除其他组织部门"}
	ErrDepListTwoLevelAnother         = &Errno{Code: 20134, Message: "二级组织管理员不能查看其他组织部门"}
	ErrRolePermAddOnlyTwoLevelAdmin   = &Errno{Code: 20135, Message: "角色添加权限,仅二级组织管理员可操作"}
	ErrRolePermAddOtherOrg            = &Errno{Code: 20136, Message: "角色添加权限,不能操作其他组织的角色"}
	ErrUserRoleAddOnlySameOrg         = &Errno{Code: 20137, Message: "用户角色添加,仅限于操作本组织成员"}
	ErrExportUserOperateLog           = &Errno{Code: 20138, Message: "导出成员操作日志错误"}
	ErrRolePermConfig                 = &Errno{Code: 20139, Message: "角色配置权限错误"}
	ErrUserRoleConfig                 = &Errno{Code: 20140, Message: "用户配置角色错误"}

	ErrAppIdNotFound = &Errno{Code: 20206, Message: "没有找到对应的AppId,请确保LOLA-ACCESS-APPID数据正确"}
	ErrSmsCode       = &Errno{Code: 20207, Message: "短信验证码错误"}
	ErrParas         = &Errno{Code: 20208, Message: "参数验证失败"}

	ErrFileUpload = &Errno{Code: 20301, Message: "文件上传错误"}
	ErrFileInfo   = &Errno{Code: 20302, Message: "文件信息获取错误"}
)
