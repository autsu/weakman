package pkg

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
	errno2 "vote/trash/internal/v1/errno"
)

var (
	Issuer    = "zz"               // Issuer 签发者
	ExpiresAt = time.Hour * 24     // ExpiresAt 24 hours
	Secret    = []byte("SGlsb3g=") // Secret 加密秘钥
)

const (
	Head       = "Authorization" // Head 请求头
	HeadPrefix = "Bearer "       // HeadPrefix token 前缀
)

func CreateJwt(id string) string {
	// 指定信息
	claims := jwt.StandardClaims{
		Audience:  "",                               // 受众
		ExpiresAt: time.Now().Add(ExpiresAt).Unix(), // 过期时间
		Id:        id,                               // 编号
		IssuedAt:  time.Now().Unix(),                // 签发时间
		Issuer:    Issuer,                           // 签发人
		NotBefore: time.Now().Unix(),                // 生效时间
		Subject:   "login",                          // 主题
	}
	// 创建 token
	token, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(Secret)
	if err != nil {
		panic(err)
	}
	return token
}

func ParseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return Secret, nil
		})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*jwt.StandardClaims)
	if ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, errno2.TOKEN_INVALID
}

// GetRawToken 获取原生 token
func GetRawToken(token string) (string, error) {
	var rawToken string
	if len(token) != 0 {
		// 以空格为分隔符，将字符串切割为多个子串
		// 获取初始 token 值，即去掉了前缀 Bearer 后的值
		rawToken = strings.Fields(token)[1]
		return rawToken, nil
	} else {
		return "", errno2.TOKEN_INVALID
	}
}
