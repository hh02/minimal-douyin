package db

import (
	"context"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/gorm"
)

// User 用户服务
type User struct {
	gorm.Model
	// username's maxlen is 32
	Username string `gorm:"type:varchar(32);not null"`
	// md5's length is 32
	Password      string `gorm:"type:char(32);not null"`
	FollowCount   int64  `gorm:"default:0"`
	FollowerCount int64  `gorm:"default:0"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser 用户注册，向数据库中添加用户。返回错误，以及当前用户ID。
func CreateUser(ctx context.Context, user *User) (error, int64) {
	err := DB.WithContext(ctx).Create(&user).Error
	return err, int64(user.ID)
}

// QueryUserByName 通过名字，查询用户
func QueryUserByName(ctx context.Context, UserName string) (*User, error) {
	res := &User{}
	err := DB.WithContext(ctx).Where("username = ?", UserName).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// QueryUserById 通过 id 查询用户
func QueryUserById(ctx context.Context, Id int64) (*User, error) {
	res := &User{}
	err := DB.WithContext(ctx).Where("id = ?", uint(Id)).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MQueryUserById 通过多个 id 查询多个用户
func MQueryUserById(ctx context.Context, IDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(IDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", IDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// AddFollowCount 对指定用户增加 count 关注数
func AddFollowCount(ctx context.Context, Id int64, count int32) error {
	return DB.WithContext(ctx).Where("id = ?", Id).Update("follow_count", gorm.Expr("follow_count + ?", count)).Error
}

// AddFollowerCount 对指定用户增加 count 粉丝数
func AddFollowerCount(ctx context.Context, Id int64, count int32) error {
	return DB.WithContext(ctx).Where("id = ?", Id).Update("follower_count", gorm.Expr("follow_counter + ?", count)).Error
}
