package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/favorite/dal/db"
	"github.com/hh02/minimal-douyin/cmd/favorite/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/favoriterpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type DeleteFavoriteService struct {
	ctx context.Context
}

func NewDeleteFavoriteService(ctx context.Context) *DeleteFavoriteService {
	return &DeleteFavoriteService{ctx: ctx}
}

func (d *DeleteFavoriteService) DeleteFavorite(req *favoriterpc.DeleteFavoriteRequest) error {
	favoriteModel := &db.Favorite{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	}
	rowsAffected, err := db.DeleteFavorite(d.ctx, favoriteModel)
	if err != nil {
		return err
	}
	if rowsAffected > 0 {
		// 取消点赞让 video 点赞数减一
		err = rpc.AddFavoriteCount(d.ctx, &videorpc.AddFavoriteCountRequest{
			VideoId: req.VideoId,
			Count:   -1,
		})
		return err
	} else {
		return errno.DeleteErr
	}
}
