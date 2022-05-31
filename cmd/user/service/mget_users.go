package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/user/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type MGetUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// UserRegister 用户服务,返回用户信息，错误
func (s *MGetUserService) MGetUser(req *userrpc.MGetUserRequest) ([]*db.User, error) {
	// 查询用户是否存在
	users, err := db.MQueryUserById(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	// 用户不存在则返回 用户不存在（UserNotExistErr）的错误
	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}
	return users, err
}
