package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/cmd/api/utils"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/response"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

func SendCommentActionResponse(c *gin.Context, err error, comment *commentrpc.Comment) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.CommentActionResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		Comment:    comment,
	})
}

func SendCommentListResponse(c *gin.Context, err error, comments []*commentrpc.Comment) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.CommentListResponse{
		StatusCode:  Err.ErrCode,
		StatusMsg:   Err.ErrMsg,
		CommentList: comments,
	})
}

// CommentAction create comment or delete comment
func CommentAction(c *gin.Context) {
	if err := utils.CheckAuth(c); err != nil {
		SendCommentActionResponse(c, err, nil)
		return
	}
	type CommentParam struct {
		Token       string `form:"token"`
		VideoId     int64  `form:"video_id"`
		ActionType  uint8  `form:"action_type"`
		CommentText string `form:"comment_text"`
		CommentId   int64  `form:"comment_id"`
	}

	var commentVar CommentParam
	if err := c.ShouldBind(&commentVar); err != nil {
		SendCommentActionResponse(c, err, nil)
		return
	}

	tokenId := utils.GetIdFromClaims(c)
	if tokenId == 0 {
		SendCommentActionResponse(c, errno.AuthErr, nil)
		return
	}

	fmt.Println(commentVar)
	// 1 create comment ; 2 delete comment
	if commentVar.ActionType == 1 {
		req := &commentrpc.CreateCommentRequest{
			UserId:      tokenId,
			VideoId:     commentVar.VideoId,
			CommentText: commentVar.CommentText,
		}
		comment, err := rpc.CreateComment(context.Background(), req)
		if err != nil {
			SendCommentActionResponse(c, err, nil)
			return
		}
		SendCommentActionResponse(c, errno.Success, comment)
	} else if commentVar.ActionType == 2 {
		req := &commentrpc.DeleteCommentRequest{
			UserId:    tokenId,
			CommentId: commentVar.CommentId,
		}
		err := rpc.DeleteComment(context.Background(), req)
		if err != nil {
			SendCommentActionResponse(c, err, nil)
			return
		}
		SendCommentActionResponse(c, errno.Success, nil)

	} else {
		SendCommentActionResponse(c, errno.ParamErr, nil)
	}
}

// CommentList query comment by video id
func CommentList(c *gin.Context) {
	if err := utils.CheckAuth(c); err != nil {
		SendCommentListResponse(c, err, nil)
		return
	}
	type CommentParam struct {
		Token   string `form:"token"`
		VideoId int64  `form:"video_id"`
	}

	var commentVar CommentParam
	if err := c.ShouldBind(&commentVar); err != nil {
		SendCommentListResponse(c, err, nil)
		return
	}

	tokenId := utils.GetIdFromClaims(c)
	if tokenId == 0 {
		SendCommentListResponse(c, errno.AuthErr, nil)
		return
	}

	req := &commentrpc.QueryCommentByVideoIdRequest{
		VideoId:     commentVar.VideoId,
		TokenUserId: tokenId,
	}
	comments, err := rpc.QueryCommentByVideo(context.Background(), req)
	if err != nil {
		SendCommentListResponse(c, err, nil)
		return
	}

	SendCommentListResponse(c, errno.Success, comments)
}
