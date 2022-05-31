package main

import (
	"context"

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
	if req.UserId <= 0 || req.FollowId <= 0 || req.UserId == req.FollowId {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateFollowService(ctx).CreateFollow(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
	}

	resp.Status = errno.BuildStatus(errno.Success)

	return resp, nil
}

// DeleteFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) DeleteFollow(ctx context.Context, req *followrpc.DeleteFollowRequest) (resp *followrpc.DeleteFollowResponse, err error) {
	resp = new(followrpc.DeleteFollowResponse)

	if req.UserId <= 0 || req.FollowId <= 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	err = service.NewDeleteFollowService(ctx).DeleteFollow(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
	}
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}

// QueryFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) QueryFollow(ctx context.Context, req *followrpc.QueryFollowRequest) (resp *followrpc.QueryFollowResponse, err error) {
	resp = new(followrpc.QueryFollowResponse)

	if req.UserId <= 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewQueryFollowService(ctx).QueryFollow(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.Users = users
	return resp, nil
}

// QueryFollower implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) QueryFollower(ctx context.Context, req *followrpc.QueryFollowerRequest) (resp *followrpc.QueryFollowerResponse, err error) {
	resp = new(followrpc.QueryFollowerResponse)

	if req.UserId <= 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewQueryFollowerService(ctx).QueryFollower(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.Users = users
	return resp, nil
}

// CheckFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) CheckFollow(ctx context.Context, req *followrpc.CheckFollowRequest) (resp *followrpc.CheckFollowResponse, err error) {
	resp = new(followrpc.CheckFollowResponse)
	if req.UserId <= 0 || req.FollowId <= 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}
	// 自己不能关注自己
	if req.UserId == req.FollowId {
		resp.Status = errno.BuildStatus(errno.Success)
		resp.IsFollow = false
		return resp, nil
	}

	isFollow, err := service.NewCheckFollowService(ctx).CheckFollow(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.IsFollow = isFollow
	return resp, nil
}
