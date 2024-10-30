package router

import (
	"github.com/gin-gonic/gin"
)

const (
	Prefix = "/api"
)

func InitRouters(r *gin.Engine) {
	UserGroup(r)
	CourseRouters(r)
	LoginRouters(r)
}
