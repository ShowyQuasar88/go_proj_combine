package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/showyquasar88/proj-combine/gin_jwt_demo/api/v1"
)

const (
	Login = "/login"
)

func LoginRouters(r *gin.Engine) {
	login := r.Group(Prefix + Login)
	loginV1 := login.Group("/v1")
	{
		loginV1.POST("", v1.Login)
	}
}
