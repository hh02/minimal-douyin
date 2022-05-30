package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
)


type CreateFollowService struct {
	ctx context.Context
}

func NewCreateFollowService(ctx context.Context) *CreateFollowService {
	return &CreateFollowService{ctx: ctx}
}

func (s *CreateFollowService) CreateFollow(req *followrpc.CreateFollowRequest) error {
	followModel := &db.Follow{
		UserId: req.UserId,
		FollowId: req.FollowId,
	}
	return db.CreateFollow(s.ctx, followModel)
}
