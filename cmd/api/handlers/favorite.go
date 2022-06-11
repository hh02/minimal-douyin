package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/cmd/api/utils"
	"github.com/hh02/minimal-douyin/kitex_gen/favoriterpc"
	"github.com/hh02/minimal-douyin/kitex_gen/response"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

func SendFavoriteActionResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.FavoriteActionResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}

func SendFavoriteListResponse(c *gin.Context, err error, videos []*videorpc.Video) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.FavoriteListResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		VideoList:  videos,
	})
}

func FavoriteAction(c *gin.Context) {
	if err := utils.CheckAuth(c); err != nil {
		SendFavoriteActionResponse(c, err)
	}

	type formParam struct {
		Token      string `form:"token" binding:"required"`
		VideoId    int64  `form:"video_id" binding:"required"`
		ActionType int32  `form:"action_type" binding:"required"`
	}

	var param formParam
	if err := c.ShouldBind(&param); err != nil {
		SendFavoriteActionResponse(c, err)
		return
	}

	userId := utils.GetIdFromClaims(c)
	if userId == 0 {
		SendFavoriteActionResponse(c, errno.GetIdFromClaimsErr)
		return
	}

	if param.ActionType == 1 {
		err := rpc.CreateFavorite(context.Background(), &favoriterpc.CreateFavoriteRequest{
			UserId:  userId,
			VideoId: param.VideoId,
		})
		if err != nil {
			SendFavoriteActionResponse(c, err)
			return
		}
	} else if param.ActionType == 2 {
		err := rpc.DeleteFavorite(c, &favoriterpc.DeleteFavoriteRequest{
			UserId:  userId,
			VideoId: param.VideoId,
		})
		if err != nil {
			SendFavoriteActionResponse(c, err)
			return
		}
	} else {
		SendFavoriteActionResponse(c, errno.ParamErr)
		return
	}
	SendFavoriteActionResponse(c, errno.Success)
}

func FavoriteList(c *gin.Context) {
	if err := utils.CheckAuth(c); err != nil {
		SendFavoriteListResponse(c, err, nil)
		return
	}
	type formParam struct {
		UserId int64  `form:"user_id" binding:"required"`
		Token  string `form:"token" binding:"required"`
	}

	var param formParam
	if err := c.ShouldBind(&param); err != nil {
		SendFavoriteListResponse(c, err, nil)
		return
	}

	userId := utils.GetIdFromClaims(c)
	if userId == 0 {
		SendFavoriteListResponse(c, errno.GetIdFromClaimsErr, nil)
		return
	}

	videos, err := rpc.QueryFavorite(context.Background(), &favoriterpc.QueryFavoriteByUserIdRequest{
		UserId: userId,
	})
	if err != nil {
		SendFavoriteListResponse(c, err, nil)
		return
	}
	SendFavoriteListResponse(c, errno.Success, videos)
}
