package main

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/follow/pack"
	"github.com/hh02/minimal-douyin/cmd/follow/service"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// CreateFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) CreateFollow(ctx context.Context, req *followrpc.CreateFollowRequest) (resp *followrpc.CreateFollowResponse, err error) {
	resp = new(followrpc.CreateFollowResponse)

	// 参数检查：ID > 0 && 自己不能关注自己
	if req.UserId <= 0 || req.FollowId <= 0 || req.UserId == req.FollowId{
		resp.StatusCode, resp.StatusMessage = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateFollowService(ctx).CreateFollow(req)
	if err != nil {
		resp.StatusCode, resp.StatusMessage = pack.BuildBaseResp(err)
	}

	resp.StatusCode, resp.StatusMessage = pack.BuildBaseResp(errno.Success)

	return resp, nil
}

// DeleteFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) DeleteFollow(ctx context.Context, req *followrpc.DeleteFollowRequest) (resp *followrpc.DeleteFollowResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) QueryFollow(ctx context.Context, req *followrpc.QueryFollowRequest) (resp *followrpc.QueryFollowResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryFollower implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) QueryFollower(ctx context.Context, req *followrpc.QueryFollowerRequest) (resp *followrpc.QueryFollowerResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) CheckFollow(ctx context.Context, req *followrpc.CheckFollowRequest) (resp *followrpc.CheckFollowResponse, err error) {
	// TODO: Your code here...
	return
}
