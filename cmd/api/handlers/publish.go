package handlers

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
	uuid "github.com/satori/go.uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// 获取视频封面
func getSnapshot(videoPath, snapshotPath string, frameNum int) (err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		return err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		return err
	}
	return nil
}

func PublishAction(c *gin.Context) {
	type formParam struct {
		Data  *multipart.FileHeader `form:"data" binding:"required"`
		Token string                `form:"token" binding:"required"`
		Title string                `form:"title" binding:"required"`
	}

	var formParamVar formParam
	if err := c.ShouldBind(&formParamVar); err != nil {
		SendStatusResponse(c, errno.ParamErr)
		return
	}

	if len(formParamVar.Title) > 100 {
		SendStatusResponse(c, errno.ParamErr.WithMessage("标题长度超过100"))
		return
	}

	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))

	videoName := uuid.NewV4().String()
	videoPath := constants.VideoFolder + videoName + ".mp4"
	coverPath := constants.CoverFolder + videoName + ".png"

	c.SaveUploadedFile(formParamVar.Data, videoPath)
	err := getSnapshot(videoPath, coverPath, 1)
	if err != nil {
		SendStatusResponse(c, errno.ConvertErr(err).WithMessage("获取封面失败"))
		return
	}

	playUrl := constants.FileServer + videoPath
	coverUrl := constants.FileServer + coverPath

	err = rpc.CreateVideo(context.Background(), &videorpc.CreateVideoRequest{
		UserId:   userId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    formParamVar.Title,
	})

	if err != nil {
		SendStatusResponse(c, errno.ConvertErr(err).WithMessage("创建视频失败"))
		return
	}
	SendStatusResponse(c, errno.Success)

}

func PublishList(c *gin.Context) {
	type param struct {
		Token  string `form:"token"`
		UserId int64  `form:"user_id"`
	}
	type videoListResponse struct {
		StatusCode int32             `json:"status_code"` // 状态码，0-成功，其他值-失败
		StatusMsg  string            `json:"status_msg"`  // 返回状态描述
		VideoList  []*videorpc.Video `json:"video_list"`  // 用户信息列表
	}

	var paramVar param
	if err := c.ShouldBind(&paramVar); err != nil {
		SendStatusResponse(c, errno.ConvertErr(err))
		return
	}

	videos, err := rpc.QueryVideoByUserId(context.Background(), &videorpc.QueryVideoByUserIdRequest{UserId: paramVar.UserId})

	if err != nil {
		SendStatusResponse(c, errno.ConvertErr(err))
		return
	}

	c.JSON(http.StatusOK, videoListResponse{
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  errno.Success.ErrMsg,
		VideoList:  videos,
	})

}
