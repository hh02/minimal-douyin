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
		Token string `form:"token" binding:"required"`
		VideoId int64 `form:"video_id" binding:"required"`
		ActionType int32 `form:"action_type" binding:"required"`
	}

	var param formParam
	if err := c.ShouldBind(&param); err != nil {
		SendStatusResponse(c, errno.ParamErr)
		return
	}

	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))

	if param.ActionType == 1 {
		rpc.CreateFavorite(context.Background(), &favoriterpc.CreateFavoriteRequest{
			UserId: userId,
	
		})

	}





}