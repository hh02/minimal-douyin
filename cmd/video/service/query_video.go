package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/video/dal/db"
	"github.com/hh02/minimal-douyin/cmd/video/rpc"
	"github.com/hh02/minimal-douyin/cmd/video/pack"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
)

type QueryVideoService struct {
	ctx context.Context
}

func NewQueryVideoService(ctx context.Context) *QueryVideoService {
	return &QueryVideoService{ctx: ctx}
}

func (s *QueryVideoService) QueryVideoByUserId(req *videorpc.QueryVideoByUserIdRequest) ([]*videorpc.Video, error) {
	videoModels, err := db.QueryVideoByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	user, err := rpc.GetUser(s.ctx, req.UserId)

	if err != nil {
		return nil, err
	}

	videos := pack.VideoModels2RpcMessages(videoModels)

	for i := 0; i < len(videos); i++ {
		videos[i].UserName = user.Name
		videos[i].
	}



}