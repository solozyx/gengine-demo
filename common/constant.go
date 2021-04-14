package common

const (
	PathStatic = "/static/"

	// 用户初始密码
	UserInitPassword = "asdf1234"
	UserPasswordSalt = "e5f396a1-f18f-411f-b9f1-892928e0c723_gengine"
	UserIdKey        = "user_id"

	// 验证码
	VerifyCodeExpiration = 300 // 单位:秒

	// sign
	HeaderSignSalt = "com.obencn.gengine.c27d32c5-c878-470a-8a6f-107f7d3700dd"

	// jwt
	JwtTokenSecret     = "e2df5341-4f15-4c87-93fb-5f8f36b1865c_obencn.com_gengine"
	JwtPayloadIssuer   = "gengine.gateway"
	JwtPayloadAudience = "https://gengine.obencn.com"

	FileStoreLocalPath = "file/%s"
)
