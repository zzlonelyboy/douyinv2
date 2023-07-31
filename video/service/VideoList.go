package service

import (
	"context"
	"douyinv2/user/errons"
	"douyinv2/video/dal/db"
	"douyinv2/video/kitex_gen/video"
)

type VideoListService struct {
	ctx context.Context
}

func NewVideoListService(ctx context.Context) *VideoListService {
	return &VideoListService{
		ctx: ctx,
	}
}

func (s *VideoListService) GetVideoList(request *video.VideoPublishedListRequest) (*video.VideoPublishedListResponse, error) {
	VidoeList, err := db.GetVideoList(s.ctx, request.UserId)
	if err != errons.Successcode {
		err = errons.ConverterrtoErr(err)
	}
	resp := &video.VideoPublishedListResponse{
		StatusCode: err.Errcode,
		StatusMsg:  err.Errmessage,
		VideoList:  VidoeList,
	}
	return resp, nil
}
