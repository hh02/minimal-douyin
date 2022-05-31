package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/user/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type UserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

// UserRegister 用户服务,返回用户信息，错误
func (s *UserService) User(req *userrpc.UserRequest) (*db.User, error) {
	// 查询用户是否存在
	user, err := db.QueryUserById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	// 用户不存在则返回 用户不存在（UserNotExistErr）的错误
	if user == nil {
		return nil, errno.UserNotExistErr
	}
	return user, err
}
