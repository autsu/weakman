package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strings"
	result2 "vote/v2/enum/result"
	"vote/v2/pkg"
)

func AdminAuth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	logrus.Info(token)
	parseToken, err := pkg.ParseTokenWithBearer(token)
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

func StuAuth(c *gin.Context)  {
	token := c.GetHeader("Authorization")
	logrus.Info(token)

	_, err := pkg.ParseTokenWithBearer(token)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusForbidden, result2.NewWithCode(result2.TOKEN_INVALID))
		c.Abort()
	}
	c.Next()
}

