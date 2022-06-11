package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/cmd/api/utils"
	"github.com/hh02/minimal-douyin/kitex_gen/response"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
)

func SendFeedResponse(c *gin.Context, err error, nextTime int64, videos []*videorpc.Video) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.FeedResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		NextTime:   nextTime,
		VideoList:  videos,
	})
	// 防止继续运行鉴权中间件
	c.Abort()
}

func Feed(c *gin.Context) {
	latestTimeStr := c.DefaultQuery("latest_time", "")
	var (
		latestTime int64
		err        error
	)
	if len(latestTimeStr) > 0 {
		latestTime, err = strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			SendFeedResponse(c, errno.ParamErr, 0, nil)
			return
		}
	} else {
		latestTime = time.Now().Unix()
	}

	tokenUserId := utils.GetIdFromClaims(c)

	// tokenUserId 如果为0，表示不查询关注信息
	nextTime, videos, err := rpc.QueryVideoByTime(context.Background(), &videorpc.QueryVideoByTimeRequest{
		LatestTime:  latestTime,
		TokenUserId: tokenUserId,
	})
	if err != nil {
		SendFeedResponse(c, err, 0, nil)
		return
	}
	SendFeedResponse(c, errno.Success, nextTime, videos)
}
