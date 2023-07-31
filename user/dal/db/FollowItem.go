package db

import (
	"context"
	_const "douyinv2/user/const"
	"douyinv2/user/dal/model"
	"douyinv2/user/errons"
	"douyinv2/user/kitex_gen/user"
)

func CreateFollowItem(ctx context.Context, fromid int64, toid int64) errons.Errorn {
	followiten := &model.UserFollow{
		FromID: fromid,
		ToID:   toid,
	}
	err := Q.WithContext(ctx).UserFollow.Create(followiten)
	if err != nil {
		return errons.ConverterrtoErr(err)
	}
	return errons.Successcode
}
func FollowList(ctx context.Context, fromid int64) ([]*user.User, errons.Errorn) {
	var err2 errons.Errorn
	var users []*user.User
	err2 = errons.Successcode
	temps, err := Q.WithContext(ctx).UserFollow.Where(Q.UserFollow.FromID.Eq(fromid)).Find()
	if err != nil {
		err2 = errons.ConverterrtoErr(err)
		return users, err2
	}
	for i := 0; i < len(temps); i++ {
		followcount, _ := Q.WithContext(ctx).UserFollow.Where(Q.UserFollow.FromID.Eq(temps[i].ToID)).Count()
		followercount, _ := Q.WithContext(ctx).UserFollow.Where(Q.UserFollow.ToID.Eq(temps[i].ToID)).Count()
		theuser, err := UserInfo(ctx, temps[i].ID)
		if err != nil {
			err2 = errons.ConverterrtoErr(err)
			return users, err2
		}
		tempuser := &user.User{
			Id:              theuser.ID,
			Name:            theuser.Username,
			Avatar:          theuser.Avatar,
			IsFollow:        true,
			FollowCount:     followcount,
			FollowerCount:   followercount,
			BackgroundImage: theuser.BackgroundImage,
			Signature:       theuser.Signature,
			TotalFavorited:  theuser.FavoriteCount,
			WorkCount:       theuser.WorkCount,
			FavoriteCount:   theuser.FavoriteCount,
		}
		users = append(users, tempuser)
	}
	return users, err2
}
func DeleteFollowItem(ctx context.Context, fromid int64, toid int64) errons.Errorn {
	_, err := Q.WithContext(ctx).UserFollow.Where(Q.UserFollow.FromID.Eq(fromid)).Where(Q.UserFollow.ToID.Eq(toid)).Delete()
	if err != nil {
		return errons.ConverterrtoErr(err)
	}
	return errons.Successcode
}
func FollowerList(ctx context.Context, toid int64) ([]*user.User, errons.Errorn) {
	var err2 errons.Errorn
	var users []*user.User
	err2 = errons.Successcode
	temps, err := Q.WithContext(ctx).UserFollow.Where(Q.UserFollow.ToID.Eq(toid)).Find()
	if err != nil {
		err2 = errons.ConverterrtoErr(err)
		return users, err2
	}
	for i := 0; i < len(temps); i++ {
		followcount, _ := Q.WithContext(ctx).UserFollow.Where(Q.UserFollow.FromID.Eq(temps[i].FromID)).Count()
		followercount, _ := Q.WithContext(ctx).UserFollow.Where(Q.UserFollow.ToID.Eq(temps[i].FromID)).Count()
		test, _ := Q.WithContext(ctx).UserFollow.Where(Q.UserFollow.FromID.Eq(temps[i].FromID)).Where(Q.UserFollow.ToID.Eq(toid)).Find()
		iffollow := false
		if test != nil {
			iffollow = true
		}
		theuser, err := UserInfo(ctx, temps[i].ID)
		if err != nil {
			err2 = errons.ConverterrtoErr(err)
			return users, err2
		}
		tempuser := &user.User{
			Id:              theuser.ID,
			Name:            theuser.Username,
			Avatar:          theuser.Avatar,
			IsFollow:        iffollow,
			FollowCount:     followcount,
			FollowerCount:   followercount,
			BackgroundImage: theuser.BackgroundImage,
			Signature:       theuser.Signature,
			TotalFavorited:  theuser.FavoriteCount,
			WorkCount:       theuser.WorkCount,
			FavoriteCount:   theuser.FavoriteCount,
		}
		users = append(users, tempuser)
	}
	return users, err2
}

func FriendList(ctx context.Context, userid int64) ([]*user.FriendUser, errons.Errorn) {
	var friends []*user.FriendUser
	p := Q.UserFollow
	q := Q.UserFollow
	//关注该用户的人
	followers, _ := q.WithContext(ctx).Where(q.ToID.Eq(userid)).Select(q.FromID).Find()
	var followersid []int64
	for i := 0; i < len(followers); i++ {
		followersid = append(followersid, followers[i].FromID)
	}
	//用户关注的人中关注该用户的人的id
	res, err2 := p.Where(p.FromID.Eq(userid)).Where(p.ToID.In(followersid...)).Find()
	if err2 != nil {
		return friends, errons.ConverterrtoErr(err2)
	}
	for i := 0; i < len(res); i++ {
		theuser, err := UserInfo(ctx, res[i].ToID)
		if err != nil {
			return friends, errons.ConverterrtoErr(err)
		}
		followcount, _ := Q.WithContext(ctx).UserFollow.Where(Q.UserFollow.FromID.Eq(res[i].ToID)).Count()
		followercount, _ := Q.WithContext(ctx).UserFollow.Where(Q.UserFollow.ToID.Eq(res[i].ToID)).Count()
		tempuser := &user.User{
			Id:              theuser.ID,
			Name:            theuser.Username,
			Avatar:          theuser.Avatar,
			IsFollow:        true,
			FollowCount:     followcount,
			FollowerCount:   followercount,
			BackgroundImage: theuser.BackgroundImage,
			Signature:       theuser.Signature,
			TotalFavorited:  theuser.FavoriteCount,
			WorkCount:       theuser.WorkCount,
			FavoriteCount:   theuser.FavoriteCount,
		}
		friends = append(friends, _const.ConvertUserToFriend(tempuser, "", 1))
	}
	return friends, errons.Successcode
}
func IFFollow(ctx context.Context, fromid int64, toid int64) (res bool) {
	user, err := Q.WithContext(ctx).UserFollow.Where(Q.UserFollow.FromID.Eq(fromid)).Where(Q.UserFollow.ToID.Eq(toid)).First()
	if err != nil || user == nil {
		return false
	}
	return true
}
