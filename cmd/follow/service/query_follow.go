package service

import (
	"context"

	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
)

type QueryFollowService struct {
	ctx context.Context
}

func NewQueryFollowService(ctx context.Context) *QueryFollowService {
	return &QueryFollowService{ctx: ctx}
}

func (s *QueryFollowService) QueryFollow(req *followrpc.QueryFollowRequest) *followrpc.QueryFollowResponse {
	return nil
}