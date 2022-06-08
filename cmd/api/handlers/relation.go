package handlers

import (
	"context"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type UserListResponse struct {
	StatusCode int32           `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string          `json:"status_msg"`  // 返回状态描述
	UserList   []*userrpc.User `json:"user_list"`   // 用户信息列表
}

func RelationAction(c *gin.Context) {
	type relationParam struct {
		Token      string `form:"token" binding:"required"`
		ToUserId   int64  `form:"to_user_id" binding:"required"`
		ActionType uint8  `form:"action_type" binding:"required"`
	}

	var relationVar relationParam
	if err := c.ShouldBind(&relationVar); err != nil {
		SendStatusResponse(c, errno.ParamErr)
		return
	}

	if relationVar.ToUserId <= 0 || (relationVar.ActionType != 1 && relationVar.ActionType != 2) {
		SendStatusResponse(c, errno.ParamErr)
		return
	}

	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))
	// 1 for follow, 2 for unfollow
	if relationVar.ActionType == 1 {
		err := rpc.CreateFollow(context.Background(), &followrpc.CreateFollowRequest{
			UserId:   userId,
			FollowId: relationVar.ToUserId,
		})
		if err != nil {
			SendStatusResponse(c, errno.ConvertErr(err))
			return
		}
	} else if relationVar.ActionType == 2 {
		err := rpc.DeleteFollow(c, &followrpc.DeleteFollowRequest{
			UserId:   userId,
			FollowId: relationVar.ToUserId,
		})
		if err != nil {
			SendStatusResponse(c, errno.ConvertErr(err))
			return
		}
	}
	SendStatusResponse(c, errno.Success)
}

func FollowList(c *gin.Context) {

	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))

	users, err := rpc.QueryFollow(context.Background(), &followrpc.QueryFollowRequest{
		UserId: userId,
	})

	if err != nil {
		SendStatusResponse(c, errno.ConvertErr(err))
		return
	}

	c.JSON(http.StatusOK, UserListResponse{
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  errno.Success.ErrMsg,
		UserList:   users,
	})
}

func FollowerList(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))

	users, err := rpc.QueryFollower(context.Background(), &followrpc.QueryFollowerRequest{
		UserId: userId,
	})

	if err != nil {
		SendStatusResponse(c, errno.ConvertErr(err))
		return
	}

	c.JSON(http.StatusOK, UserListResponse{
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  errno.Success.ErrMsg,
		UserList:   users,
	})
}
