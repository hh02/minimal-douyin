package handlers

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hh02/minimal-douyin/cmd/api/rpc"
	"github.com/hh02/minimal-douyin/cmd/api/storage"
	"github.com/hh02/minimal-douyin/cmd/api/utils"
	"github.com/hh02/minimal-douyin/kitex_gen/response"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"github.com/hh02/minimal-douyin/pkg/errno"
	uuid "github.com/satori/go.uuid"
)

func SendPublishActionResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.PublishActionResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}
func SendPublishListResponse(c *gin.Context, err error, videos []*videorpc.Video) {
	Err := errno.ConvertErr(err)
	utils.PbJSONResponse(c, http.StatusOK, &response.PublishListResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		VideoList:  videos,
	})
}

func PublishAction(c *gin.Context) {
	type formParam struct {
		Data  *multipart.FileHeader `form:"data" binding:"required"`
		Token string                `form:"token" binding:"required"`
		Title string                `form:"title" binding:"required"`
	}

	var formParamVar formParam
	if err := c.ShouldBind(&formParamVar); err != nil {
		SendPublishActionResponse(c, err)
		return
	}

	if len(formParamVar.Title) > 100 {
		SendPublishActionResponse(c, errno.ParamErr.WithMessage("标题长度超过100"))
		return
	}

	c.AddParam("token", formParamVar.Token)
	// 调用鉴权
	c.Next()

	// 鉴权失败则返回
	if err := utils.CheckAuth(c); err != nil {
		SendPublishActionResponse(c, err)
		return
	}

	userId := utils.GetIdFromClaims(c)
	if userId == 0 {
		SendPublishActionResponse(c, errno.GetIdFromClaimsErr)
		return
	}
	videoFile, err := formParamVar.Data.Open()
	if err != nil {
		SendPublishActionResponse(c, errno.OpenVideoFileErr.WithMessage(err.Error()))
		return
	}
	// 使用 uuid 作为上传文件的名称，防止冲突
	videoName := uuid.NewV4().String()
	videoFilename := videoName + ".mp4"
	coverFilename := videoName + ".png"
	videoLocalPath := constants.TempFoler + videoFilename
	if err := c.SaveUploadedFile(formParamVar.Data, videoLocalPath); err != nil {

	}

	cover, err := utils.GetSnapshot(videoLocalPath, 1)
	if err != nil {
		SendPublishActionResponse(c, errno.GetCoverErr)
		return
	}

	videoUrl := constants.OssUrlPrefix + videoFilename
	coverUrl := constants.OssUrlPrefix + coverFilename
	if err := storage.PutObject(videoFilename, videoFile); err != nil {
		SendPublishActionResponse(c, errno.OssPutObjectErr.WithMessage(err.Error()))
		return
	}
	if err := storage.PutObject(coverFilename, cover); err != nil {
		SendPublishActionResponse(c, errno.OssPutObjectErr.WithMessage(err.Error()))
		return
	}

	err = rpc.CreateVideo(context.Background(), &videorpc.CreateVideoRequest{
		UserId:   userId,
		PlayUrl:  videoUrl,
		CoverUrl: coverUrl,
		Title:    formParamVar.Title,
	})

	if err != nil {
		SendPublishActionResponse(c, err)
		return
	}
	SendPublishActionResponse(c, errno.Success)

}

func PublishList(c *gin.Context) {
	if err := utils.CheckAuth(c); err != nil {
		SendPublishListResponse(c, err, nil)
		return
	}
	type param struct {
		Token  string `form:"token"`
		UserId int64  `form:"user_id"`
	}

	var paramVar param
	if err := c.ShouldBind(&paramVar); err != nil {
		SendPublishListResponse(c, err, nil)
		return
	}

	userId := utils.GetIdFromClaims(c)
	if userId == 0 {
		SendPublishListResponse(c, errno.GetIdFromClaimsErr, nil)
		return
	}

	videos, err := rpc.QueryVideoByUserId(context.Background(), &videorpc.QueryVideoByUserIdRequest{UserId: userId})

	if err != nil {
		SendPublishListResponse(c, err, nil)
		return
	}
	SendPublishListResponse(c, errno.Success, videos)
}
