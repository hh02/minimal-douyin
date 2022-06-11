package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/cmd/api/utils"
	"github.com/hh02/minimal-douyin/kitex_gen/response"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
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
	fmt.Println("latest time is: ", latestTimeStr)
	token := c.DefaultQuery("token", "")
	var (
		latestTime  int64
		tokenUserId int64
		err         error
	)
	if len(latestTimeStr) > 0 {
		latestTime, err = strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			SendFeedResponse(c, errno.ParamErr.WithMessage("时间参数错误"), 0, nil)
			return
		}
	} else {
		latestTime = time.Now().Unix()
	}

	if len(token) > 0 {
		// 调用鉴权服务
		c.Next()
		// 如果失败就返回
		if c.IsAborted() {
			SendFeedResponse(c, errno.AuthErr, 0, nil)
			return
		}
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
		SendFeedResponse(c, err, 0, nil)
		return
	}
	SendFeedResponse(c, errno.Success, nextTime, videos)

}
