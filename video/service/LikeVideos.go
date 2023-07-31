package service

import (
	"context"
	"douyinv2/video/dal/db"
	"douyinv2/video/errons"
	"douyinv2/video/kitex_gen/video"
)

func LikeVideos(ctx context.Context, req *video.FavoriteVideoListRequest) (*video.FavoriteVideoListResponse, error) {
	var videos []*video.Video
	video_id, err := db.LikeVideoList(ctx, req)
	if err != errons.Successcode {
		return &video.FavoriteVideoListResponse{StatusCode: err.Errcode, StatusMsg: err.Errmessage}, nil
	}
	for i := 0; i < len(video_id); i++ {
		temp, err2 := db.QueryVideos(ctx, video_id[i])
		if err2 != nil {
			err = errons.ConverterrtoErr(err2)
			return &video.FavoriteVideoListResponse{StatusCode: err.Errcode, StatusMsg: err.Errmessage}, nil
		}
		videos = append(videos, temp)
	}
	return &video.FavoriteVideoListResponse{VideoList: videos, StatusMsg: err.Errmessage, StatusCode: err.Errcode}, nil
}
