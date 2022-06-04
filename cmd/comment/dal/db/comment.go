package db

import (
	"context"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/gorm"
)

// User 用户服务
type Comment struct {
	gorm.Model

	UserId  int64
	VideoId int64
	content string `gorm:"type:char(100);not null"`
}

func (u *Comment) TableName() string {
	return constants.CommentTableName
}

// CreateCommwnt  创建评论
func CreateComment(ctx context.Context, comment *Comment) error {
	return DB.WithContext(ctx).Create(comment).Error
}
