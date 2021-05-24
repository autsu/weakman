package result

type Code int

const (
	SUCCESS                   Code = 20000
	UNKNOW_ERROR              Code = 50010
	LOGIN_FAIL_WRONG_PASSWORD Code = 30001
	USER_NOT_EXIST            Code = 30013
	CREATE_TOKEN_ERROR        Code = 30002
	SERVER_ERROR              Code = 3
	REGISTER_FAIL_PHONE_EXIST Code = 40001
	TOKEN_INVALID             Code = 30003
	TOKEN_PARSE_ERROR         Code = 6
	BAD_REQUEST               Code = 7
	ADMIN_FORBIDDEN           Code = 8
	USER_IS_VOTED             Code = 30022
	NO_DATA                   Code = 100
	TOPIC_PASSWORD_ERROR      Code = 110
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
		message = "注册失败：该手机号已被注册"
	case TOKEN_INVALID:
		message = "提供的登录凭证无效"
	case BAD_REQUEST:
		message = "提供的参数无效"
	case ADMIN_FORBIDDEN:
		message = "管理员访问权限错误"
	case TOKEN_PARSE_ERROR:
		message = "登录凭证解析失败"
	case UNKNOW_ERROR:
		message = "未知错误"
	case CREATE_TOKEN_ERROR:
		message = "创建登录凭证失败，请再试一次"
	case USER_NOT_EXIST:
		message = "该用户不存在"
	case USER_IS_VOTED:
		message = "您已经投过票了"
	case NO_DATA:
		message = "未查询到数据"
	case TOPIC_PASSWORD_ERROR:
		message = "输入的密码错误"
	}
	return message
}
