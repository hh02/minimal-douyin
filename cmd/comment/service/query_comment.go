package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/comment/dal/db"
	"github.com/hh02/minimal-douyin/cmd/comment/pack"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
)

type QueryCommentService struct {
	ctx context.Context
}

// NewQueryCommentService new QueryCommentService
func NewQueryCommentService(ctx context.Context) *QueryCommentService {
	return &QueryCommentService{ctx: ctx}
}

// QueryComment 用户服务,返回用户信息，错误
func (s *QueryCommentService) QueryComment(req *commentrpc.QueryCommentByVideoIdRequest) ([]*commentrpc.Comment, error) {
	comments, err := db.QueryCommentByVideoId(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	return pack.MComment(s.ctx, comments, req.TokenUserId)
}
