package main

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/favorite/service"
	"github.com/hh02/minimal-douyin/kitex_gen/favoriterpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// CreateFavorite implements the UserServiceImpl interface.
func (s *FavoriteServiceImpl) CreateFavorite(ctx context.Context, req *favoriterpc.CreateFavoriteRequest) (resp *favoriterpc.CreateFavoriteResponse, err error) {
	// TODO: Your code here...
	resp = new(favoriterpc.CreateFavoriteResponse)
	if req.UserId <= 0 || req.VideoId <= 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateFavoriteService(ctx).CreateFavorite(req)

	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}

// DeleteFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) DeleteFavorite(ctx context.Context, req *favoriterpc.DeleteFavoriteRequest) (resp *favoriterpc.DeleteFavoriteResponse, err error) {
	// TODO: Your code here...
	resp = new(favoriterpc.DeleteFavoriteResponse)
	if req.UserId <= 0 || req.VideoId <= 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	err = service.NewDeleteFavoriteService(ctx).DeleteFavorite(req)

	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}

// QueryFavoriteByUserId implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) QueryFavoriteByUserId(ctx context.Context, req *favoriterpc.QueryFavoriteByUserIdRequest) (resp *favoriterpc.QueryFavoriteByUserIdResponse, err error) {
	// TODO: Your code here...
	resp = new(favoriterpc.QueryFavoriteByUserIdResponse)
	if req.UserId <= 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}
	videos, err := service.NewQueryFavoriteService(ctx).QueryFavorite(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.Videos = videos
	return resp, nil
}
