package errno

import "errors"

var (
	PASSWORD_NOT_EQUAL = errors.New("密码的 md5 值不匹配")
)

// jwt
var (
	TOKEN_INVALID = errors.New("提供的 token 无效")
	
)

// db error
var (
	DB_COLLEGE_NAME_EXISIT = errors.New("系名已存在")
	DB_INSERT_ERROR = errors.New("数据库插入失败")
)


