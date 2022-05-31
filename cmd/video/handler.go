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

// MGetVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetVideo(ctx context.Context, req *videorpc.MGetVideoRequest) (resp *videorpc.MGetVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryVideoByUserId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryVideoByUserId(ctx context.Context, req *videorpc.QueryVideoByUserIdRequest) (resp *videorpc.QueryVideoByUserIdResponse, err error) {
	// TODO: Your code here...
	return
}
