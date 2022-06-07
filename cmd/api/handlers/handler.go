package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type UserInfoResponse struct {
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg"`
	User       *userrpc.User `json:"user"`
}

// SendUserInfoResponse pack UserInfoResponse
func SendUserInfoResponse(c *gin.Context, err error, user *userrpc.User) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode: int32(Err.ErrCode),
		StatusMsg:  Err.ErrMsg,
		User:       user,
	})
}

type RegisterResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

// SendRegisterResponse pack RegisterResponse
func SendRegisterResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, RegisterResponse{
		StatusCode: int32(Err.ErrCode),
		StatusMsg:  string(Err.ErrMsg),
		UserId:     0,
		Token:      "",
	})
}

type UserParam struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type StatusResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func SendStatusResponse(c *gin.Context, err error) {
	status := errno.BuildStatus(err)
	c.JSON(http.StatusOK, StatusResponse{
		StatusCode: status.StatusCode,
		StatusMsg:  status.StatusMessage,
	})
}