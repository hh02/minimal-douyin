package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/favorite/dal/db"
	"github.com/hh02/minimal-douyin/cmd/favorite/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/favoriterpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type CreateFavoriteService struct {
	ctx context.Context
}

func NewCreateFavoriteService(ctx context.Context) *CreateFavoriteService {
	return &CreateFavoriteService{ctx: ctx}
}

func (s *CreateFavoriteService) CreateFavorite(req *favoriterpc.CreateFavoriteRequest) error {
	favoriteModel := &db.Favorite{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	}

	// 查询是否存在
	favorite, err := db.GetFavorite(s.ctx, favoriteModel)
	if err != nil {
		return err
	}
	if favorite == nil {
		err = db.CreateFavorite(s.ctx, favoriteModel)
		if err != nil {
			return err
		}
		// video 点赞数加一
		err = rpc.AddFavoriteCount(s.ctx, &videorpc.AddFavoriteCountRequest{
			VideoId: req.VideoId,
			Count:   1,
		})
		return err
	} else {
		return errno.FavoriteAlreadyExistErr
	}
}
