package main

//type Userlogin struct {
//	gorm.Model
//	Username         string `json:"username"`
//	Password         string `json:"password"`
//	Avatar           string `gorm:"column:avatar;" json:"avatar"`
//	Background_image string `gorm:"column.background_image" json:"background_image"`
//	Signature        string `gorm:"column.signature; default:hello world" json:"signature"`
//	Total_favorited  string `gorm:"column.total_favorited; default:0" json:"total_favorited"`
//	Work_count       int    `gorm:"column.work_count;default:0" json:"Work_count"`
//	Favorite_count   int    `gorm:"column.favorite_count;default:0" json:"favorite_count"`
//}
//
//type UserFollow struct {
//	gorm.Model
//	From_id uint `json:"from_Id"`
//	To_id   uint `json:"to_Id"`
//}

//建表

//func main() {
//	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", _const.DBUser, _const.DBpass, _const.DBCONNECT, _const.DBIP, _const.DBPORT, _const.DBNAME, _const.DBCHAR)
//	fmt.Println(dsn)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("can't connect to the dbs")
//	}
//	err = db.AutoMigrate(&UserFollow{})
//	if err != nil {
//		panic("creaete_talbe_failed")
//	}
//}

// gen生成CRUD（增删改查）
//func main() {
//	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", _const.DBUser, _const.DBpass, _const.DBCONNECT, _const.DBIP, _const.DBPORT, _const.DBNAME, _const.DBCHAR)
//	fmt.Println(dsn)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("connect to db failed")
//	}
//	g := gen.NewGenerator(gen.Config{
//		OutPath: "/home/unbuntu/GolandProjects/douyinv2/user/dal",
//		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
//	})
//
//	g.UseDB(db)
//	g.ApplyBasic(g.GenerateAllTable()...)
//	g.Execute()
//}

// 测试CRUD
//
//	func main() {
//		dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", _const.DBUser, _const.DBpass, _const.DBCONNECT, _const.DBIP, _const.DBPORT, _const.DBNAME, _const.DBCHAR)
//		fmt.Println(dsn)
//		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//		if err != nil {
//			panic("connect to db failed")
//		}
//		query.SetDefault(db)
//		u := query.Userlogin
//		ctx := context.Background()
//		users := []*model.Userlogin{{Username: "ieee", Password: "123"}, {Username: "zz", Password: "321"}}
//		//var user *model.Userlogin
//		//err = u.WithContext(ctx).Create(users...)
//		users, err = u.WithContext(ctx).Where(u.Username.Eq("iee")).Find()
//		if err != nil {
//			fmt.Print("写入失败")
//		}
//		fmt.Println(len(users))
//	}
//func main() {
//	db.Init()
//	db.FriendList(context.Background(), 1)
//}
