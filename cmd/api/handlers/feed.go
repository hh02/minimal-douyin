package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

func Feed(c *gin.Context) {

	type feedResponse struct {
		StatusCode int32             `json:"status_code"` // 状态码，0-成功，其他值-失败
		StatusMsg  string            `json:"status_msg"`  // 返回状态描述
		NextTime   int64             `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
		VideoList  []*videorpc.Video `json:"video_list"`  // 视频列表

	}

	latestTimeStr := c.Query("latest_time")
	token := c.Query("token")
	var (
		latestTime  int64
		tokenUserId int64
		err         error
	)
	if len(latestTimeStr) > 0 {
		latestTime, err = strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			SendStatusResponse(c, errno.ParamErr.WithMessage("时间参数错误"))
			return
		}
	} else {
		latestTime = time.Now().Unix()
	}

	if len(token) > 0 {
		claims := jwt.ExtractClaims(c)
		tokenUserId = int64(claims[constants.IdentityKey].(float64))
	} else {
		tokenUserId = 0
	}

	nextTime, videos, err := rpc.QueryVideoByTime(context.Background(), &videorpc.QueryVideoByTimeRequest{
		LatestTime:  latestTime,
		TokenUserId: tokenUserId,
	})
	if err != nil {
		SendStatusResponse(c, errno.ConvertErr(err))
		return
	}

	c.JSON(http.StatusOK, feedResponse{
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  errno.Success.ErrMsg,
		NextTime:   nextTime,
		VideoList:  videos,
	})

}
