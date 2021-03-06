package main

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/video/service"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *videorpc.CreateVideoRequest) (resp *videorpc.CreateVideoResponse, err error) {
	resp = new(videorpc.CreateVideoResponse)
	if req.UserId <= 0 || len(req.CoverUrl) == 0 || len(req.PlayUrl) == 0 || len(req.Title) == 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}
	err = service.NewCreateVideoService(ctx).CreateVideo(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}

// MGetVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetVideo(ctx context.Context, req *videorpc.MGetVideoRequest) (resp *videorpc.MGetVideoResponse, err error) {
	resp = new(videorpc.MGetVideoResponse)
	if len(req.VideoIds) == 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	videos, err := service.NewMGetVideoService(ctx).MGetVideo(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.Videos = videos
	return resp, nil
}

// QueryVideoByUserId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryVideoByUserId(ctx context.Context, req *videorpc.QueryVideoByUserIdRequest) (resp *videorpc.QueryVideoByUserIdResponse, err error) {
	resp = new(videorpc.QueryVideoByUserIdResponse)
	if req.UserId <= 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}
	videos, err := service.NewQueryVideoService(ctx).QueryVideoByUserId(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.Videos = videos
	return resp, nil
}

func (s *VideoServiceImpl) QueryVideoByTime(ctx context.Context, req *videorpc.QueryVideoByTimeRequest) (resp *videorpc.QueryVideoByTimeResponse, err error) {
	resp = new(videorpc.QueryVideoByTimeResponse)
	nextTime, videos, err := service.NewQueryVideoService(ctx).QueryVideoByTime(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.NextTime = nextTime
	resp.Videos = videos
	return resp, nil
}

func (s *VideoServiceImpl) AddFavoriteCount(ctx context.Context, req *videorpc.AddFavoriteCountRequest) (resp *videorpc.AddFavoriteCountResponse, err error) {
	resp = new(videorpc.AddFavoriteCountResponse)
	if req.VideoId <= 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	err = service.NewAddFavoriteCountService(ctx).AddFavoriteCount(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}

func (s *VideoServiceImpl) AddCommentCount(ctx context.Context, req *videorpc.AddCommentCountRequest) (resp *videorpc.AddCommentCountResponse, err error) {
	resp = new(videorpc.AddCommentCountResponse)
	if req.VideoId <= 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	err = service.NewAddCommentCountService(ctx).AddCommentCount(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}
