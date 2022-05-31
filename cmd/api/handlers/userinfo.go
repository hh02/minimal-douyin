package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/douyin/core"
	"github.com/hh02/minimal-douyin/pkg/errno"
	"net/http"
)

func UserInfo(c *gin.Context) {
	var userVar UserInfoParam
	if err := c.ShouldBind(&userVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if userVar.UserId == 0 || len(userVar.Token) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	user, err := rpc.UsreInfo(context.Background(), &core.UserRequest{
		UserId: userVar.UserId,
		Token:  userVar.Token,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(http.StatusOK, core.UserResponse{
		StatusCode: errno.Success.ErrCode,
		StatusMsg: errno.Success.ErrMsg,
		User:     user,
	})
}
