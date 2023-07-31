package db

import (
	_const "douyinv2/user/const"
	"douyinv2/user/dal/query"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Q *query.Query

func Init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", _const.DBUser, _const.DBpass, _const.DBCONNECT, _const.DBIP, _const.DBPORT, _const.DBNAME, _const.DBCHAR)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect to dbs failed")
	}
	query.SetDefault(db)
	Q = query.Q
	if err != nil {
		panic(err)
	}
}
