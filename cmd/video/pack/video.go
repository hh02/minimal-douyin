package pack

import (
	"github.com/hh02/minimal-douyin/cmd/video/dal/db"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"github.com/hh02/minimal-douyin/kitex_gen/videorpc"
)

// convert db.Video to videorpc.Video
func Video(video *db.Video) *videorpc.Video {
	if video == nil {
		return nil
	}
	return &videorpc.Video{
		VideoId:       int64(video.ID),
		PlayUrl:       video.PlayUrl,
		Author:        &userrpc.User{UserId: video.UserId},
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		Title:         video.Title,
	}
}

// convert multiple db.Video to multiple videorpc.Video
func Videos(videos []*db.Video) []*videorpc.Video {
	res := make([]*videorpc.Video, 0)
	for _, video := range videos {
		if convreted := Video(video); convreted != nil {
			res = append(res, convreted)
		}
	}
	return res
}

func UserIds(videos []*db.Video) []int64 {
	userIds := make([]int64, 0)
	if len(videos) == 0 {
		return userIds
	}

	userIdMap := make(map[int64]struct{})
	for _, video := range videos {
		if video != nil {
			userIdMap[video.UserId] = struct{}{}
		}
	}
	for userId := range userIdMap {
		userIds = append(userIds, userId)
	}
	return userIds
}
