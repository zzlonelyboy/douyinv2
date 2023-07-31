package db

import (
	"context"
	"douyinv2/user/errons"
	"douyinv2/video/dal/model"
	"douyinv2/video/kitex_gen/video"
	"time"
)

func CreaetVideo(ctx context.Context, video []*model.Videos) error {
	err := Q.WithContext(ctx).Videos.Create(video...)
	return err
}

func GetFeedVideo(ctx context.Context, lasttime int64) ([]*video.Video, int64, errons.Errorn) {
	if lasttime == 0 {
		lasttime = time.Now().Unix()
	}
	timestr := time.Unix(lasttime, 0)
	var Videos []*model.Videos
	var errn errons.Errorn
	var retVideos []*video.Video
	if timestr.Year() == 55535 {
		timestr = time.Now()
	}
	Videos, err := Q.WithContext(ctx).Videos.Where(Q.Videos.CreatedAt.Lt(timestr)).Limit(30).Order(Q.Videos.CreatedAt.Desc()).Find()
	errn = errons.Successcode
	if err != nil {
		errn = errons.ConverterrtoErr(err)
	}
	for i := 0; i < len(Videos); i++ {
		auther := &video.User{Id: Videos[i].PostUserID}
		likecount, _ := Q.WithContext(ctx).VideoLikes.Where(Q.VideoLikes.VideoID.Eq(Videos[i].ID)).Count()
		commentcount, _ := Q.WithContext(ctx).VideoComments.Where(Q.VideoComments.VideoID.Eq(Videos[i].ID)).Count()
		temp := &video.Video{
			Id:       Videos[i].ID,
			PlayUrl:  Videos[i].Playurl,
			CoverUrl: Videos[i].Coverurl,
			//后期修改
			Author:        auther,
			FavoriteCount: likecount,
			CommentCount:  commentcount,
			IsFavorite:    false,
			Title:         Videos[i].Title,
		}
		retVideos = append(retVideos, temp)
	}

	var returntime int64
	returntime = time.Now().Unix()
	if len(Videos)-1 >= 0 {
		returntime = Videos[len(Videos)-1].CreatedAt.Unix()
	}
	return retVideos, returntime, errn
}

func GetVideoList(ctx context.Context, userid int64) ([]*video.Video, errons.Errorn) {
	var Videos []*model.Videos
	var errn errons.Errorn
	var retVideos []*video.Video
	Videos, err := Q.WithContext(ctx).Videos.Where(Q.Videos.PostUserID.Eq(userid)).Find()
	errn = errons.Successcode
	if err != nil {
		errn = errons.ConverterrtoErr(err)
	}
	for i := 0; i < len(Videos); i++ {
		likecount, _ := Q.WithContext(ctx).VideoLikes.Where(Q.VideoLikes.VideoID.Eq(Videos[i].ID)).Count()
		commentcount, _ := Q.WithContext(ctx).VideoComments.Where(Q.VideoComments.VideoID.Eq(Videos[i].ID)).Count()
		auther := &video.User{Id: Videos[i].PostUserID}
		temp := &video.Video{
			Id:       Videos[i].ID,
			PlayUrl:  Videos[i].Playurl,
			CoverUrl: Videos[i].Coverurl,
			//后期修改
			Author:        auther,
			FavoriteCount: likecount,
			CommentCount:  commentcount,
			IsFavorite:    false,
			Title:         Videos[i].Title,
		}
		retVideos = append(retVideos, temp)
	}
	return retVideos, errn
}

func QueryVideos(ctx context.Context, video_id int64) (*video.Video, error) {
	Video, err := Q.WithContext(ctx).Videos.Where(Q.Videos.ID.Eq(video_id)).First()
	likecount, _ := Q.WithContext(ctx).VideoLikes.Where(Q.VideoLikes.VideoID.Eq(video_id)).Count()
	if err != nil {
		return nil, err
	}
	temp := &video.Video{
		Id:       video_id,
		PlayUrl:  Video.Playurl,
		CoverUrl: Video.Coverurl,
		Author: &video.User{
			Id: Video.PostUserID,
		},
		FavoriteCount: likecount,
		CommentCount:  0,
		IsFavorite:    true,
		Title:         Video.Title,
	}
	return temp, nil
}
