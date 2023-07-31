package service

import (
	"context"
	"douyinv2/biz/mw"
	"douyinv2/user/dal/db"
	"douyinv2/user/errons"
	"douyinv2/user/kitex_gen/user"
	"github.com/form3tech-oss/jwt-go"
	"strconv"
)

type UserInfoService struct {
	ctx context.Context
}

func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{
		ctx: ctx,
	}
}

func (s *UserInfoService) UserInfo(req *user.UserInfoRequest) (user.User, errons.Errorn) {
	user1, err := db.UserInfo(s.ctx, req.UserId)
	var ID int64
	ID = -1
	isfollow := false
	if err != nil {
		return *user.NewUser(), errons.ConverterrtoErr(err)
	}
	temp, _ := strconv.Atoi(user1.TotalFavorited)
	resolvetoken, err := jwt.ParseWithClaims(req.Token, &mw.Myclaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("douyintest"), nil
	})
	if err == nil {
		myclaims := resolvetoken.Claims.(*mw.Myclaims)
		ID = myclaims.Id
	}
	if ID != -1 {
		isfollow = db.IFFollow(s.ctx, ID, req.UserId)
	}
	followcount, _ := db.Q.WithContext(s.ctx).UserFollow.Where(db.Q.UserFollow.FromID.Eq(req.UserId)).Count()
	followercount, _ := db.Q.WithContext(s.ctx).UserFollow.Where(db.Q.UserFollow.ToID.Eq(req.UserId)).Count()
	return user.User{
		Id:              user1.ID,
		Name:            user1.Username,
		FollowCount:     followcount,
		FollowerCount:   followercount,
		IsFollow:        isfollow,
		Avatar:          user1.Avatar,
		BackgroundImage: user1.BackgroundImage,
		Signature:       user1.Signature,
		TotalFavorited:  int64(temp),
		FavoriteCount:   user1.FavoriteCount,
		WorkCount:       user1.WorkCount,
	}, errons.Successcode
}
