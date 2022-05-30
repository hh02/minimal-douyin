package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type GetFollowService struct {
	ctx context.Context
}

func NewGetFollowService(ctx context.Context) *GetFollowService {
	return &GetFollowService{ctx: ctx}
}

func (s *GetFollowService) GetFollow(req *followrpc.GetFollowRequest) *followrpc.GetFollowResponse {
	followModel := &db.Follow{
		UserId:   req.UserId,
		FollowId: req.FollowId,
	}

	_, err := db.GetFollow(s.ctx, followModel)
	if err != nil {
		return &followrpc.GetFollowResponse{
			StatusCode:    errno.FollowNotExistErrCode,
			StatusMessage: err.Error(),
		}
	}
	return &followrpc.GetFollowResponse{
		StatusCode: errno.SuccessCode,
	}
}
