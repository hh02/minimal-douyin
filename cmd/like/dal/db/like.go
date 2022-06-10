package db

import (
	"context"
	"errors"

	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/gorm"
)

type Like struct {
	// 用户ID
	UserId int64 `gorm:"primaryKey;autoIncrement:false"`
	// 视频ID(联合主键的第二个键需要添加索引)
	VideoId    int64 `gorm:"primaryKey;index;autoIncrement:false"`
	CreateTime int64 `gorm:"autoCreateTime:nano"`
}

func (f *Like) TableName() string {
	return constants.LikeTableName
}

// 点赞
func CreateLike(ctx context.Context, like *Like) error {
	return DB.WithContext(ctx).Create(like).Error
}

// 取消点赞
func DeleteLike(ctx context.Context, like *Like) (int64, error) {
	result := DB.WithContext(ctx).Delete(like)
	return result.RowsAffected, result.Error
}

// QueryLikeList 返回点赞视频 id 列表
func QueryLikeList(ctx context.Context, userId int64) ([]int64, error) {
	var res []int64
	if err := DB.WithContext(ctx).Table(constants.LikeTableName).Order("create_time desc").Select("video_id").Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetLike 根据复合主键查询记录
func GetLike(ctx context.Context, like *Like) (*Like, error) {
	var res Like
	if err := DB.WithContext(ctx).Where(map[string]interface{}{
		"user_id":   like.UserId,
		"follow_id": like.VideoId,
	}).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}
