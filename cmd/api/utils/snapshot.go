package utils

import (
	"bytes"
	"fmt"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// 获取视频的某一帧，一般用第一帧作为视频封面
func GetSnapshot(videoPath string, frameNum int) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	return buf, nil
}
