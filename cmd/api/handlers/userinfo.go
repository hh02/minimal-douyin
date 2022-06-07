package handlers

import (
	"context"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

// UserInfo query user info
func UserInfo(c *gin.Context) {
	claims := c.Query("token") //c.Get("token") //

	//测试
	fmt.Println(claims)
	fmt.Println(jwt.ExtractClaims(c))
	//fmt.Println(base64.RawURLEncoding.DecodeString(claims))

	tokenID := int64(0) //int64(claims[constants.IdentityKey].(float64))
	var queryVar struct {
		UserId int64 `json:"user_id" form:"user_id"`
	}
	if err := c.BindQuery(&queryVar); err != nil {
		SendUserInfoResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if queryVar.UserId == 0 {
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
