package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/video/dal/db"
	"github.com/hh02/minimal-douyin/cmd/video/pack"
	"github.com/hh02/minimal-douyin/cmd/video/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
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
	user, err := rpc.GetUser(s.ctx, &userrpc.GetUserRequest{UserId: req.UserId, ReturnIsFollow: false})
	if err != nil {
		return nil, err
	}
	user.IsFollow = false // 自己不能关注自己
	videos := pack.Videos(videoModels)
	for i := 0; i < len(videos); i++ {
		videos[i].Author = user
	}
	return videos, nil
}

func (s *QueryVideoService) QueryVideoByTime(req *videorpc.QueryVideoByTimeRequest) (nextTime int64, videos []*videorpc.Video, err error) {
	videoModels, err := db.QueryVideoByTime(s.ctx, req.LatestTime)
	if err != nil {
		return 0, nil, err
	}
	if len(videoModels) == 0 {
		return 0, nil, errno.ServiceErr.WithMessage("no more videos")
	}

	returnIsFollow := (req.TokenUserId > 0)

	// 提取出不重复的 userIds
	userIds := pack.UserIds(videoModels)
	userMap, err := rpc.MGetUserMap(s.ctx, &userrpc.MGetUserRequest{
		UserIds:        userIds,
		TokenUserId:    req.TokenUserId,
		ReturnIsFollow: returnIsFollow,
	})
	if err != nil {
		return 0, nil, err
	}
	videos = pack.Videos(videoModels)

	for i := 0; i < len(videos); i++ {
		if user, ok := userMap[videos[i].Author.UserId]; ok {
			videos[i].Author = user
		}
	}

	nextTime = videoModels[len(videoModels)-1].CreateTime
	err = nil
	return
}
