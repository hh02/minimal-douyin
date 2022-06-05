package db

import (
	"context"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/gorm"
)

// Comment 用户服务
type Comment struct {
	gorm.Model

	UserId  int64
	VideoId int64
	Content string `gorm:"type:char(100);not null"`
}

func (u *Comment) TableName() string {
	return constants.CommentTableName
}

// CreateComment  创建评论
func CreateComment(ctx context.Context, comment *Comment) error {
	return DB.WithContext(ctx).Create(comment).Error
}

// QueryCommentByCommentId 根据 CommentId 查询评论
func QueryCommentByCommentId(ctx context.Context, commentId int64) (*Comment, error) {
	res := &Comment{}
	err := DB.WithContext(ctx).Where("id = ?", commentId).First(&res).Error
	return res, err
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, commentId int64) error {
	return DB.WithContext(ctx).Delete(&Comment{}, commentId).Error
}

// QueryCommentByVideoId 根据 VideoId 查询评论
func QueryCommentByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	err := DB.WithContext(ctx).Where("video_id = ?", videoId).Find(&res).Error
	return res, err
}
