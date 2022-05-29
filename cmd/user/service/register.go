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

type UserRegisterService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewUserRegiaterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{ctx: ctx}
}

// UserRegister 用户注册服务,返回用户ID、token、错误
func (s *UserRegisterService) UserRegister(req *core.UserRegisterRequest) (uint, string, error) {
	// 查询用户是否存在
	user, err := db.QueryUserByName(s.ctx, req.Username)
	if err != nil {
		return 0, "", err
	}
	// 用户已经存在则返回 用户已经存在（UserAlreadyExistErr）的错误
	if user != nil {
		return 0, "", errno.UserAlreadyExistErr
	}

	// 对密码进行 md5 加密
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, "", err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	//得到加密后的密码 ， 与用户名组成 token 作为唯一标识。
	token := req.Username + passWord
	err, Id := db.UserRegister(s.ctx, &db.User{
		Name:     req.Username,
		Password: passWord,
	})
	return Id, token, err
}
