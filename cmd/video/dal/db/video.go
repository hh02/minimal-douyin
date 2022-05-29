package db

import (
	"context"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Id int64
	AuthorId int64
	PlayUrl string
	CoverUrl string
	Title string
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func CreateVideo(ctx context.Context, videos []*Video) error {
	return DB.WithContext(ctx).Create(videos).Error
}

func QueryVideoByUserId(ctx context.Context, userId int64) ([]*Video, error){
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("author_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}