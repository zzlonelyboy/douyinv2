package _const

import "douyinv2/user/kitex_gen/user"

func ConvertUserToFriend(user1 *user.User, msg string, msgtype int64) *user.FriendUser {

	return &user.FriendUser{
		Id:              user1.Id,
		Name:            user1.Name,
		FollowCount:     user1.FollowCount,
		FollowerCount:   user1.FollowerCount,
		IsFollow:        user1.IsFollow,
		Avatar:          user1.Avatar,
		BackgroundImage: user1.BackgroundImage,
		Signature:       user1.Signature,
		TotalFavorited:  user1.TotalFavorited,
		FavoriteCount:   user1.FavoriteCount,
		WorkCount:       user1.WorkCount,
		Message:         msg,
		MsgType:         msgtype,
	}
}
