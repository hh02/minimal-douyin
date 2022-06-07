package service

import (
	"context"
	"errors"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/cmd/follow/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"gorm.io/gorm"
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

	_, err := db.GetFollow(s.ctx, followModel)
	if err != nil {
		// 如果关注不存在再关注
		if errors.Is(err, gorm.ErrRecordNotFound) {

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
			return nil
		}
		return err
	}
	return nil
}
