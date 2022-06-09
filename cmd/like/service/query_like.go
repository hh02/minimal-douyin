package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/like/dal/db"
	"github.com/hh02/minimal-douyin/cmd/like/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/likerpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
)

type QueryLikeService struct {
	ctx context.Context
}

func NewQueryLikeService(ctx context.Context) *QueryLikeService {
	return &QueryLikeService{ctx: ctx}
}
func (q *QueryLikeService) QueryLike(req *likerpc.QueryLikeByUserIdRequest) ([]*videorpc.Video, error) {
	// 查询用户点赞的所以视频 id
	videoIds, err := db.QueryLikeList(q.ctx, req.UserId)
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
