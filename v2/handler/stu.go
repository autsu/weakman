package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"vote/v2/enum/result"
	"vote/v2/errno"
	"vote/v2/model"
	"vote/v2/service"
)

type StuHandler struct {
	StuService service.StuService
}

func (h *StuHandler) Login(c *gin.Context) {
	var loginVO = struct {
		Username string `binding:"required"`
		Password string `binding:"required"`
	}{}
	//var s model.Stu
	if err := c.ShouldBindJSON(&loginVO); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusOK, result.NewWithCode(result.BAD_REQUEST))
		return
	}

	token, err := h.StuService.Login(loginVO.Username, loginVO.Password)
	if err != nil {
		logrus.Error(err)
		if errors.Is(err, errno.MysqlConnectError) {
			c.JSON(http.StatusOK,
				result.NewWithCode(result.SERVER_ERROR))
			return
		}
		if errors.Is(err, errno.LoginPasswordWrong) {
			c.JSON(http.StatusOK,
				result.NewWithCode(result.LOGIN_FAIL_WRONG_PASSWORD))
			return
		}
		if errors.Is(err, errno.JwtCreateError) {
			c.JSON(http.StatusOK,
				result.NewWithCode(result.CREATE_TOKEN_ERROR))
			return
		}
		if errors.Is(err, errno.MysqlSelectNoData) {
			c.JSON(http.StatusOK,
				result.NewWithCode(result.USER_NOT_EXIST))
			return
		}
		c.JSON(http.StatusOK,
			result.NewWithCode(result.UNKNOW_ERROR))
		return
	}

	c.JSON(http.StatusOK,
		result.NewWithCodeAndData(result.SUCCESS, gin.H{
			"token": token,
		}))
}

func (h *StuHandler) Register(c *gin.Context) {
	var s model.Stu
	if err := c.ShouldBindJSON(&s); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusOK,
			result.NewWithCode(result.BAD_REQUEST))
		return
	}

	_, err := h.StuService.Register(&s)
	if err != nil {
		logrus.Error(err)
		if errors.Is(err, errno.RegisterPhoneIsExist) {
			c.JSON(http.StatusOK,
				result.NewWithCode(result.REGISTER_FAIL_PHONE_EXIST))
			return
		}
		c.JSON(http.StatusOK,
			result.NewWithCode(result.SERVER_ERROR))
		return
	}
	c.JSON(http.StatusOK,
		result.New(result.SUCCESS, "注册成功", nil))
}

func (h *StuHandler) GetInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	stu, err := h.StuService.GetInfoByToken(token)
	if err != nil {
		if errors.Is(err, errno.TokenInvalid) {
			c.JSON(http.StatusOK, result.NewWithCode(result.TOKEN_INVALID))
			return
		}
		if errors.Is(err, errno.MysqlSelectNoData) {
			c.JSON(http.StatusOK, result.NewWithCode(result.USER_NOT_EXIST))
			return
		}
		c.JSON(http.StatusOK, result.NewWithCode(result.SERVER_ERROR))
		return
	}
	stu.Password = ""
	c.JSON(http.StatusOK,
		result.NewWithCodeAndData(result.SUCCESS, stu))
}

func (h *StuHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, result.NewWithCode(result.SUCCESS))
}
