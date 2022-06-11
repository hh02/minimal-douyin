package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

type VideoListResponse struct {
	StatusCode int32             `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string            `json:"status_msg"`  // 返回状态描述
	VideoList  []*videorpc.Video `json:"video_list"`  // 用户信息列表
}

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
	})
	c.Abort() // 注册失败，就停止后面中间件的运行
}

type UserParam struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type FeedResponse struct {
	StatusCode int32             `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string            `json:"status_msg"`  // 返回状态描述
	NextTime   int64             `json:"next_time"`
	VideoList  []*videorpc.Video `json:"video_list"` // 用户信息列表
}
