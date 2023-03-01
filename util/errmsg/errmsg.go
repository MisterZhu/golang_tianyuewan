package errmsg

const (
	SUCCSE = 200 //成功
	ERROR  = 400 //系统错误

	// // code 400 通用提示语报错
	// ERR_GENERAL = 400

	//code 1000 用户模块错误
	ERR_USER_USED        = 1001
	ERR_PASSWORD_WRONG   = 1002
	ERR_USER_NOT_EXIST   = 1003
	ERR_TOKEN_NOT_EXIST  = 1004
	ERR_TOKEN_EXPIRE     = 1005
	ERR_TOKEN_WRONG      = 1006
	ERR_TOKEN_TYPE_WRONG = 1007

	//code 2000 分类模块错误
	ERR_CATE_USED = 2001
)

var codemsg = map[int]string{
	SUCCSE: "请求成功",
	ERROR:  "请求失败",

	ERR_USER_USED:        "用户名已存在！",
	ERR_PASSWORD_WRONG:   "密码错误",
	ERR_USER_NOT_EXIST:   "用户不存在",
	ERR_TOKEN_NOT_EXIST:  "token不存在",
	ERR_TOKEN_EXPIRE:     "token已过期",
	ERR_TOKEN_WRONG:      "token不正确",
	ERR_TOKEN_TYPE_WRONG: "token格式错误",

	// ERR_GENERAL: 			"自定义参数",
	ERR_CATE_USED: "分类已存在",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
