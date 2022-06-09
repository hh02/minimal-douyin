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
	videoIds, err := db.QueryLike(q.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if len(videoIds) == 0 {
		return nil, nil
	}
	rpcRequest := &videorpc.MGetVideoRequest{
		VideoIds: videoIds,
	}
	videoID, err := rpc.MGetUser(q.ctx, rpcRequest)
	if err != nil {
		return nil, err
	}
	return videoID, nil
}
