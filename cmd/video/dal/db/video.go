package db

import (
	"context"

	"github.com/hh02/minimal-douyin/pkg/constants"
)

type Video struct {
	Id     int64 `gorm:"primarykey"`
	UserId int64
	// 2083 reference to "https://www.racecoder.com/archives/508/"
	PlayUrl       string `gorm:"type:varchar(2083);not null"`
	CoverUrl      string `gorm:"type:varchar(2083);not null"`
	Title         string `gorm:"type:varchar(100);not null"`
	FavoriteCount int64  `gorm:"default:0"`
	CommentCount  int64  `gorm:"default:0"`
	CreateTime    int64  `gorm:"autoCreateTime"`
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func CreateVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

func QueryVideoByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	var res []*Video
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryVideoByTime(ctx context.Context, latestTime int64) ([]*Video, error) {
	var res []*Video
	if err := DB.WithContext(ctx).Where("create_time <= ?", latestTime).Order("create_time desc").Limit(constants.TotalFeedNumber).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func MgetVideo(ctx context.Context, videoIds []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if len(videoIds) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", videoIds).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
