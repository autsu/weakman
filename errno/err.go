package errno

import "errors"

var (
	LoginPasswordWrong   = errors.New("用户名或密码错误")
	RegisterPhoneIsExist = errors.New("手机号已被注册")
)

var (
	TopicUserIsVoted = errors.New("您已经投过票了")
	TopicPasswordIsWrong = errors.New("输入的密码不正确")
)

var (
	MysqlConnectError          = errors.New("连接数据库错误")
	MysqlSelectError           = errors.New("数据库查询错误")
	MysqlSelectNoData          = errors.New("未查询到数据")
	MysqlInsertError           = errors.New("数据库插入错误")
	MysqlGetLastInsertIdError  = errors.New("获取最后一次插入的 id 错误")
	MysqlStartTransactionError = errors.New("开启事务错误")
	MysqlLimitParamError       = errors.New("提供的分页参数错误，page 和 size 必须大于 0")
	MysqlScanError             = errors.New("scan 参数时发生错误")
	MysqlUpdateError           = errors.New("更新失败")
)

var (
	JwtCreateError = errors.New("token 创建失败")
	TokenInvalid   = errors.New("提供的 token 无效")
)

var (
	EncryptPasswordError = errors.New("密码加密为 md5 错误")
	PasswordNotEqual     = errors.New("密码的 md5 值不匹配")
)
