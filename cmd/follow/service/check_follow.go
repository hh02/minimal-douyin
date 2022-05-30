package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type CheckFollowService struct {
	ctx context.Context
}

func NewCheckFollowService(ctx context.Context) *CheckFollowService {
	return &CheckFollowService{ctx: ctx}
}

func (s *CheckFollowService) CheckFollow(req *followrpc.CheckFollowRequest) *followrpc.CheckFollowResponse {
	followModel := &db.Follow{
		UserId:   req.UserId,
		FollowId: req.FollowId,
	}

	_, err := db.GetFollow(s.ctx, followModel)
	if err != nil {
		return &followrpc.CheckFollowResponse{
			StatusCode:    errno.FollowNotExistErrCode,
			StatusMessage: err.Error(),
		}
	}
	return &followrpc.CheckFollowResponse{
		StatusCode: errno.SuccessCode,
	}
}
