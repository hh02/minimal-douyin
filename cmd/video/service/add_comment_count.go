package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/video/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
)

type AddCommentCountService struct {
	ctx context.Context
}

func NewAddCommentCountService(ctx context.Context) *AddCommentCountService {
	return &AddCommentCountService{ctx: ctx}
}

func (s *AddCommentCountService) AddCommentCount(req *videorpc.AddCommentCountRequest) error {
	return db.AddCommentCount(s.ctx, req.VideoId, req.Count)
}
