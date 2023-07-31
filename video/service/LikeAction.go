package service

import (
	"context"
	"douyinv2/video/dal/db"
	"douyinv2/video/errons"
	"douyinv2/video/kitex_gen/video"
)

func LikeAction(ctx context.Context, req *video.FavoriteActionRequest) (resp *video.FavoriteActionResponse) {
	if req.ActionType == 1 {
		user, err := db.QueryLikeItem(ctx, req)
		if user != nil {
			err = errons.ConverterrtoErr(err)
			resp = &video.FavoriteActionResponse{
				StatusCode: err.Errcode,
				StatusMsg:  err.Errmessage,
			}
			return resp
		}
		err = db.CreaetLikeItem(ctx, req)

		resp = &video.FavoriteActionResponse{StatusCode: err.Errcode, StatusMsg: err.Errmessage}
	}
	if req.ActionType == 2 {
		user, err := db.QueryLikeItem(ctx, req)
		if user == nil {
			err = errons.ConverterrtoErr(err)
			resp = &video.FavoriteActionResponse{
				StatusCode: err.Errcode,
				StatusMsg:  err.Errmessage,
			}
			return resp
		}
		err = db.DeleteLikeItem(ctx, req)
		resp = &video.FavoriteActionResponse{StatusCode: err.Errcode, StatusMsg: err.Errmessage}
	}
	return resp
}
