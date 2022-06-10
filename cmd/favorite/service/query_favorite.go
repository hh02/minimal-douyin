package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/favorite/dal/db"
	"github.com/hh02/minimal-douyin/cmd/favorite/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/favoriterpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
)

type QueryFavoriteService struct {
	ctx context.Context
}

func NewQueryFavoriteService(ctx context.Context) *QueryFavoriteService {
	return &QueryFavoriteService{ctx: ctx}
}
func (q *QueryFavoriteService) QueryFavorite(req *favoriterpc.QueryFavoriteByUserIdRequest) ([]*videorpc.Video, error) {
	// 查询用户点赞的所有视频 id
	videoIds, err := db.QueryFavoriteList(q.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	// 根据 videoIds , userId 远程调用 video 服务获取 video 列表
	rpcRequest := &videorpc.MGetVideoRequest{
		TokenUserId: req.UserId,
		VideoIds:    videoIds,
	}
	videos, err := rpc.MGetVideo(q.ctx, rpcRequest)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
