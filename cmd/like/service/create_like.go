package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/like/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/likerpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type CreateLikeService struct {
	ctx context.Context
}

func NewCreateLikeService(ctx context.Context) *CreateLikeService {
	return &CreateLikeService{ctx: ctx}
}

func (s *CreateLikeService) CreateLike(req *likerpc.CreateLikeRequest) error {
	likeModel := &db.Like{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	}

	// 查询是否存在
	like, err := db.GetLike(s.ctx, likeModel)
	if err != nil {
		return err
	}
	if like == nil {
		err = db.CreateLike(s.ctx, likeModel)
		return err
	} else {
		return errno.LikeAlreadyExistErr
	}
}
