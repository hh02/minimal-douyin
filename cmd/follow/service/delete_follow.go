package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/cmd/follow/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
)

type DeleteFollowService struct {
	ctx context.Context
}

func NewDeleteFollowService(ctx context.Context) *DeleteFollowService {
	return &DeleteFollowService{ctx: ctx}
}

func (s *DeleteFollowService) DeleteFollow(req *followrpc.DeleteFollowRequest) error {
	followModel := &db.Follow{
		UserId:   req.UserId,
		FollowId: req.FollowId,
	}

	rowsAffected, err := db.DeleteFollow(s.ctx, followModel)
	if err != nil {
		return err
	}
	if rowsAffected > 0 {
		err := rpc.AddFollowCount(s.ctx, &userrpc.AddFollowCountRequest{UserId: req.UserId, Count: -1})
		if err != nil {
			return err
		}
		err = rpc.AddFollowerCount(s.ctx, &userrpc.AddFollowerCountRequest{UserId: req.FollowId, Count: -1})
		if err != nil {
			return err
		}
	}
	return nil
}
