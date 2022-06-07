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

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new CreateUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser 用户服务,返回用户信息，错误
func (s *MGetUserService) MGetUser(req *userrpc.MGetUserRequest) ([]*userrpc.User, error) {
	// 查询用户是否存在
	users, err := db.MQueryUserById(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	// 用户不存在则返回 用户不存在（UserNotExistErr）的错误
	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}
	us := make([]*userrpc.User, 0)

	var isFollows []bool

	// 如果需要调用 isfollow 函数，则调用。
	// 如果不需要 则设置为 true
	if req.ReturnIsFollow {
		isFollows, err = rpc.BatchIsFollow(s.ctx, &followrpc.MCheckFollowRequest{
			UserId:    req.TokenUserId,
			FollowIds: req.UserIds,
		})
		if err != nil {
			return nil, errno.FollowNotExistErr
		}
	}

	for i, u := range users {
		if user2 := pack.User(u); user2 != nil {
			if req.ReturnIsFollow {
				user2.IsFollow = isFollows[i]
			} else {
				user2.IsFollow = true
			}
			us = append(us, user2)
		}
	}
	return us, nil
}
