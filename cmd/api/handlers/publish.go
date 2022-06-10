package handlers

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
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

	err = imaging.Save(img, snapshotPath)
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

	c.AddParam("token", formParamVar.Token)
	// 调用鉴权
	c.Next()

	// 鉴权失败则返回
	if c.IsAborted() {
		return
	}

	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))

	videoName := uuid.NewV4().String()

	videoLocalPath := constants.VideoLocalPath + "/" + videoName + ".mp4"
	coverLocalPath := constants.CoverLocalPath + "/" + videoName + ".png"

	c.SaveUploadedFile(formParamVar.Data, videoLocalPath)
	err := getSnapshot(videoLocalPath, coverLocalPath, 1)
	if err != nil {
		SendStatusResponse(c, errno.ConvertErr(err).WithMessage("获取封面失败"))
		return
	}

	playUrl := constants.VideoServerPath + "/" + videoName + ".mp4"
	coverUrl := constants.CoverServerPath + "/" + videoName + ".png"

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

	var paramVar param
	if err := c.ShouldBind(&paramVar); err != nil {
		SendVideoListResponse(c, err, nil)
		return
	}

	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64))

	videos, err := rpc.QueryVideoByUserId(context.Background(), &videorpc.QueryVideoByUserIdRequest{UserId: userId})

	if err != nil {
		SendVideoListResponse(c, err, nil)
		return
	}
	SendVideoListResponse(c, errno.Success, videos)
}
