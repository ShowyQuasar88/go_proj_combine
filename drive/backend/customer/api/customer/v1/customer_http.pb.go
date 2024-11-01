// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v5.28.2
// source: api/customer/customer.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCustomerGetVerifyCode = "/api.customer.Customer/GetVerifyCode"

type CustomerHTTPServer interface {
	// GetVerifyCode 获取验证码
	GetVerifyCode(context.Context, *GetVerifyCodeReq) (*GetVerifyCodeResp, error)
}

func RegisterCustomerHTTPServer(s *http.Server, srv CustomerHTTPServer) {
	r := s.Route("/")
	r.GET("/customer/get-verify-code", _Customer_GetVerifyCode0_HTTP_Handler(srv))
}

func _Customer_GetVerifyCode0_HTTP_Handler(srv CustomerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVerifyCodeReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCustomerGetVerifyCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVerifyCode(ctx, req.(*GetVerifyCodeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetVerifyCodeResp)
		return ctx.Result(200, reply)
	}
}

type CustomerHTTPClient interface {
	GetVerifyCode(ctx context.Context, req *GetVerifyCodeReq, opts ...http.CallOption) (rsp *GetVerifyCodeResp, err error)
}

type CustomerHTTPClientImpl struct {
	cc *http.Client
}

func NewCustomerHTTPClient(client *http.Client) CustomerHTTPClient {
	return &CustomerHTTPClientImpl{client}
}

func (c *CustomerHTTPClientImpl) GetVerifyCode(ctx context.Context, in *GetVerifyCodeReq, opts ...http.CallOption) (*GetVerifyCodeResp, error) {
	var out GetVerifyCodeResp
	pattern := "/customer/get-verify-code"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCustomerGetVerifyCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
