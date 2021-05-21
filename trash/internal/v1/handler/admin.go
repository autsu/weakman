package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	model2 "vote/v1/model"
	result2 "vote/v1/result"
	service2 "vote/v1/service"
)

var AdminLogin = func(c *gin.Context) {
	var admin model2.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest,
			result2.New(result2.BAD_REQUEST,
				result2.BAD_REQUEST.String()+" : "+err.Error(), nil))
		return
	}
	log.Println(admin)

	token, err := service2.AdminLogin(admin.Username, admin.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			result2.New(
				result2.LOGIN_FAIL_WRONG_PASSWORD,
				result2.LOGIN_FAIL_WRONG_PASSWORD.String(), nil))
		return
	}

	c.JSON(http.StatusOK,
		result2.New(
			result2.SUCCESS,
			result2.SUCCESS.String(),
			token))
}

var AdminQueryAllCollege = func(c *gin.Context) {
	colleges, err := service2.CollegeGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			result2.New(
				result2.SERVER_ERROR,
				result2.SERVER_ERROR.String()+" : "+err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, result2.New(
		result2.SUCCESS,
		result2.SUCCESS.String(),
		colleges))
}

var AdminAddCollege = func(c *gin.Context) {
	var col model2.College
	if err := c.ShouldBindJSON(&col); err != nil {
		c.JSON(http.StatusBadRequest,
			result2.New(result2.BAD_REQUEST,
				result2.BAD_REQUEST.String()+" : "+err.Error(), nil))
		return
	}
	if err := service2.CollegeAdd(&col); err != nil {
		c.JSON(http.StatusInternalServerError, result2.New(
			result2.SERVER_ERROR,
			result2.SERVER_ERROR.String()+" : "+err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, result2.New(
		result2.SUCCESS,
		result2.SUCCESS.String(), nil))
}

var AdminDeleteCollege = func(c *gin.Context) {
	id := c.Param("id")
	n, err := service2.CollegeSoftDelete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result2.New(
			result2.SERVER_ERROR,
			result2.SERVER_ERROR.String()+" : "+err.Error(), nil))
		return
	}
	msg := fmt.Sprintf(", 删除了 %d 条数据", n)
	c.JSON(http.StatusOK, result2.New(
		result2.SUCCESS,
		result2.SUCCESS.String()+msg, nil))
}

var AdminUpdateCollege = func(c *gin.Context) {
	var col model2.College
	if err := c.ShouldBindJSON(&col); err != nil {
		c.JSON(http.StatusBadRequest,
			result2.New(result2.BAD_REQUEST,
				result2.BAD_REQUEST.String()+" : "+err.Error(), nil))
		return
	}

	n, err := service2.CollegeUpdate(&col)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result2.New(
			result2.SERVER_ERROR,
			result2.SERVER_ERROR.String()+" : "+err.Error(), nil))
		return
	}
	msg := fmt.Sprintf(", 更新了 %d 条数据", n)
	c.JSON(http.StatusOK, result2.New(
		result2.SUCCESS,
		result2.SUCCESS.String()+msg, nil))
}

var AdminQueryCollegeByIdWithProfession = func(c *gin.Context) {
	id := c.Param("id")
	cs, err := service2.CollegeQueryByIdWithProfession(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result2.New(
			result2.SERVER_ERROR,
			result2.SERVER_ERROR.String()+" : "+err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, result2.New(
		result2.SUCCESS,
		result2.SUCCESS.String(), cs))
}
