package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/comment/dal/db"
	"github.com/hh02/minimal-douyin/cmd/comment/pack"
	"github.com/hh02/minimal-douyin/cmd/comment/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
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

	comment := &db.Comment{
		UserId:  req.UserId,
		VideoId: req.VideoId,
		Content: req.CommentText,
	}

	err := db.CreateComment(s.ctx, comment)
	if err != nil {
		return nil, err
	}
	err = rpc.AddCommentCount(s.ctx, &videorpc.AddCommentCountRequest{
		VideoId: req.VideoId,
		Count:   1,
	})
	if err != nil {
		return nil, err
	}
	return pack.Comment(s.ctx, comment), nil
}
