package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/cmd/follow/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
)

// 查询关注列表
type QueryFollowerService struct {
	ctx context.Context
}

func NewQueryFollowerService(ctx context.Context) *QueryFollowerService {
	return &QueryFollowerService{ctx: ctx}
}

func (s *QueryFollowerService) QueryFollower(req *followrpc.QueryFollowerRequest) ([]*userrpc.User, error) {
	userIds, err := db.QueryFollower(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	rpcRequest := new(userrpc.MGetUserRequest)
	rpcRequest.UserIds = userIds
	rpcRequest.TokenUserId = req.UserId
	users, err := rpc.MGetUser(s.ctx, rpcRequest)
	if err != nil {
		return nil, err
	}
	return users, nil
}
