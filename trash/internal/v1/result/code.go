package result

type Code int

const (
	LOGIN_FAIL_WRONG_PASSWORD Code = iota + 1
	SUCCESS
	SERVER_ERROR
	REGISTER_FAIL_PHONE_EXIST
	TOKEN_INVALID
	TOKEN_PARSE_ERROR
	BAD_REQUEST
	ADMIN_FORBIDDEN
)

func (c Code) String() string {
	var message string
	switch c {
	case LOGIN_FAIL_WRONG_PASSWORD:
		message = "用户名或密码错误"
	case SUCCESS:
		message = "操作成功"
	case SERVER_ERROR:
		message = "服务器内部错误"
	case REGISTER_FAIL_PHONE_EXIST:
		message = "注册失败：手机号已被注册"
	case TOKEN_INVALID:
		message = "提供的 token 无效"
	case BAD_REQUEST:
		message = "提供的参数无效"
	case ADMIN_FORBIDDEN:
		message = "管理员访问权限错误"
	case TOKEN_PARSE_ERROR:
		message = "token 解析失败"
	}
	return message
}
