package handlers

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/errno"
	"google.golang.org/protobuf/proto"
)

type PbJSON struct {
	Data interface{}
}

func WritePbJSON(w http.ResponseWriter, obj *) error {
	writecontenttype()
	var buffer bytes.Buffer
	jsonpbMarshaler.Marshal(&buffer, &obj)
}
func (r PbJSON) Render(w http.ResponseWriter) (err error) {
	if err = WritePbJSON(w, r.Data); err != nil {
		panic(err)
	}
	return
}
var jsonpbMarshaler = &jsonpb.Marshaler {
	EnumsAsInts: true,
	EmitDefaults: true,
}

type VideoListResponse struct {
	StatusCode int32             `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string            `json:"status_msg"`  // 返回状态描述
	VideoList  []*videorpc.Video `json:"video_list"`  // 用户信息列表
}

func SendVideoListResponse(c *gin.Context, err error, videos []*videorpc.Video) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, VideoListResponse{
		StatusCode: int32(Err.ErrCode),
		StatusMsg:  Err.ErrMsg,
		VideoList:  videos,
	})
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

type FeedResponse struct {
	StatusCode int32             `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string            `json:"status_msg"`  // 返回状态描述
	NextTime   int64             `json:"next_time"`
	VideoList  []*videorpc.Video `json:"video_list"` // 用户信息列表
}

func SendFeedResponse(c *gin.Context, err error, nextTime int64, videos []*videorpc.Video) {
	status := errno.BuildStatus(err)
	c.JSON(http.StatusOK, FeedResponse{
		StatusCode: status.StatusCode,
		StatusMsg:  status.StatusMessage,
		NextTime:   nextTime,
		VideoList:  videos,
	})
	// 防止继续运行鉴权中间件
	c.Abort()
}
