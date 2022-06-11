package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
)

type CheckFollowService struct {
	ctx context.Context
}

func NewCheckFollowService(ctx context.Context) *CheckFollowService {
	return &CheckFollowService{ctx: ctx}
}

func (s *CheckFollowService) CheckFollow(req *followrpc.CheckFollowRequest) (bool, error) {
	// 自己不能关注自己
	if req.UserId == req.FollowId {
		return true, nil
	}
	followModel := &db.Follow{
		UserId:   req.UserId,
		FollowId: req.FollowId,
	}

	follow, err := db.GetFollow(s.ctx, followModel)

	if err != nil {
		return false, err
	}
	if follow == nil {
		return false, nil
	}
	return true, nil
}
