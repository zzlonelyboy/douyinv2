package db

import (
	"context"
	"douyinv2/video/dal/model"
	"douyinv2/video/errons"
	"douyinv2/video/kitex_gen/video"
)

func NewComment(ctx context.Context, user_id int64, video_id int64, comment_text string) (resp *video.Comment, err2 errons.Errorn) {
	Comment := []*model.VideoComments{
		{UserID: user_id,
			VideoID:    video_id,
			CommentCon: comment_text},
	}
	resp = &video.Comment{}
	err := Q.WithContext(ctx).VideoComments.Create(Comment...)
	for _, comment := range Comment {
		resp.Id = comment.ID
		resp.Content = comment.CommentCon
		resp.LikeCount = 0
		resp.TeaseCount = 0
		resp.CreateDate = comment.CreatedAt.Format("04-02")
	}
	if err != nil {
		return &video.Comment{}, errons.ConverterrtoErr(err)
	}
	return resp, errons.Successcode
}

func QueryComment(ctx context.Context, vido_id int64) (resp []*video.Comment, err2 errons.Errorn) {
	comments, err := Q.WithContext(ctx).VideoComments.Where(Q.VideoComments.VideoID.Eq(vido_id)).Find()
	if err != nil {
		return nil, errons.ConverterrtoErr(err)
	}
	var List []*video.Comment
	for i := 0; i < len(comments); i++ {
		auther := &video.User{Id: comments[i].UserID}
		temp := &video.Comment{
			Id:         comments[i].ID,
			Content:    comments[i].CommentCon,
			CreateDate: comments[i].CreatedAt.Format("04-02"),
			User:       auther,
		}
		List = append(List, temp)
	}
	return List, errons.Successcode
}

func DeleteComment(ctx context.Context, comment_id int64) errons.Errorn {
	_, err := Q.WithContext(ctx).VideoComments.Where(Q.VideoComments.ID.Eq(comment_id)).Delete()
	if err != nil {
		return errons.ConverterrtoErr(err)
	}
	return errons.Successcode
}
