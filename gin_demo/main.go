package main

import (
	"github.com/showyquasar88/proj-combine/gin_demo/router"
	"log"
)

type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func main() {
	// 没有使用中间件
	//initDB("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	// 使用中间件
	LoggerMiddleware(initDB)("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")

	r := router.InitRouters()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

// initDB
// 假设存在如下的一个函数，用于初始化数据库连接。现在我想要增加一些功能，但是不想修改原始函数，此时就可以用中间件来完成
func initDB(connStr string) {
	log.Println("初始化数据库", connStr)
}

// LoggerMiddleware
// 新建一个方法，接收的参数和原始方法相同，然后在内部可以在他的前后做一些功能上的增加
func LoggerMiddleware(in func(connStr string)) func(connStr string) {
	return func(connStr string) {
		log.Println("call LoggerMiddleware begin")
		in(connStr)
		log.Println("call LoggerMiddleware end")
	}
}
