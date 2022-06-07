package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/like/dal/db"
	"github.com/hh02/minimal-douyin/cmd/like/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/likerpc"
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
	rowsAffected, err := db.CreateLike(s.ctx, likeModel)
	if err != nil {
		panic(err)
	}
	if rowsAffected > 0 {
		err := rpc.AddLikeCont(s.ctx, &likerpc.CreateLikeRequest{VideoId: req.VideoId})
		if err != nil {
			return err
		}
	}
	return nil
}
