package db

import (
	"context"
	"douyinv2/user/dal/model"
	"strconv"
)

// 增加用户
func CreateUser(ctx context.Context, users []*model.Userlogin) error {
	err := Q.WithContext(ctx).Userlogin.Create(users...)
	return err
}

// 查询用户
func QueryUser(ctx context.Context, username string) (user []*model.Userlogin, err error) {
	users, err := Q.WithContext(ctx).Userlogin.Where(Q.Userlogin.Username.Eq(username)).Find()
	return users, err
}
func UserLogin(ctx context.Context, username string, password string) (user *model.Userlogin, err error) {
	users, err := Q.WithContext(ctx).Userlogin.Where(Q.Userlogin.Username.Eq(username)).Where(Q.Userlogin.Password.Eq(password)).First()
	return users, err
}
func UserInfo(ctx context.Context, userID int64) (user *model.Userlogin, err error) {
	users, err := Q.WithContext(ctx).Userlogin.Where(Q.Userlogin.ID.Eq(userID)).First()
	return users, err
}
func CountChange(ctx context.Context, userID int64, filedname string, ifadd bool) error {
	var err error
	var value int64
	value = 1
	if !ifadd {
		value = -1
	}
	if filedname == "Like" {
		_, err = Q.WithContext(ctx).Userlogin.Where(Q.Userlogin.ID.Eq(userID)).UpdateSimple(Q.Userlogin.FavoriteCount.Add(value))
	}
	if filedname == "Work" {
		_, err = Q.WithContext(ctx).Userlogin.Where(Q.Userlogin.ID.Eq(userID)).UpdateSimple(Q.Userlogin.WorkCount.Add(value))
	}
	if filedname == "Liked" {
		user, _ := Q.WithContext(ctx).Userlogin.Where(Q.Userlogin.ID.Eq(userID)).First()
		olddata, _ := strconv.Atoi(user.TotalFavorited)
		olddata = int(int64(olddata) + value)
		newdata := strconv.Itoa(olddata)
		_, err = Q.WithContext(ctx).Userlogin.Where(Q.Userlogin.ID.Eq(userID)).UpdateSimple(Q.Userlogin.TotalFavorited.Value(newdata))
	}
	return err
}
