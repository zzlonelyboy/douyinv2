package rpc

import (
	"context"
	douyin "douyinv2/biz/model/douyinv2"
	"douyinv2/user/kitex_gen/user"
	"douyinv2/user/kitex_gen/user/apiservice"
	"douyinv2/video/errons"
	video "douyinv2/video/kitex_gen/video"
	"douyinv2/video/kitex_gen/video/videoservice"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var userclient apiservice.Client
var videoclient videoservice.Client

func Init() {
	r, err := etcd.NewEtcdResolver([]string{"192.168.217.1:2379"})
	if err != nil {
		panic(err)
	}
	newClient, err := apiservice.NewClient("Userservice", client.WithResolver(r))
	newvideoclient := videoservice.MustNewClient("VideoService", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
	userclient = newClient
	videoclient = newvideoclient
}

func UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp, err = userclient.RegisterUser(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp, err = userclient.LoginUser(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	resp, err = userclient.UserInfo(ctx, req)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// VIDEO
func VideoPublish(ctx context.Context, req *douyin.VideoPublishRequestv2) (resp *douyin.VideoPublishResponse, err error) {
	resp2, err := videoclient.VideoPublish(ctx, &video.VideoPublishRequest{ID: req.ID, Filepath: req.Filepath, Cover: req.Coverurl, Title: req.Title})
	if err != nil {
		panic(err)
	}
	resp = &douyin.VideoPublishResponse{}
	resp.StatusMsg = resp2.StatusMsg
	resp.StatusCode = resp2.StatusCode
	return resp, nil
}
func FeedVideoInfo(ctx context.Context, req *douyin.FeedRequest, userid int64) (*douyin.FeedResponse, error) {
	if time.Unix(req.LatestTime, 10).Year() > 9999 {
		req.LatestTime = time.Now().Unix()
	}
	resp, err := videoclient.Feed(ctx, (*video.FeedRequest)(req))
	if err != nil {
		err = errons.ConverterrtoErr(err)
	}
	var p []*douyin.Video
	if resp == nil {
		req.LatestTime = time.Now().Unix()
		resp, _ = videoclient.Feed(ctx, (*video.FeedRequest)(req))
	}
	for i := 0; i < len(resp.VideoList); i++ {
		user2, err := UserInfo(ctx, &user.UserInfoRequest{
			UserId: resp.VideoList[i].Author.Id,
			Token:  req.Token,
		})
		if err != nil {
			err = errons.ConverterrtoErr(err)
		}
		if userid != -1 {
			followinfo, _ := videoclient.IfLike(ctx, &video.DoesFavoriteRequest{UserId: userid, VideoId: resp.VideoList[i].Id})
			resp.VideoList[i].IsFavorite = followinfo.Iffollow
		}
		temp := douyin.Video{}
		temp.ID = resp.VideoList[i].Id
		temp.Title = resp.VideoList[i].Title
		temp.IsFavorite = resp.VideoList[i].IsFavorite
		temp.CommentCount = resp.VideoList[i].CommentCount
		temp.CoverURL = resp.VideoList[i].CoverUrl
		temp.PlayURL = resp.VideoList[i].PlayUrl
		temp.FavoriteCount = resp.VideoList[i].FavoriteCount
		auther := douyin.User{ID: user2.User.Id,
			Name:          user2.User.Name,
			FollowerCount: user2.User.FollowerCount,
			FollowCount:   user2.User.FollowCount,
			IsFollow:      user2.User.IsFollow,
		}
		if userid == auther.ID {
			auther.IsFollow = true
		}
		temp.Author = &auther
		p = append(p, &temp)
	}
	resp2 := &douyin.FeedResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		VideoList:  p,
		NextTime:   resp.NextTime,
	}
	return resp2, err
}

func VideoList(ctx context.Context, req *douyin.VideoPublishedListRequest, userid int64) (*douyin.VideoPublishedListResponse, error) {
	resp, err := videoclient.GetVideoList(ctx, &video.VideoPublishedListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		err = errons.ConverterrtoErr(err)
	}
	var p []*douyin.Video
	user2, err := UserInfo(ctx, &user.UserInfoRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	for i := 0; i < len(resp.VideoList); i++ {
		if err != nil {
			err = errons.ConverterrtoErr(err)
		}
		if userid != -1 {
			followinfo, _ := videoclient.IfLike(ctx, &video.DoesFavoriteRequest{UserId: userid, VideoId: resp.VideoList[i].Id})
			resp.VideoList[i].IsFavorite = followinfo.Iffollow
		}
		temp := douyin.Video{}
		temp.ID = resp.VideoList[i].Id
		temp.Title = resp.VideoList[i].Title
		temp.IsFavorite = resp.VideoList[i].IsFavorite
		temp.CommentCount = resp.VideoList[i].CommentCount
		temp.CoverURL = resp.VideoList[i].CoverUrl
		temp.PlayURL = resp.VideoList[i].PlayUrl
		temp.FavoriteCount = resp.VideoList[i].FavoriteCount
		auther := douyin.User{ID: user2.User.Id,
			Name:          user2.User.Name,
			FollowerCount: user2.User.FollowerCount,
			FollowCount:   user2.User.FollowCount,
			IsFollow:      user2.User.IsFollow,
		}
		temp.Author = &auther
		p = append(p, &temp)
	}
	resp2 := &douyin.VideoPublishedListResponse{StatusMsg: resp.StatusMsg, StatusCode: resp.StatusCode, VideoList: p}
	return resp2, nil
}

func CountChange(ctx context.Context, req *user.UserCountChangeRequest) (resp *user.UserCountChangeResponse, err error) {
	resp, _ = userclient.CountAdd(ctx, req)
	return resp, nil
}

// Likes
func LikeAction(ctx context.Context, req *video.FavoriteActionRequest) (resp *video.FavoriteActionResponse) {
	resp, _ = videoclient.LikeAction(ctx, req)
	return resp
}

func LikeVideosList(ctx context.Context, req *video.FavoriteVideoListRequest) (resp *video.FavoriteVideoListResponse) {
	resp, _ = videoclient.LikeVideos(ctx, req)
	for i := 0; i < len(resp.VideoList); i++ {
		user, _ := UserInfo(ctx, &user.UserInfoRequest{UserId: resp.VideoList[i].Id, Token: req.Token})
		resp.VideoList[i].Author = (*video.User)(user.User)
	}
	return resp
}

func CommentAction(ctx context.Context, req *video.CommentActionRequest) (resp *video.CommentActionResponse) {
	//resp = &video.CommentActionResponse{}
	resp, err := videoclient.CommentACtion(ctx, req)
	fmt.Println(err)
	temp, _ := UserInfo(ctx, &user.UserInfoRequest{UserId: req.Userid})
	resp.Comment.User = (*video.User)(temp.User)
	return resp
}
func CommentList(ctx context.Context, req *video.CommentListRequest) (resp *video.CommentListResponse) {
	resp, _ = videoclient.CommentList(ctx, req)
	for i := 0; i < len(resp.CommentList); i++ {
		temp, _ := UserInfo(ctx, &user.UserInfoRequest{
			UserId: resp.CommentList[i].User.Id,
		})
		resp.CommentList[i].User = (*video.User)(temp.User)
	}
	return resp
}

func RelationAction(ctx context.Context, req *user.RelationActionRequest) (resp *user.RelationActionResponse, err error) {
	resp, _ = userclient.RelationAction(ctx, req)
	return resp, nil
}
func FollowList(ctx context.Context, req *user.FollowListRequest) (resp *user.FollowListResponse, err error) {
	resp, _ = userclient.FollowList(ctx, req)
	return resp, nil
}
func FollowerList(ctx context.Context, req *user.FollowerListRequest) (resp *user.FollowerListResponse, err error) {
	resp, _ = userclient.FollowerList(ctx, req)
	return resp, nil
}
func FriendList(ctx context.Context, req *user.FriendListRequest) (resp *user.FriendListResponse, err error) {
	resp, _ = userclient.FriendList(ctx, req)
	return resp, nil
}
