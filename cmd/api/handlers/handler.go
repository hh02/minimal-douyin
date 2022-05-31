package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/pkg/errno"
	"net/http"
)

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type UserInfoParam struct {
	UserId int64 `json:"userid"`
	Token  string `json:"token"`
}

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}
