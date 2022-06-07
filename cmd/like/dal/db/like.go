package db

import (
	"context"
	"errors"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/gorm"
)

type Like struct {
	//点赞id
	UserId int64 `gorm:"primaryKey;autoIncrement:false"`

	VideoId    int64          `gorm:"primaryKey;autoIncrement:false"`
	CreateTime int64          `gorm:"autoCreateTime:nano"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (l *Like) TableName() string {
	return constants.LikeTableName
}
func CreateLike(ctx context.Context, l *Like) (int64, error) {
	create := DB.WithContext(ctx).Create(l)
	return create.RowsAffected, create.Error
}
func DeleteLike(ctx context.Context, l *Like) (int64, error) {
	tx := DB.WithContext(ctx).Delete(l)
	return tx.RowsAffected, tx.Error
}

func QueryLike(ctx context.Context, videoID int64) ([]int64, error) {
	var res []int64
	if err := DB.WithContext(ctx).Order("create_time desc").Select("user_id").Where("video_id=?", videoID).Find(&res).Error; err != nil {
		return res, nil
	}
	return res, nil
}

func Getlike(ctx context.Context, like *Like) (*Like, error) {
	var res Like
	if err := DB.WithContext(ctx).Where(map[string]interface{}{
		"userID":  like.UserId,
		"videoId": like.VideoId,
	}).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil

}
