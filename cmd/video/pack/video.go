package pack

import (
	"github.com/hh02/minimal-douyin/cmd/video/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
)

// model to rpc message
func VideoModel2RpcMessage(m *db.Video) *videorpc.Video {
	if m == nil {
		return nil
	}

	return &videorpc.Video{
		VideoId: m.VideoId,
		PlayUrl: m.PlayUrl,
		CoverUrl: m.CoverUrl,
		Title: m.Title,
		FavoriteCount: m.FavoriteCount,
		CommentCount: m.CommentCount,
	}
}

func VideoModels2RpcMessages (ms []*db.Video) []*videorpc.Video {
	videos := make([]*videorpc.Video, len(ms))
	for i, m := range ms {
		if res := VideoModel2RpcMessage(m); res != nil {
			videos[i] = res
		}
	}
	return videos
}