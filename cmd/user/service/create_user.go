package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/hh02/minimal-douyin/cmd/user/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser 用户服务,返回用户信息，错误
func (s *CreateUserService) CreateUser(req *userrpc.CreateUserRequest) (error, int64) {
	// 查询用户是否存在
	user, err := db.QueryUserByName(s.ctx, req.Username)

	if err != nil {
		return err, 0
	}

	// 用户存在则返回 用户存在（UserAlreadyExistErr）的错误
	if user != nil {
		return errno.UserAlreadyExistErr, 0
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err, 0
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	return db.CreateUser(s.ctx, &db.User{
		Username: req.Username,
		Password: passWord,
	})
}
