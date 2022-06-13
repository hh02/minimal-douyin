package main

import (
	"context"
	"fmt"

	"github.com/hh02/minimal-douyin/cmd/comment/service"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CreateComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateComment(ctx context.Context, req *commentrpc.CreateCommentRequest) (resp *commentrpc.CreateCommentResponse, err error) {
	// TODO: Your code here...
	resp = new(commentrpc.CreateCommentResponse)
	fmt.Println(req.UserId, req.VideoId, req.CommentText)
	if req.UserId == 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	comment, err := service.NewCreateCommentService(ctx).CreateComment(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.Comment = comment
	return resp, nil
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *commentrpc.DeleteCommentRequest) (resp *commentrpc.DeleteCommentResponse, err error) {
	// TODO: Your code here...
	resp = new(commentrpc.DeleteCommentResponse)

	if req.UserId == 0 || req.CommentId == 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	err = service.NewDeleteCommentService(ctx).DeleteComment(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	return resp, nil
}

// QueryCommentByVideo implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) QueryCommentByVideo(ctx context.Context, req *commentrpc.QueryCommentByVideoIdRequest) (resp *commentrpc.QueryCommentByVideoIdResponse, err error) {
	// TODO: Your code here...
	resp = new(commentrpc.QueryCommentByVideoIdResponse)

	if req.VideoId == 0 {
		resp.Status = errno.BuildStatus(errno.ParamErr)
		return resp, nil
	}

	comments, err := service.NewQueryCommentService(ctx).QueryComment(req)
	if err != nil {
		resp.Status = errno.BuildStatus(err)
		return resp, nil
	}
	resp.Status = errno.BuildStatus(errno.Success)
	resp.CommentList = comments
	return resp, nil
}
