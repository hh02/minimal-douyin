package dal

import "github.com/hh02/minimal-douyin/cmd/user/dal/db"

func Init() {
	//初始化 mysql 数据库
	db.Init()
}
