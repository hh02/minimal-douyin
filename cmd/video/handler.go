package main

import (
	"context"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *videorpc.CreateVideoRequest) (resp *videorpc.CreateVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideoByUserId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoByUserId(ctx context.Context, req *videorpc.GetVideoByUserIdRequest) (resp *videorpc.GetVideoByUserIdResponse, err error) {
	// TODO: Your code here...
	return
}
