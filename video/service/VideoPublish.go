package service

import (
	"context"
	"douyinv2/video/dal/db"
	"douyinv2/video/dal/model"
	"douyinv2/video/errons"
	"douyinv2/video/structers"
)

type CreaetVideoService struct {
	ctx context.Context
}

func NewCraetVideoService(ctx context.Context) *CreaetVideoService {
	return &CreaetVideoService{ctx: ctx}
}
func (s CreaetVideoService) CreateVideo(req *structers.VideoInfo) errons.Errorn {
	//http框架由jwt token获取用户ID数据
	err := db.CreaetVideo(s.ctx, []*model.Videos{{PostUserID: req.PosterID, Playurl: req.Playurl, Coverurl: req.Coverurl, Title: req.Title}})
	if err != nil {
		return errons.ConverterrtoErr(err)
	}
	return errons.Successcode
}
