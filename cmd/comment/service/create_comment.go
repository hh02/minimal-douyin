package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/comment/dal/db"
	"github.com/hh02/minimal-douyin/cmd/comment/pack"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
)

type CreateCommentService struct {
	ctx context.Context
}

// NewCreateCommentService new CreateCommentService
func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{ctx: ctx}
}

// CreateComment 用户服务,返回用户信息，错误
func (s *CreateCommentService) CreateComment(req *commentrpc.CreateCommentRequest) (*commentrpc.Comment, error) {
	err := db.CreateComment(s.ctx, &db.Comment{
		UserId:  req.UserId,
		VideoId: req.VideoId,
		Content: req.CommentText,
	})
	if err != nil {
		return nil, err
	}
	return pack.Comment(s.ctx, &db.Comment{
		UserId:  req.UserId,
		VideoId: req.VideoId,
		Content: req.CommentText,
	}), nil
}
