package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/cmd/follow/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
)

type CreateFollowService struct {
	ctx context.Context
}

func NewCreateFollowService(ctx context.Context) *CreateFollowService {
	return &CreateFollowService{ctx: ctx}
}

func (s *CreateFollowService) CreateFollow(req *followrpc.CreateFollowRequest) error {
	followModel := &db.Follow{
		UserId:   req.UserId,
		FollowId: req.FollowId,
	}

	follow, err := db.GetFollow(s.ctx, followModel)
	if err != nil {
		return err
	}

	if follow == nil {
		// 下面三个操作应该是原子的
		err = db.CreateFollow(s.ctx, followModel)
		if err != nil {
			return err
		}
		err = rpc.AddFollowCount(s.ctx, &userrpc.AddFollowCountRequest{UserId: req.UserId, Count: 1})
		if err != nil {
			return err
		}
		err = rpc.AddFollowerCount(s.ctx, &userrpc.AddFollowerCountRequest{UserId: req.FollowId, Count: 1})
		if err != nil {
			return err
		}
	}
	return nil
}
