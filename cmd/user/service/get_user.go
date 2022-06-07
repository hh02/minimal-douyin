package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/user/dal/db"
	"github.com/hh02/minimal-douyin/cmd/user/pack"
	"github.com/hh02/minimal-douyin/cmd/user/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type GetUserService struct {
	ctx context.Context
}

// NewGetUserService new CreateUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// GetUser 用户服务,返回用户信息，错误
func (s *GetUserService) GetUser(req *userrpc.GetUserRequest) (*userrpc.User, error) {
	// 查询用户是否存在
	user, err := db.QueryUserById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	// 用户不存在则返回 用户不存在（UserNotExistErr）的错误
	if user == nil {
		return nil, errno.UserNotExistErr
	}

	// 如果不需要查询是否关注，则直接设置为关注
	if !req.ReturnIsFollow {
		User := pack.User(user)
		User.IsFollow = true
		return User, err
	}
	// 否则，通过 follow 服务查询
	isFollow, err := rpc.IsFollow(s.ctx, &followrpc.CheckFollowRequest{
		UserId:   req.TokenUserId,
		FollowId: req.UserId,
	})
	if err != nil {
		return nil, errno.FollowNotExistErr
	}
	User := pack.User(user)
	User.IsFollow = isFollow
	return User, err
}
