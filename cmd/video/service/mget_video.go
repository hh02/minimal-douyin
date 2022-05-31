package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/video/dal/db"
	"github.com/hh02/minimal-douyin/cmd/video/pack"
	"github.com/hh02/minimal-douyin/cmd/video/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
)

type MGetVideoService struct {
	ctx context.Context
}

func NewMGetVideoService(ctx context.Context) *MGetVideoService {
	return &MGetVideoService{ctx: ctx}
}

func (s *MGetVideoService) MGetVideoService(req *videorpc.MGetVideoRequest) ([]*videorpc.Video, error) {

	videoModels, err := db.MgetVideo(s.ctx, req.VideoIds)
	if err != nil {
		return nil, err
	}

	// 提取出所有用户ID并去重
	userIds := pack.UserIds(videoModels)

	userMap, err := rpc.MGetUser(s.ctx, &userrpc.MGetUserRequest{UserIds: userIds})
	if err != nil {
		return nil, err
	}

	// convert db.Video to videorpc.Video
	videos := pack.Videos(videoModels)

	for i := 0; i < len(videos); i++ {
		if user, ok := userMap[videos[i].Author.UserId]; ok {
			videos[i].Author = user
		}
	}

	return videos, nil

}