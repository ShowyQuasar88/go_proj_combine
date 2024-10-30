package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/showyquasar88/proj-combine/gin_demo/api/v1"
)

const (
	Prefix  = "/api"
	Version = "/v1"
	User    = "/user"
	Course  = "/course"
)

func AuthCheck(param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("call check auth func ", param)
		c.Next()
	}
}

func InitRouters() *gin.Engine {
	// 下面两行等效于 gin.Default()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), AuthCheck("123123"))
	courseGroup := r.Group(Prefix + Version + Course)
	{
		courseGroup.GET("", v1.GetCourse)
		courseGroup.POST("", v1.AddCourse)
	}
	userGroup := r.Group(Prefix + Version + User)
	{
		userGroup.GET("", v1.GetUser)
		userGroup.POST("", v1.AddUser)
	}
	return r
}
