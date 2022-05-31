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

func (s *CheckFollowService) CheckFollow(req *followrpc.CheckFollowRequest) error {
	followModel := &db.Follow{
		UserId:   req.UserId,
		FollowId: req.FollowId,
	}

	_, err := db.GetFollow(s.ctx, followModel)
	if err != nil {
		return err
	}
	return nil
}
