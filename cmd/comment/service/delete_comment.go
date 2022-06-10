package service

import (
	"context"
	"github.com/hh02/minimal-douyin/cmd/comment/dal/db"
	"github.com/hh02/minimal-douyin/cmd/comment/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type DeleteCommentService struct {
	ctx context.Context
}

// NewDeleteCommentService new DeleteCommentService
func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{ctx: ctx}
}

// DeleteComment 用户服务,返回用户信息，错误
func (s *DeleteCommentService) DeleteComment(req *commentrpc.DeleteCommentRequest) error {

	// 查询是否有这条评论
	comment, err := db.QueryCommentByCommentId(s.ctx, req.CommentId)
	if err != nil {
		return err
	}

	// 判断用户有没有权限删除改评论
	if comment.UserId != req.UserId {
		return errno.PermissionErr
	}

	err = db.DeleteComment(s.ctx, req.CommentId)
	if err != nil {
		return err
	}
	err = rpc.AddCommentCount(s.ctx, &videorpc.AddCommentCountRequest{
		VideoId: comment.VideoId,
		Count:   -1,
	})
	if err != nil {
		return err
	}
	return nil
}
