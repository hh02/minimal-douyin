package db

import (
	"context"

	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/gorm"
)

type Video struct {
	VideoId int64 `gorm:"primarykey"`
	// 2083 reference to "https://www.racecoder.com/archives/508/"
	PlayUrl  string `gorm:"type:varchar(2083);not null"`
	CoverUrl string `gorm:"type:varchar(2083);not null"`
	Title    string `gorm:"type:varchar(100);not null"`
	CreateTime int64 `gorm:"autoUpdateTime:nano"`
	FavoriteCount int64
	CommentCount int64
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func CreateVideo(ctx context.Context, videos []*Video) error {
	return DB.WithContext(ctx).Create(videos).Error
}

func QueryVideoByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
