package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/user/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type AddFollowCountService struct {
	ctx context.Context
}

// NewAddFollowCountService new AddFollowCountService
func NewAddFollowCountService(ctx context.Context) *AddFollowCountService {
	return &AddFollowCountService{ctx: ctx}
}

// AddFollowCount 对指定用户增加关注数
func (s *AddFollowCountService) AddFollowCount(req *userrpc.AddFollowCountRequest) error {
	// 查询用户是否存在
	user, err := db.QueryUserById(s.ctx, req.UserId)
	if err != nil {
		return err
	}
	// 用户不存在则返回 用户不存在（UserNotExistErr）的错误
	if user == nil {
		return errno.UserNotExistErr
	}
	return db.AddFollowCount(s.ctx, req.UserId, req.Count)
}
