package handlers

import (
	"context"

	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

type AuthResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

func UserLoginResponse(c *gin.Context) {

	claims := jwt.ExtractClaims(c)
	tokenId := int64(claims[constants.IdentityKey].(float64))
	tokenString := jwt.GetToken(c)

	c.JSON(http.StatusOK, AuthResponse{
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  errno.Success.ErrMsg,
		UserId:     tokenId,
		Token:      tokenString,
	})

}

// UserInfo query user info
func UserInfo(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	tokenID := int64(claims[constants.IdentityKey].(float64))
	var queryVar struct {
		UserId int64 `json:"user_id" form:"user_id"`
	}
	if err := c.BindQuery(&queryVar); err != nil {
		SendUserInfoResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if queryVar.UserId <= 0 {
		SendUserInfoResponse(c, errno.ParamErr, nil)
		return
	}

	req := &userrpc.GetUserRequest{
		UserId:         queryVar.UserId,
		TokenUserId:    tokenID,
		ReturnIsFollow: true,
	}

	user, err := rpc.GetUser(context.Background(), req)
	if err != nil {
		SendUserInfoResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendUserInfoResponse(c, errno.Success, user)
}

func UserRegister(c *gin.Context) {
	var userVar UserParam
	if err := c.ShouldBind(&userVar); err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err))
		return
	}
	if len(userVar.Username) == 0 || len(userVar.Password) == 0 {
		SendRegisterResponse(c, errno.ParamErr)
		return
	}

	err := rpc.CreateUser(context.Background(), &userrpc.CreateUserRequest{
		Username: userVar.Username,
		Password: userVar.Password,
	})

	if err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err))
		return
	}
	c.Next()
}
