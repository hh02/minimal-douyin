package main

import (
	"context"
	"github.com/hh02/minimal-douyin/kitex_gen/video_rpc"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *video_rpc.CreateVideoRequest) (resp *video_rpc.CreateVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideoByUserId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoByUserId(ctx context.Context, req *video_rpc.GetVideoByUserIdRequest) (resp *video_rpc.GetVideoByUserIdResponse, err error) {
	// TODO: Your code here...
	return
}
