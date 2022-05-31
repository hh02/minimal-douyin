package db

import (
	"context"

	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/gorm"
)

type Follow struct {
	// 粉丝ID
	UserId int64 `gorm:"primaryKey;autoIncrement:false"`
	// 关注ID(联合主键的第二个键需要添加索引)
	FollowId   int64          `gorm:"primaryKey;index;autoIncrement:false"`
	CreateTime int64          `gorm:"autoCreateTime:nano"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (f *Follow) TableName() string {
	return constants.FollowTableName
}

// 关注
func CreateFollow(ctx context.Context, follow *Follow) error {
	return DB.WithContext(ctx).Create(follow).Error
}

// 取消关注
func DeleteFollow(ctx context.Context, follow *Follow) error {
	return DB.WithContext(ctx).Delete(follow).Error
}

// 返回关注列表
func QueryFollow(ctx context.Context, userId int64) ([]int64, error) {
	var res []int64
	if err := DB.WithContext(ctx).Order("create_time desc").Select("follow_id").Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// 返回粉丝列表
func QueryFollower(ctx context.Context, userId int64) ([]int64, error) {
	var res []int64
	if err := DB.WithContext(ctx).Select("user_id").Where("follow_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// 根据复合主键查询记录
func GetFollow(ctx context.Context, follow *Follow) (*Follow, error) {
	var res Follow
	if err := DB.WithContext(ctx).Where(map[string]interface{}{
		"user_id":   follow.UserId,
		"follow_id": follow.FollowId,
	}).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
