package db

import (
	"context"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/gorm"
)

// 用户服务
type User struct {
	gorm.Model
	// username's maxlen is 64
	Username string `gorm:"type:varchar(64);not null"`
	// md5's length is 128
	Password      string `gorm:"type:char(128);not null"`
	FollowCount   int64  `gorm:"default:0"`
	FollowerCount int64  `gorm:"default:0"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// 用户注册，向数据库中添加用户。返回错误，以及当前用户ID。
func CreateUser(ctx context.Context, user *User) (error, int64) {
	err := DB.WithContext(ctx).Create(user).Error
	return err, int64(user.ID)
}

// 通过名字，查询用户
func QueryUserByName(ctx context.Context, UserName string) (*User, error) {
	res := &User{}
	err := DB.WithContext(ctx).Where("name = ?", UserName).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 通过 id 查询用户
func QueryUserById(ctx context.Context, Id int64) (*User, error) {
	res := &User{}
	err := DB.WithContext(ctx).Where("id = ?", uint(Id)).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

//通过多个 id 查询多个用户
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
