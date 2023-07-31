package service

import (
	"context"
	"douyinv2/user/dal/db"
	"douyinv2/user/errons"
	"douyinv2/user/kitex_gen/user"
)

type FollowService struct {
	ctx context.Context
}

func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{ctx: ctx}
}

func (s *FollowService) FollowAction(req *user.RelationActionRequest) (resp *user.RelationActionResponse, err2 error) {
	var err errons.Errorn
	if req.ActionType == 1 {
		err = db.CreateFollowItem(s.ctx, req.Fromid, req.Toid)
	}
	if req.ActionType == 2 {
		err = db.DeleteFollowItem(s.ctx, req.Fromid, req.Toid)
	}

	resp = &user.RelationActionResponse{StatusCode: err.Errcode, StatusMsg: err.Errmessage}
	return resp, nil
}
func (s *FollowService) FollowList(req *user.FollowListRequest) (resp *user.FollowListResponse, err2 error) {
	users, err := db.FollowList(s.ctx, req.Userid)
	resp = &user.FollowListResponse{
		UserList:   users,
		StatusMsg:  err.Errmessage,
		StatusCode: err.Errcode,
	}
	return resp, nil
}
func (s *FollowService) FollowerList(req *user.FollowerListRequest) (resp *user.FollowerListResponse, err2 error) {
	users, err := db.FollowerList(s.ctx, req.Userid)
	resp = &user.FollowerListResponse{
		UserList:   users,
		StatusMsg:  err.Errmessage,
		StatusCode: err.Errcode,
	}
	return resp, nil
}
func (s *FollowService) FriendList(req *user.FriendListRequest) (resp *user.FriendListResponse, err2 error) {
	friends, err := db.FriendList(s.ctx, req.ID)
	resp = &user.FriendListResponse{UserList: friends, StatusMsg: err.Errmessage, StatusCode: err.Errcode}
	return resp, nil
}
