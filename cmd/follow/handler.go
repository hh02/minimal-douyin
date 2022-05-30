package main

import (
	"context"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// CreateFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) CreateFollow(ctx context.Context, req *followrpc.CreateFollowRequest) (resp *followrpc.CreateFollowResponse, err error) {
	// TODO: Your code here...
	return
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

// GetFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) GetFollow(ctx context.Context, req *followrpc.GetFollowRequest) (resp *followrpc.GetFollowResponse, err error) {
	// TODO: Your code here...
	return
}
