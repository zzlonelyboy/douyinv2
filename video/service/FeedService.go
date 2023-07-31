package service

import (
	"context"
	"douyinv2/video/dal/db"
	"douyinv2/video/kitex_gen/video"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) Feed(req *video.FeedRequest) (*video.FeedResponse, error) {
	videolist, lasttime, err := db.GetFeedVideo(s.ctx, req.LatestTime)
	resp := &video.FeedResponse{
		StatusCode: err.Errcode,
		StatusMsg:  err.Errmessage,
		VideoList:  videolist,
		NextTime:   lasttime,
	}
	return resp, nil
}
