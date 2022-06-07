package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/like/dal/db"
	"github.com/hh02/minimal-douyin/cmd/like/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/likerpc"
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
		panic(err)
	}
	if rowsAffected > 0 {
		err := rpc.AddLikeCont(d.ctx, &likerpc.CreateLikeRequest{VideoId: req.VideoId, UserId: req.UserId})
		if err != nil {
			return err
		}
	}
	return nil
}
