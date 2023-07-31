package service

import (
	"context"
	"douyinv2/video/dal/db"
	"douyinv2/video/kitex_gen/video"
)

func Iflike(ctx context.Context, req *video.DoesFavoriteRequest) *video.DoesFavoriteResponse {
	likeinfo, _ := db.QueryLikeItem(ctx, &video.FavoriteActionRequest{ID: req.UserId, Video_ID: req.VideoId})
	resp := &video.DoesFavoriteResponse{}
	resp.Iffollow = false
	if likeinfo != nil {
		resp.Iffollow = true
	}
	return resp
}
