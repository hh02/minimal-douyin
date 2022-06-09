package handlers

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/likerpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type VideoListResponse struct {
	StatusCode int32             `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
	VideoList  []*videorpc.Video `json:"video_list"` //点赞视频列表
}

func LikeAction(c *gin.Context) {
	type likeParam struct {
		Token      string `form:"token" binding:"required"`
		VideoId    int64  `form:"video_id" binding:"required"`
		ActionType uint8  `form:"action_type" binding:"required"`
	}
	var likeVar likeParam
	if err := c.ShouldBind(&likeVar); err != nil {
		SendStatusResponse(c, errno.ParamErr)
	}
	if likeVar.VideoId <= 0 || (likeVar.ActionType != 1 && likeVar.ActionType != 2) {
		SendStatusResponse(c, errno.ParamErr)
	}
	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))
	if likeVar.ActionType == 1 {
		err := rpc.CreateLike(context.Background(), &likerpc.CreateLikeRequest{
			UserId:  userId,
			VideoId: likeVar.VideoId,
		})
		if err != nil {
			SendStatusResponse(c, errno.ParamErr)
			return
		}

	} else if likeVar.ActionType == 2 {
		err := rpc.DeleteLike(context.Background(), &likerpc.DeleteLikeRequest{
			UserId:  userId,
			VideoId: likeVar.VideoId,
		})
		if err != nil {
			SendStatusResponse(c, errno.ParamErr)
			return
		}
	}
	SendStatusResponse(c, errno.ParamErr)
}
func QueryLike(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))
	videos, err := rpc.QueryLike(context.Background(), &likerpc.QueryLikeByUserIdRequest{
		UserId: userId,
	})
	if err != nil {
		SendStatusResponse(c, errno.ParamErr)
		return
	}
	c.JSON(200, VideoListResponse{
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  errno.Success.ErrMsg,
		VideoList:  videos,
	})
}
