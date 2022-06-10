package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/video/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
)


type AddFavoriteCountService struct {
	ctx context.Context
}

func NewAddFavoriteCountService(ctx context.Context) *AddFavoriteCountService {
	return &AddFavoriteCountService{ctx: ctx}
}

func (s *AddFavoriteCountService) AddFavoriteCount(req *videorpc.AddFavoriteCountRequest) error {
	return db.AddFavoriteCount(s.ctx, req.VideoId, req.Count)
}