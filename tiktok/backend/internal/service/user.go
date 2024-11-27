package service

import (
	v1 "backend/api/v1"
	"backend/internal/biz"
	"backend/internal/pkg/response"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/transport"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"net/http"
	"net/url"
)

// UserService 是用户服务的实现
type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUseCase
}

// NewUserService 创建一个用户服务
func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.Response, error) {
	user := &biz.User{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	if err := s.uc.Register(ctx, user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return response.SuccessWithMsg("用户已存在"), nil
		}
		return response.Error(v1.ErrorCode_SYSTEM_ERROR, err.Error()), err
	}
	return response.SuccessWithMsg("注册成功"), nil
}

func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.Response, error) {
	user := &biz.User{
		Username: req.Username,
		Password: req.Password,
	}
	userResp, err := s.uc.Login(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(v1.ErrorCode_USER_NOT_FOUND, "用户不存在或密码错误"), nil
		}
		return response.Error(v1.ErrorCode_SYSTEM_ERROR, err.Error()), nil
	}

	// 设置 cookies
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return response.Error(v1.ErrorCode_SYSTEM_ERROR, "系统错误"), nil
	}
	userTokenCookie := &http.Cookie{
		Name:     "user_token",
		Value:    userResp.Token,
		Path:     "/",
		MaxAge:   int(userResp.ExpiresIn),
		HttpOnly: true,  // 设置 httponly 前端将无法通过 getCookie 来获取，防止 xss 攻击
		Secure:   false, // true的话只有https才会发送cookie，http就不会发送
		SameSite: http.SameSiteStrictMode,
	}
	usernameCookie := &http.Cookie{
		Name:     "username",
		Value:    url.QueryEscape(userResp.UserInfo.Username),
		Path:     "/",
		MaxAge:   int(userResp.ExpiresIn),
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	userIDCookie := &http.Cookie{
		Name:     "userID",
		Value:    userResp.UserInfo.Id,
		Path:     "/",
		MaxAge:   int(userResp.ExpiresIn),
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	ht, ok := tr.(khttp.Transporter)
	if !ok {
		return response.Error(v1.ErrorCode_SYSTEM_ERROR, "系统错误"), nil
	}
	ht.ReplyHeader().Add("Set-Cookie", userTokenCookie.String())
	ht.ReplyHeader().Add("Set-Cookie", usernameCookie.String())
	ht.ReplyHeader().Add("Set-Cookie", userIDCookie.String())

	// 类型转化成 any
	data, err := anypb.New(userResp)
	if err != nil {
		return response.Error(v1.ErrorCode_SYSTEM_ERROR, err.Error()), nil
	}

	return response.SuccessWithData(data), nil
}

func (s *UserService) Logout(ctx context.Context, empty *emptypb.Empty) (resp *v1.Response, err error) {
	// 删除浏览器 cookie
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return response.Error(v1.ErrorCode_SYSTEM_ERROR, "系统错误"), nil
	}

	ht, ok := tr.(khttp.Transporter)
	if !ok {
		return response.Error(v1.ErrorCode_SYSTEM_ERROR, "系统错误"), nil
	}

	// 从 cookie 中获取信息
	cookies := ht.Request().Cookies()
	var username, userID string
	for _, cookie := range cookies {
		switch cookie.Name {
		case "username":
			username, err = url.QueryUnescape(cookie.Value)
			if err != nil {
				return response.Error(v1.ErrorCode_SYSTEM_ERROR, "系统错误"), nil
			}
		case "userID":
			userID = cookie.Value
		}
	}

	userTokenCookie := &http.Cookie{
		Name:     "user_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // 立即过期
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	usernameCookie := &http.Cookie{
		Name:     "username",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	userIDCookie := &http.Cookie{
		Name:     "userID",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	ht.ReplyHeader().Add("Set-Cookie", userTokenCookie.String())
	ht.ReplyHeader().Add("Set-Cookie", usernameCookie.String())
	ht.ReplyHeader().Add("Set-Cookie", userIDCookie.String())

	err = s.uc.Logout(ctx, username, userID)
	if err != nil {
		return response.Error(v1.ErrorCode_SYSTEM_ERROR, "系统错误"), nil
	}

	return response.Success(), nil
}
