package db

import (
	"context"
	"errors"

	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/gorm"
)

type Favorite struct {
	// 用户ID
	UserId int64 `gorm:"primaryKey;autoIncrement:false"`
	// 视频ID(联合主键的第二个键需要添加索引)
	VideoId    int64 `gorm:"primaryKey;index;autoIncrement:false"`
	CreateTime int64 `gorm:"autoCreateTime:nano"`
}

func (f *Favorite) TableName() string {
	return constants.FavoriteTableName
}

// 点赞
func CreateFavorite(ctx context.Context, favorite *Favorite) error {
	return DB.WithContext(ctx).Create(favorite).Error
}

// 取消点赞
func DeleteFavorite(ctx context.Context, favorite *Favorite) (int64, error) {
	result := DB.WithContext(ctx).Delete(favorite)
	return result.RowsAffected, result.Error
}

// QueryFavoriteList 返回点赞视频 id 列表
func QueryFavoriteList(ctx context.Context, userId int64) ([]int64, error) {
	var res []int64
	if err := DB.WithContext(ctx).Table(constants.FavoriteTableName).Order("create_time desc").Select("video_id").Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetFavorite 根据复合主键查询记录
func GetFavorite(ctx context.Context, favorite *Favorite) (*Favorite, error) {
	var res Favorite
	if err := DB.WithContext(ctx).Where(map[string]interface{}{
		"user_id":  favorite.UserId,
		"video_id": favorite.VideoId,
	}).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}
