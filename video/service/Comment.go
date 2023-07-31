package service

import (
	"context"
	"douyinv2/video/dal/db"
	"douyinv2/video/errons"
	"douyinv2/video/kitex_gen/video"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{
		ctx: ctx,
	}
}

func (s *CommentService) CommentAction(req *video.CommentActionRequest) (resp *video.CommentActionResponse) {
	//发表评论
	resp = &video.CommentActionResponse{}
	if req.ActionType == 1 {
		comment, err := db.NewComment(s.ctx, req.Userid, req.VideoId, req.CommentText)
		resp.StatusCode = err.Errcode
		resp.StatusMsg = err.Errmessage
		if err != errons.Successcode {
			resp.Comment = &video.Comment{}
			return resp
		}
		resp.Comment = comment
		resp.Comment.User = &video.User{Id: req.Userid}
	}
	if req.ActionType == 2 {
		err := db.DeleteComment(s.ctx, req.CommentId)
		resp.StatusCode = err.Errcode
		resp.StatusMsg = err.Errmessage
		resp.Comment = &video.Comment{}
	}
	return resp
}

func (s *CommentService) CommentList(req *video.CommentListRequest) (resp *video.CommentListResponse) {
	resp = &video.CommentListResponse{}
	videos, err := db.QueryComment(s.ctx, req.VideoId)
	resp.StatusCode = err.Errcode
	resp.StatusMsg = err.Errmessage
	if err != errons.Successcode {
		resp.CommentList = nil
		return resp
	}
	resp.CommentList = videos
	return resp
}
