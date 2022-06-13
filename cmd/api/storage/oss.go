package storage

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/hh02/minimal-douyin/pkg/constants"
)

var ossClient *oss.Client
var ossBucket *oss.Bucket

func initOss() {
	ossClient, err := oss.New(constants.OssEndpoint, constants.OssAccessKeyId, constants.OssAccessKeySecret)
	if err != nil {
		panic(err)
	}

	ossBucket, err = ossClient.Bucket(constants.OssBucketName)
	if err != nil {
		panic(err)
	}
}

func PutObject(filename string, file io.Reader) error {
	return ossBucket.PutObject(filename, file)
}
