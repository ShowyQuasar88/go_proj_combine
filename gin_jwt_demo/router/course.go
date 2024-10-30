package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/showyquasar88/proj-combine/gin_jwt_demo/api/v1"
	"github.com/showyquasar88/proj-combine/gin_jwt_demo/middleware"
)

const (
	Course = "/course"
)

func CourseRouters(r *gin.Engine) {
	course := r.Group(Prefix + Course)
	course.Use(middleware.Auth())
	courseV1 := course.Group("/v1")
	{
		courseV1.GET("", v1.GetCourse)
		courseV1.POST("", v1.AddCourse)
	}
}
