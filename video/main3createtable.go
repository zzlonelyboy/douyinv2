package main

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	PostUserId int    `json:"postUserId"`
	Playurl    string `gorm:"column:playurl" json:"playurl"`
	Coveurl    string `gorm:"column:coverurl" json:"coveurl"`
	Title      string `gorm:"column:title" json:"title"`
}

type VideoLike struct {
	gorm.Model
	UserID  int `json:"userID"`
	VideoID int `json:"videoID"`
}
type VideoComment struct {
	gorm.Model
	UserID     int    `json:"userID"`
	VideoID    int    `json:"videoID"`
	CommentCon string `json:"commentCon"`
}

//func main() {
//	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", _const.DBUser, _const.DBpass, _const.DBCONNECT, _const.DBIP, _const.DBPORT, _const.DBNAME, _const.DBCHAR)
//	fmt.Println(dsn)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("can't connect to the dbs")
//	}
//	err = db.AutoMigrate(&Video{})
//	err = db.AutoMigrate(VideoComment{})
//	err = db.AutoMigrate(VideoLike{})
//	if err != nil {
//		panic("creaete_talbe_failed")
//	}
//}

//func main() {
//	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", _const.DBUser, _const.DBpass, _const.DBCONNECT, _const.DBIP, _const.DBPORT, _const.DBNAME, _const.DBCHAR)
//	fmt.Println(dsn)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("connect to db failed")
//	}
//	g := gen.NewGenerator(gen.Config{
//		OutPath: "/home/unbuntu/GolandProjects/douyinv2/video/dal",
//		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
//	})
//
//	g.UseDB(db)
//	g.ApplyBasic(g.GenerateModelAs("videos", "Videos"))
//	g.ApplyBasic(g.GenerateModelAs("video_likes", "VideoLikes"))
//	g.ApplyBasic(g.GenerateModelAs("video_comments", "VideoComments"))
//	g.Execute()
//}

func maimn() {
	p := time.Now().Unix()
	timelayout := "2019-01-01 07:08:09.12"
	timestr := time.Unix(p, 0).Format(timelayout)
	fmt.Println(timestr)
}
