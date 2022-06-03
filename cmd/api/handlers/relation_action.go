package handlers

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

func RelationAction(c *gin.Context) {
	type RelationParam struct {
		Token string `json:"token"`
		ToUserId int64 `json:"to_user_id"`
		ActionType uint8 `json:"action_type"`
	}

	var relationVar RelationParam
	if err := c.ShouldBind(&relationVar); err != nil {
		SendStatusResponse(c, err)
	}

	if relationVar.ToUserId <= 0 || (relationVar.ActionType != 1 && relationVar.ActionType != 2) {
		SendStatusResponse(c, errno.ParamErr)
	}

	clams := jwt.ExtractClaims(c)
	userId := int64(clams[constants.IdentityKey].(float64))

	// 1 for follow, 2 for unfollow
	if relationVar.ActionType == 1 {
		err := rpc.CreateFollow(context.Background(), &followrpc.CreateFollowRequest{
			UserId: userId,
			FollowId: relationVar.ToUserId,
		})
		if err != nil {
			SendStatusResponse(c, errno.ConvertErr(err))
			return
		}
		SendStatusResponse(c, errno.Success)
	} else if relationVar.ActionType == 2 {
		err := rpc.DeleteFollow(c, &followrpc.DeleteFollowRequest{
			UserId: userId,
			FollowId: relationVar.ToUserId,
		})
		if err != nil {
			SendStatusResponse(c, errno.ConvertErr(err))
			return
		}
	}
	SendStatusResponse(c, errno.Success)
}