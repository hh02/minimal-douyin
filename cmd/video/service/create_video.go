package service

import (
	"context"

	"github.com/hh02/minimal-douyin/kitex_gen/video_rpc/"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
)

type CreateVideoService struct {
	ctx context.Context
}

func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

func (s *CreateVideoService) CreateVideo(req *videorpc.CreateVideoRequest) 