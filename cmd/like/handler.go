package main

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/like/service"
	"github.com/hh02/minimal-douyin/kitex_gen/likerpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct{}

// CreateLike implements the UserServiceImpl interface.
func (s *LikeServiceImpl) CreateLike(ctx context.Context, req *likerpc.CreateLikeRequest) (resp *likerpc.CreateLikeResponse, err error) {
	// TODO: Your code here...
	resp = new(likerpc.CreateLikeResponse)
	if req.UserId > 0 || req.VideoId <= 0 {
		return resp, nil
	}
	err = service.NewCreateLikeService(ctx).CreateLike(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *LikeServiceImpl) GetUser(ctx context.Context, req *likerpc.DeleteLikeRequest) (resp *likerpc.DeleteLikeResponse, err error) {
	// TODO: Your code here...
	resp = new(likerpc.DeleteLikeResponse)
	if req.UserId <= 0 || req.VideoId <= 0 {
		return resp, nil
	}
	err = service.NewDeleteLikeService(ctx).DeleteLike(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *LikeServiceImpl) MGetUser(ctx context.Context, req *likerpc.QueryLikeByUserIdRequest) (resp *likerpc.QueryLikeByUserIdResponse, err error) {
	// TODO: Your code here...
	resp = new(likerpc.QueryLikeByUserIdResponse)
	if req.UserId <= 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}
	videos, err := service.NewQueryLikeService(ctx).QueryLike(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.Videos = videos
	return resp, nil
}
