package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
)


type DeleteFollowService struct {
	ctx context.Context
}

func NewDeleteFollowService(ctx context.Context) *DeleteFollowService {
	return &DeleteFollowService{ctx: ctx}
}

func (s *DeleteFollowService) DeleteFollow(req *followrpc.DeleteFollowRequest) error {
	followModel := &db.Follow{
		UserId: req.UserId,
		FollowId: req.FollowId,
	}
	return db.DeleteFollow(s.ctx, followModel)
}
