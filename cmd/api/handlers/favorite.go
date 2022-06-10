package handlers

import (
	"context"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/favoriterpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

func FavoriteAction(c *gin.Context) {
	type formParam struct {
		Token      string `form:"token" binding:"required"`
		VideoId    int64  `form:"video_id" binding:"required"`
		ActionType int32  `form:"action_type" binding:"required"`
	}

	var param formParam
	if err := c.ShouldBind(&param); err != nil {
		SendStatusResponse(c, errno.ParamErr)
		return
	}

	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))

	if param.ActionType == 1 {
		err := rpc.CreateFavorite(context.Background(), &favoriterpc.CreateFavoriteRequest{
			UserId:  userId,
			VideoId: param.VideoId,
		})
		if err != nil {
			SendStatusResponse(c, errno.ConvertErr(err))
			return
		}
	} else if param.ActionType == 2 {
		err := rpc.DeleteFavorite(c, &favoriterpc.DeleteFavoriteRequest{
			UserId:  userId,
			VideoId: param.VideoId,
		})
		if err != nil {
			SendStatusResponse(c, errno.ConvertErr(err))
		}
	}
	SendStatusResponse(c, errno.Success)
}

func FavoriteList(c *gin.Context) {
	type formParam struct {
		UserId int64  `form:"user_id" binding:"required"`
		Token  string `form:"token" binding:"required"`
	}

	var param formParam
	if err := c.ShouldBind(&param); err != nil {
		SendVideoListResponse(c, err, nil)
		return
	}

	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))

	videos, err := rpc.QueryFavorite(context.Background(), &favoriterpc.QueryFavoriteByUserIdRequest{
		UserId: userId,
	})
	if err != nil {
		SendVideoListResponse(c, err, nil)
		return
	}
	SendVideoListResponse(c, errno.Success, videos)
}
