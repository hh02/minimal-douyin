package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/hh02/minimal-douyin/cmd/comment/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
	"gorm.io/gorm"
	"io"
)

type CreateCommentService struct {
	ctx context.Context
}

// NewCreateCommentService new CreateUserService
func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{ctx: ctx}
}

// CreateComment 用户服务,返回用户信息，错误
func (s *CreateCommentService) CreateComment(req *commentrpc.CreateCommentRequest) (*db.Comment, error) {

	return req.CommentText, nil
}
