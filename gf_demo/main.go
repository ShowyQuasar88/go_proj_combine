package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

/*
 * 统一返回结构
 */

type Response struct {
	Message string      `json:"message" dc:"消息提示"`
	Data    interface{} `json:"data" dc:"执行结果"`
}

type HelloReq struct {
	g.Meta `path:"/" method:"get" tags:"Test" summary:"Hello world"`
	Name   string `v:"required" json:"name" dc:"姓名"`
	Age    int    `v:"required" json:"age" dc:"年龄"`
}

type HelloResp struct {
	Content string `json:"content" dc:"返回结果"`
}

type Hello struct{}

func (Hello) Say(ctx context.Context, req *HelloReq) (res *HelloResp, err error) {
	res = &HelloResp{
		Content: fmt.Sprintf(
			"Hello %s! Your age is %d!",
			req.Name,
			req.Age,
		),
	}
	return
}

func Middleware(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		msg string
		res = r.GetHandlerResponse()
		err = r.GetError()
	)
	if err != nil {
		msg = err.Error()
	} else {
		msg = "success"
	}
	r.Response.WriteJson(&Response{
		Message: msg,
		Data:    res,
	})
}

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(Middleware)
		group.Bind(
			new(Hello),
		)
	})
	s.SetOpenApiPath("/api.json")
	s.SetSwaggerPath("/swagger")
	s.SetPort(8000)
	s.Run()
}
