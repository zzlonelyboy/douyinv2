package main

import (
	"context"
	user "douyinv2/user/kitex_gen/user"
	"douyinv2/user/service"
	"fmt"
)

// ApiServiceImpl implements the last service interface defined in the IDL.
type ApiServiceImpl struct{}

// RegisterUser implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) RegisterUser(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, rerr error) {
	// TODO: Your code here...
	fmt.Println("register")
	//name = req.GetUsername()
	//passwrod = req.GetPassword()
	id, err := service.NewUserRegisterService(ctx).UserRegister(req)
	resp = new(user.UserRegisterResponse)
	resp.UserId = id
	resp.StatusCode = err.Errcode
	resp.StatusMsg = err.Errmessage
	resp.Token = "yes"
	return resp, nil
}

// LoginUser implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) LoginUser(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, rerr error) {
	// TODO: Your code here...
	fmt.Println("login")
	//name = req.GetUsername()
	//passwrod = req.GetPassword()
	id, err := service.NewUserLoginService(ctx).UserLogin(req)
	resp = new(user.UserLoginResponse)
	resp.UserId = id
	resp.StatusCode = err.Errcode
	resp.StatusMsg = err.Errmessage
	resp.Token = "yes"
	return resp, nil
}

// UserInfo implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, reserr error) {
	// TODO: Your code here...
	fmt.Println("info")
	getuser, err := service.NewUserInfoService(ctx).UserInfo(req)
	resp = new(user.UserInfoResponse)
	resp.StatusCode = err.Errcode
	resp.StatusMsg = err.Errmessage
	resp.User = &getuser
	return resp, nil
}

// CountAdd implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) CountAdd(ctx context.Context, req *user.UserCountChangeRequest) (resp *user.UserCountChangeResponse, err error) {
	// TODO: Your code here...
	err2 := service.UserCountChanged(ctx, req)
	resp.StatusCode = err2.Errcode
	resp.StatusMsg = err2.Errmessage
	return resp, nil
}

// RelationAction implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) RelationAction(ctx context.Context, req *user.RelationActionRequest) (resp *user.RelationActionResponse, err error) {
	// TODO: Your code here...
	resp, err = service.NewFollowService(ctx).FollowAction(req)
	return resp, err
}

// FollowList implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) FollowList(ctx context.Context, req *user.FollowListRequest) (resp *user.FollowListResponse, err error) {
	// TODO: Your code here...
	resp, err = service.NewFollowService(ctx).FollowList(req)
	return resp, err
}

// FollowerList implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) FollowerList(ctx context.Context, req *user.FollowerListRequest) (resp *user.FollowerListResponse, err error) {
	// TODO: Your code here...
	resp, err = service.NewFollowService(ctx).FollowerList(req)
	return resp, err
}

// FriendList implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) FriendList(ctx context.Context, req *user.FriendListRequest) (resp *user.FriendListResponse, err error) {
	// TODO: Your code here...
	resp, err = service.NewFollowService(ctx).FriendList(req)
	return resp, err
}
