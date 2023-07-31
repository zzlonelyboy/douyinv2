package main

import (
	"context"
	video "douyinv2/video/kitex_gen/video"
	"douyinv2/video/service"
	"douyinv2/video/structers"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// VideoPublish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublish(ctx context.Context, req *video.VideoPublishRequest) (resp *video.VideoPublishResponse, err error) {
	// TODO: Your code here...
	//cover := "https://t9.baidu.com/it/u=2881085827,2025487051&fm=218&app=126&size=f242,150&n=0&f=JPEG&fmt=auto?s=DF87E0054A33C3DC1C10218D01005082&sec=1690390800&t=f9b6d64b85ee17bd5cbb44b965ae6ba9"
	terr := service.NewCraetVideoService(ctx).CreateVideo(&structers.VideoInfo{Title: req.Title, Playurl: req.Filepath, Coverurl: req.Cover, PosterID: req.ID})
	resp = &video.VideoPublishResponse{}
	resp.StatusCode = terr.Errcode
	resp.StatusMsg = terr.Errmessage
	return resp, nil
}

// GetVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoList(ctx context.Context, req *video.VideoPublishedListRequest) (resp *video.VideoPublishedListResponse, err error) {
	resp, err = service.NewVideoListService(ctx).GetVideoList(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (*video.FeedResponse, error) {
	// TODO: Your code here...
	resp, err := service.NewFeedService(ctx).Feed(req)
	return resp, err
}

// LikeAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) LikeAction(ctx context.Context, req *video.FavoriteActionRequest) (resp *video.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	resp = service.LikeAction(ctx, req)
	return resp, nil
}

// LikeVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) LikeVideos(ctx context.Context, req *video.FavoriteVideoListRequest) (resp *video.FavoriteVideoListResponse, err error) {
	// TODO: Your code here...
	resp, _ = service.LikeVideos(ctx, req)
	return resp, nil
}

// IfLike implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) IfLike(ctx context.Context, req *video.DoesFavoriteRequest) (resp *video.DoesFavoriteResponse, err error) {
	// TODO: Your code here...
	resp = &video.DoesFavoriteResponse{}
	resp = service.Iflike(ctx, req)
	return resp, nil
}

// CommentACtion implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentACtion(ctx context.Context, req *video.CommentActionRequest) (resp *video.CommentActionResponse, err error) {
	// TODO: Your code here...
	//resp = &video.CommentActionResponse{}
	resp = service.NewCommentService(ctx).CommentAction(req)
	return resp, nil
}

// CommentListResponse implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentList(ctx context.Context, req *video.CommentListRequest) (resp *video.CommentListResponse, err error) {
	// TODO: Your code here...
	resp = &video.CommentListResponse{}
	resp = service.NewCommentService(ctx).CommentList(req)
	return resp, nil
}
