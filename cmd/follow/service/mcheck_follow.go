package service

import (
	"context"

	"github.com/hh02/minimal-douyin/cmd/follow/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
)

type MCheckFollowService struct {
	ctx context.Context
}

func NewMCheckFollowService(ctx context.Context) *MCheckFollowService {
	return &MCheckFollowService{ctx: ctx}
}

func (s *MCheckFollowService) MCheckFollow(req *followrpc.MCheckFollowRequest) ([]bool, error) {

	followIds, err := db.QueryFollow(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	set := make(map[int64]struct{}, 0)
	for _, followId := range followIds {
		set[followId] = struct{}{}
	}
	res := make([]bool, 0)
	for _, followId := range req.FollowIds {
		if _, ok := set[followId]; ok {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res, nil
}
