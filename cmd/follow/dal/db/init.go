package db

import (
	"github.com/hh02/minimal-douyin/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init() {
	var err error

	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	m := DB.Migrator()

	if m.HasTable(&Follow{}) {
		return
	}
	if err = m.CreateTable(&Follow{}); err != nil {
		panic(err)
	}

}
