package handlers

import (
	"context"

	"net/http"

	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/cmd/api/utils"
	"github.com/hh02/minimal-douyin/kitex_gen/response"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

type UserParam struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AuthResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

func SendUserLoginResponse(c *gin.Context, err error, userId int64, token string) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.UserLoginResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserId:     userId,
		Token:      token,
	})
}

func SendUserResponse(c *gin.Context, err error, user *userrpc.User) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.UserResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		User:       user,
	})
}

func SendUserRegisterResponse(c *gin.Context, err error, userId int64, token string) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.UserLoginResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserId:     userId,
		Token:      token,
	})
}

// UserInfo query user info
func UserInfo(c *gin.Context) {
	if err := utils.CheckAuth(c); err != nil {
		SendUserResponse(c, err, nil)
		return
	}
	tokenID := utils.GetIdFromClaims(c)
	if tokenID == 0 {
		SendUserResponse(c, errno.AuthErr, nil)
		return
	}
	var queryVar struct {
		UserId int64 `json:"user_id" form:"user_id"`
	}
	if err := c.BindQuery(&queryVar); err != nil {
		SendUserResponse(c, err, nil)
		return
	}

	if queryVar.UserId <= 0 {
		SendUserResponse(c, errno.ParamErr, nil)
		return
	}

	req := &userrpc.GetUserRequest{
		UserId:         queryVar.UserId,
		TokenUserId:    tokenID,
		ReturnIsFollow: true,
	}

	user, err := rpc.GetUser(context.Background(), req)
	if err != nil {
		SendUserResponse(c, err, nil)
		return
	}
	SendUserResponse(c, errno.Success, user)
}

func UserRegister(c *gin.Context) {
	var param UserParam
	if err := c.ShouldBind(&param); err != nil {
		SendUserRegisterResponse(c, err, 0, "")
		return
	}

	userId, err := rpc.CreateUser(context.Background(), &userrpc.CreateUserRequest{
		Username: param.Username,
		Password: param.Password,
	})
	if err != nil {
		SendUserRegisterResponse(c, err, 0, "")
		return
	}

	tokenString, err := utils.GenerateTokenString(userId)
	if err != nil {
		SendUserRegisterResponse(c, err, 0, "")
		return
	}
	SendUserRegisterResponse(c, errno.Success, userId, tokenString)
}

func UserLogin(c *gin.Context) {

	var param UserParam
	if err := c.ShouldBind(&param); err != nil {
		SendUserLoginResponse(c, err, 0, "")
	}

	userId, err := rpc.CheckUser(context.Background(), &userrpc.CheckUserRequest{
		Username: param.Username,
		Password: param.Password,
	})
	if err != nil {
		SendUserLoginResponse(c, err, 0, "")
		return
	}

	tokenString, err := utils.GenerateTokenString(userId)
	if err != nil {
		SendUserLoginResponse(c, err, 0, "")
		return
	}

	SendUserLoginResponse(c, errno.Success, userId, tokenString)
}
