package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/hh02/minimal-douyin/cmd/user/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CreateUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

// CheckUser 用户服务，检查用户名密码是否匹配
func (s *CheckUserService) CheckUser(req *userrpc.CheckUserRequest) (int64, error) {
	// 先用 md5 对密码进行加密
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username
	user, err := db.QueryUserByName(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	// 如果用户不存在
	if user == nil {
		return 0, errno.UserNotExistErr
	}
	if user.Password != passWord {
		return 0, errno.LoginErr
	}
	return int64(user.ID), nil
}
