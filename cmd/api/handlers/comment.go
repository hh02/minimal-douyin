package handlers

import (
	"context"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/commentrpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
	"net/http"
)

type CommentActionResponse struct {
	StatusCode int32               `json:"status_code"`
	StatusMsg  string              `json:"status_msg"`
	Comment    *commentrpc.Comment `json:"comment"`
}

type CommentListResponse struct {
	StatusCode  int32                 `json:"status_code"`
	StatusMsg   string                `json:"status_msg"`
	CommentList []*commentrpc.Comment `json:"comment_list"`
}

// CommentAction create comment or delete comment
func CommentAction(c *gin.Context) {
	type CommentParam struct {
		Token       string `form:"token"`
		VideoId     int64  `form:"video_id"`
		ActionType  uint8  `form:"action_type"`
		CommentText string `form:"comment_text"`
		CommentId   int64  `form:"comment_id"`
	}

	var commentVar CommentParam
	if err := c.ShouldBind(&commentVar); err != nil {
		SendStatusResponse(c, err)
		return
	}
	clams := jwt.ExtractClaims(c)
	tokenId := int64(clams[constants.IdentityKey].(float64))

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
			SendStatusResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{
			StatusCode: errno.Success.ErrCode,
			StatusMsg:  errno.Success.ErrMsg,
			Comment:    comment,
		})
	} else if commentVar.ActionType == 2 {
		req := &commentrpc.DeleteCommentRequest{
			UserId:    tokenId,
			CommentId: commentVar.CommentId,
		}
		err := rpc.DeleteComment(context.Background(), req)
		if err != nil {
			SendStatusResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{
			StatusCode: errno.Success.ErrCode,
			StatusMsg:  errno.Success.ErrMsg,
			Comment:    nil,
		})
	} else {
		SendStatusResponse(c, errno.ParamErr)
	}
}

// CommentList query comment by video id
func CommentList(c *gin.Context) {
	type CommentParam struct {
		Token   string `form:"token"`
		VideoId int64  `form:"video_id"`
	}

	var commentVar CommentParam
	if err := c.ShouldBind(&commentVar); err != nil {
		SendStatusResponse(c, err)
		return
	}

	clams := jwt.ExtractClaims(c)
	tokenId := int64(clams[constants.IdentityKey].(float64))

	req := &commentrpc.QueryCommentByVideoIdRequest{
		VideoId:     commentVar.VideoId,
		TokenUserId: tokenId,
	}
	fmt.Println("???")
	comments, err := rpc.QueryCommentByVideo(context.Background(), req)
	if err != nil {
		SendStatusResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, &CommentListResponse{
		StatusCode:  errno.Success.ErrCode,
		StatusMsg:   errno.Success.ErrMsg,
		CommentList: comments,
	})
}
