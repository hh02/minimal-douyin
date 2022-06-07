package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/cmd/follow/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
)

// 查询关注列表
type QueryFollowService struct {
	ctx context.Context
}

func NewQueryFollowService(ctx context.Context) *QueryFollowService {
	return &QueryFollowService{ctx: ctx}
}

func (s *QueryFollowService) QueryFollow(req *followrpc.QueryFollowRequest) ([]*userrpc.User, error) {
	userIds, err := db.QueryFollow(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if len(userIds) == 0 {
		return nil, nil
	}

	rpcRequest := &userrpc.MGetUserRequest{
		UserIds:        userIds,
		TokenUserId:    req.UserId,
		ReturnIsFollow: false,
	}

	users, err := rpc.MGetUser(s.ctx, rpcRequest)
	if err != nil {
		return nil, err
	}
	return users, nil
}
