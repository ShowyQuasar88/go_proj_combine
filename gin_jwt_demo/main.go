package main

import (
	"github.com/gin-gonic/gin"
	"github.com/showyquasar88/proj-combine/gin_jwt_demo/middleware"
	"github.com/showyquasar88/proj-combine/gin_jwt_demo/router"
)

type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	router.InitRouters(r)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
