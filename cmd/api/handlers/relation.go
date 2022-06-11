package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/cmd/api/utils"
	"github.com/hh02/minimal-douyin/kitex_gen/followrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/response"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type UserListResponse struct {
	StatusCode int32           `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string          `json:"status_msg"`  // 返回状态描述
	UserList   []*userrpc.User `json:"user_list"`   // 用户信息列表
}

func SendRelationActionResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.RelationActionResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}

func SendRelationFollowListResponse(c *gin.Context, err error, users []*userrpc.User) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.RelationFollowListResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserList:   users,
	})
}

func SendRelationFollowerListResponse(c *gin.Context, err error, users []*userrpc.User) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.RelationFollowerListResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserList:   users,
	})
}

func RelationAction(c *gin.Context) {
	if err := utils.CheckAuth(c); err != nil {
		SendRelationActionResponse(c, err)
		return
	}
	type RelationParam struct {
		Token      string `form:"token" binding:"required"`
		ToUserId   int64  `form:"to_user_id" binding:"required"`
		ActionType uint8  `form:"action_type" binding:"required"`
	}

	var relationVar RelationParam
	if err := c.ShouldBind(&relationVar); err != nil {
		SendRelationActionResponse(c, err)
		return
	}
	if relationVar.ToUserId <= 0 || (relationVar.ActionType != 1 && relationVar.ActionType != 2) {
		SendRelationActionResponse(c, errno.ParamErr)
		return
	}

	tokenId := utils.GetIdFromClaims(c)
	if tokenId == 0 {
		SendUserResponse(c, errno.AuthErr, nil)
		return
	}

	// 1 for follow, 2 for unfollow
	if relationVar.ActionType == 1 {
		err := rpc.CreateFollow(context.Background(), &followrpc.CreateFollowRequest{
			UserId:   tokenId,
			FollowId: relationVar.ToUserId,
		})
		if err != nil {
			SendRelationActionResponse(c, err)
			return
		}
	} else if relationVar.ActionType == 2 {
		err := rpc.DeleteFollow(context.Background(), &followrpc.DeleteFollowRequest{
			UserId:   tokenId,
			FollowId: relationVar.ToUserId,
		})
		if err != nil {
			SendRelationActionResponse(c, err)
			return
		}
	}
	SendRelationActionResponse(c, errno.Success)
}

func FollowList(c *gin.Context) {
	if err := utils.CheckAuth(c); err != nil {
		SendRelationFollowListResponse(c, err, nil)
		return
	}
	type FollowListParam struct {
		UserId int64  `form:"user_id"`
		Token  string `form:"token"`
	}

	var followListVar FollowListParam
	if err := c.ShouldBindQuery(&followListVar); err != nil {
		SendRelationFollowListResponse(c, err, nil)
		return
	}

	fmt.Println(followListVar.UserId)
	users, err := rpc.QueryFollow(context.Background(), &followrpc.QueryFollowRequest{
		UserId: followListVar.UserId,
	})

	if err != nil {
		SendRelationFollowListResponse(c, err, nil)
		return
	}

	SendRelationFollowListResponse(c, errno.Success, users)
}

func FollowerList(c *gin.Context) {
	if err := utils.CheckAuth(c); err != nil {
		SendRelationFollowerListResponse(c, err, nil)
		return
	}
	type FollowerListParam struct {
		UserId int64  `form:"user_id"`
		Token  string `form:"token"`
	}

	var followerListVar FollowerListParam
	if err := c.BindQuery(&followerListVar); err != nil {
		SendRelationFollowerListResponse(c, err, nil)
		return
	}

	users, err := rpc.QueryFollower(context.Background(), &followrpc.QueryFollowerRequest{
		UserId: followerListVar.UserId,
	})
	if err != nil {
		SendRelationFollowerListResponse(c, err, nil)
		return
	}

	SendRelationFollowerListResponse(c, errno.Success, users)
}
