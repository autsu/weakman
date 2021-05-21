package router

import (
	"github.com/gin-gonic/gin"
	handler2 "vote/trash/internal/v1/handler"
	middleware2 "vote/trash/internal/v1/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()


	admin := r.Group("/ad111")
	{
		admin.POST("/login", handler2.AdminLogin)

		// 下面的路由都会经过 AdminAuth 验证
		admin.Use(middleware2.AdminAuth)
		admin.GET("/college", handler2.AdminQueryAllCollege)
		admin.GET("/college/id/:id", handler2.AdminQueryCollegeByIdWithProfession)
		//admin.GET("/college/name/:name")
		admin.POST("/college", handler2.AdminAddCollege)
		admin.DELETE("/college/:id", handler2.AdminDeleteCollege)
		admin.PUT("/college", handler2.AdminUpdateCollege)
	}

	return r
}
