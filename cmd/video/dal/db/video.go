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

func 