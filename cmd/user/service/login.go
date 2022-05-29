package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/hh02/minimal-douyin/cmd/user/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/douyin/core"
	"github.com/hh02/minimal-douyin/pkg/errno"
	"io"
)

type UserLoginService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewUserLoginService(ctx context.Context) *UserLoginService {
	return &UserLoginService{ctx: ctx}
}

// UserRegister 用户登录服务,返回用户ID、token、错误
func (s *UserLoginService) UserLogin(req *core.UserLoginRequest) (uint, string, error) {
	// 查询用户是否存在
	user, err := db.QueryUserByName(s.ctx, req.Username)
	if err != nil {
		return 0, "", err
	}
	// 用户不存在则返回 用户不存在（UserNotExistErr）的错误
	if user == nil {
		return 0, "", errno.UserNotExistErr
	}

	// 对密码进行 md5 加密
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, "", err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	// 验证密码是否正确 不正确则返回登录错误
	if passWord != user.Password {
		return 0, "", errno.LoginErr
	}
	// 得到加密后的密码 ， 与用户名组成 token 作为唯一标识。
	token := req.Username + passWord
	// 返回用户id、token、错误
	return user.ID, token, err
}
