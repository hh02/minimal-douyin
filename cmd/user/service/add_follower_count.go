package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/user/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type AddFollowerCountService struct {
	ctx context.Context
}

// NewAddFollowerCountService new AddFollowCountService
func NewAddFollowerCountService(ctx context.Context) *AddFollowerCountService {
	return &AddFollowerCountService{ctx: ctx}
}

// AddFollowerCount 对指定用户增加粉丝数
func (s *AddFollowerCountService) AddFollowerCount(req *userrpc.AddFollowerCountRequest) error {
	// 查询用户是否存在
	user, err := db.QueryUserById(s.ctx, req.UserId)
	if err != nil {
		return err
	}
	// 用户不存在则返回 用户不存在（UserNotExistErr）的错误
	if user == nil {
		return errno.UserNotExistErr
	}
	return db.AddFollowerCount(s.ctx, req.UserId, req.Count)
}
