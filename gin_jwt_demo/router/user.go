package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/showyquasar88/proj-combine/gin_jwt_demo/api/v1"
	"github.com/showyquasar88/proj-combine/gin_jwt_demo/middleware"
)

const (
	User = "/user"
)

func UserGroup(r *gin.Engine) {
	user := r.Group(Prefix + User)
	user.Use(middleware.Auth())
	userV1 := user.Group("/v1")
	{
		userV1.GET("", v1.GetUser)
		userV1.POST("", v1.AddUser)
	}
}
