package handlers

import (
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"net/http"
)

type UserInfoResponse struct {
	StatusCode    int32         `json:"status_code"`
	StatusMessage string        `json:"status_message"`
	User          *userrpc.User `json:"user"`
}

// SendUserInfoResponse pack UserInfoResponse
func SendUserInfoResponse(c *gin.Context, err error, user *userrpc.User) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode:    int32(Err.ErrCode),
		StatusMessage: Err.ErrMsg,
		User:          user,
	})
}

type RegisterResponse struct {
	StatusCode    int32  `json:"status_code"`
	StatusMessage string `json:"status_message"`
	UserId        int64  `json:"user_id"`
	Token         string `json:"token"`
}

// SendRegisterResponse pack RegisterResponse
func SendRegisterResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, RegisterResponse{
		StatusCode:    int32(Err.ErrCode),
		StatusMessage: Err.ErrMsg,
		UserId:        0,
		Token:         "",
	})
}

type UserParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
