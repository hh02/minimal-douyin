package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/video/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
)

type CreateVideoService struct {
	ctx context.Context
}

func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

func (s *CreateVideoService) CreateVideo(req *videorpc.CreateVideoRequest) error {
	return db.CreateVideo(s.ctx, &db.Video{
		UserId:   req.UserId,
		PlayUrl:  req.PlayUrl,
		CoverUrl: req.CoverUrl,
		Title:    req.Title,
	})
}
