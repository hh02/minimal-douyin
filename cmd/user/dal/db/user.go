package db

import (
	"context"
	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `json:"name,omitempty"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// 用户注册，向数据库中添加用户。返回错误，以及当前用户ID。
func UserRegister(ctx context.Context, user *User) (error, uint) {
	err := DB.WithContext(ctx).Create(user).Error
	return err, user.ID
}

// 通过名字，查询用户是否存在
func QueryUserByName(ctx context.Context, UserName string) (*User, error) {
	res := &User{}
	err := DB.WithContext(ctx).Where("name = ?", UserName).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 通过 id 查询用户是否存在
func QueryUserById(ctx context.Context, Id int64) (*User, error) {
	res := &User{}
	err := DB.WithContext(ctx).Where("id = ?", uint(Id)).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
