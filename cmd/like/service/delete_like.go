package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/like/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/likerpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type DeleteLikeService struct {
	ctx context.Context
}

func NewDeleteLikeService(ctx context.Context) *DeleteLikeService {
	return &DeleteLikeService{ctx: ctx}
}

func (d *DeleteLikeService) DeleteLike(req *likerpc.DeleteLikeRequest) error {
	likeModel := &db.Like{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	}
	rowsAffected, err := db.DeleteLike(d.ctx, likeModel)
	if err != nil {
		return err
	}
	if rowsAffected > 0 {
		return nil
	} else {
		return errno.DeleteErr
	}
}
