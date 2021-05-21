package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	pkg2 "vote/trash/internal/v1/pkg"
	result2 "vote/trash/internal/v1/result"
)

func AdminAuth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	log.Println(token)
	// 先去掉 token 前面的 bearer ，再解析，否则解析会 panic
	token, _ = pkg2.GetRawToken(token)
	parseToken, err := pkg2.ParseToken(token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest,
			result2.New(
				result2.TOKEN_PARSE_ERROR,
				result2.TOKEN_PARSE_ERROR.String()+" : "+err.Error(), nil))
		c.Abort()
	}

	if !IsAdminId(parseToken.Id) {
		log.Println(parseToken.Id + "，该 token 中的 id 不是 admin")
		c.JSON(http.StatusForbidden,
			result2.New(
				result2.ADMIN_FORBIDDEN,
				result2.ADMIN_FORBIDDEN.String(), nil))
		c.Abort()
	}
	c.Next()
}

func IsAdminId(id string) bool {
	return strings.Contains(id, "iaabb")
}

