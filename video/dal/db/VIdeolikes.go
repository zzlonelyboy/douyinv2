package db

import (
	"context"
	"douyinv2/video/dal/model"
	"douyinv2/video/errons"
	"douyinv2/video/kitex_gen/video"
)

func CreaetLikeItem(ctx context.Context, req *video.FavoriteActionRequest) errons.Errorn {
	err := errons.Successcode
	likeitem := &model.VideoLikes{VideoID: req.Video_ID, UserID: req.ID}
	err2 := Q.WithContext(ctx).VideoLikes.Create(likeitem)
	if err2 != nil {
		err = errons.ConverterrtoErr(err2)
	}
	return err
}

func QueryLikeItem(ctx context.Context, req *video.FavoriteActionRequest) (*model.VideoLikes, errons.Errorn) {
	err := errons.Successcode
	user, err2 := Q.WithContext(ctx).VideoLikes.Where(Q.VideoLikes.UserID.Eq(req.ID)).Where(Q.VideoLikes.VideoID.Eq(req.Video_ID)).First()
	if err2 != nil {
		err = errons.ConverterrtoErr(err2)
	}
	return user, err
}
func DeleteLikeItem(ctx context.Context, req *video.FavoriteActionRequest) errons.Errorn {
	err := errons.Successcode
	_, err2 := Q.WithContext(ctx).VideoLikes.Where(Q.VideoLikes.UserID.Eq(req.ID)).Where(Q.VideoLikes.VideoID.Eq(req.Video_ID)).Delete()
	if err2 != nil {
		err = errons.ConverterrtoErr(err2)
	}
	return err
}
func LikeVideoList(ctx context.Context, req *video.FavoriteVideoListRequest) ([]int64, errons.Errorn) {
	err := errons.Successcode
	var video []int64
	likelist, err2 := Q.WithContext(ctx).VideoLikes.Where(Q.VideoLikes.UserID.Eq(req.UserId)).Find()
	for i := 0; i < len(likelist); i++ {
		video = append(video, likelist[i].VideoID)
	}
	if err2 != nil {
		err = errons.ConverterrtoErr(err2)
	}
	return video, err
}
