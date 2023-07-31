// Code generated by Kitex v0.6.1. DO NOT EDIT.

package apiservice

import (
	"context"
	user "douyinv2/user/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return apiServiceServiceInfo
}

var apiServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ApiService"
	handlerType := (*user.ApiService)(nil)
	methods := map[string]kitex.MethodInfo{
		"RegisterUser":   kitex.NewMethodInfo(registerUserHandler, newApiServiceRegisterUserArgs, newApiServiceRegisterUserResult, false),
		"LoginUser":      kitex.NewMethodInfo(loginUserHandler, newApiServiceLoginUserArgs, newApiServiceLoginUserResult, false),
		"UserInfo":       kitex.NewMethodInfo(userInfoHandler, newApiServiceUserInfoArgs, newApiServiceUserInfoResult, false),
		"CountAdd":       kitex.NewMethodInfo(countAddHandler, newApiServiceCountAddArgs, newApiServiceCountAddResult, false),
		"RelationAction": kitex.NewMethodInfo(relationActionHandler, newApiServiceRelationActionArgs, newApiServiceRelationActionResult, false),
		"FollowList":     kitex.NewMethodInfo(followListHandler, newApiServiceFollowListArgs, newApiServiceFollowListResult, false),
		"FollowerList":   kitex.NewMethodInfo(followerListHandler, newApiServiceFollowerListArgs, newApiServiceFollowerListResult, false),
		"FriendList":     kitex.NewMethodInfo(friendListHandler, newApiServiceFriendListArgs, newApiServiceFriendListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.ApiServiceRegisterUserArgs)
	realResult := result.(*user.ApiServiceRegisterUserResult)
	success, err := handler.(user.ApiService).RegisterUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newApiServiceRegisterUserArgs() interface{} {
	return user.NewApiServiceRegisterUserArgs()
}

func newApiServiceRegisterUserResult() interface{} {
	return user.NewApiServiceRegisterUserResult()
}

func loginUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.ApiServiceLoginUserArgs)
	realResult := result.(*user.ApiServiceLoginUserResult)
	success, err := handler.(user.ApiService).LoginUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newApiServiceLoginUserArgs() interface{} {
	return user.NewApiServiceLoginUserArgs()
}

func newApiServiceLoginUserResult() interface{} {
	return user.NewApiServiceLoginUserResult()
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.ApiServiceUserInfoArgs)
	realResult := result.(*user.ApiServiceUserInfoResult)
	success, err := handler.(user.ApiService).UserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newApiServiceUserInfoArgs() interface{} {
	return user.NewApiServiceUserInfoArgs()
}

func newApiServiceUserInfoResult() interface{} {
	return user.NewApiServiceUserInfoResult()
}

func countAddHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.ApiServiceCountAddArgs)
	realResult := result.(*user.ApiServiceCountAddResult)
	success, err := handler.(user.ApiService).CountAdd(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newApiServiceCountAddArgs() interface{} {
	return user.NewApiServiceCountAddArgs()
}

func newApiServiceCountAddResult() interface{} {
	return user.NewApiServiceCountAddResult()
}

func relationActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.ApiServiceRelationActionArgs)
	realResult := result.(*user.ApiServiceRelationActionResult)
	success, err := handler.(user.ApiService).RelationAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newApiServiceRelationActionArgs() interface{} {
	return user.NewApiServiceRelationActionArgs()
}

func newApiServiceRelationActionResult() interface{} {
	return user.NewApiServiceRelationActionResult()
}

func followListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.ApiServiceFollowListArgs)
	realResult := result.(*user.ApiServiceFollowListResult)
	success, err := handler.(user.ApiService).FollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newApiServiceFollowListArgs() interface{} {
	return user.NewApiServiceFollowListArgs()
}

func newApiServiceFollowListResult() interface{} {
	return user.NewApiServiceFollowListResult()
}

func followerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.ApiServiceFollowerListArgs)
	realResult := result.(*user.ApiServiceFollowerListResult)
	success, err := handler.(user.ApiService).FollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newApiServiceFollowerListArgs() interface{} {
	return user.NewApiServiceFollowerListArgs()
}

func newApiServiceFollowerListResult() interface{} {
	return user.NewApiServiceFollowerListResult()
}

func friendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.ApiServiceFriendListArgs)
	realResult := result.(*user.ApiServiceFriendListResult)
	success, err := handler.(user.ApiService).FriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newApiServiceFriendListArgs() interface{} {
	return user.NewApiServiceFriendListArgs()
}

func newApiServiceFriendListResult() interface{} {
	return user.NewApiServiceFriendListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) RegisterUser(ctx context.Context, req *user.UserRegisterRequest) (r *user.UserRegisterResponse, err error) {
	var _args user.ApiServiceRegisterUserArgs
	_args.Req = req
	var _result user.ApiServiceRegisterUserResult
	if err = p.c.Call(ctx, "RegisterUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) LoginUser(ctx context.Context, req *user.UserLoginRequest) (r *user.UserLoginResponse, err error) {
	var _args user.ApiServiceLoginUserArgs
	_args.Req = req
	var _result user.ApiServiceLoginUserResult
	if err = p.c.Call(ctx, "LoginUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, req *user.UserInfoRequest) (r *user.UserInfoResponse, err error) {
	var _args user.ApiServiceUserInfoArgs
	_args.Req = req
	var _result user.ApiServiceUserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CountAdd(ctx context.Context, req *user.UserCountChangeRequest) (r *user.UserCountChangeResponse, err error) {
	var _args user.ApiServiceCountAddArgs
	_args.Req = req
	var _result user.ApiServiceCountAddResult
	if err = p.c.Call(ctx, "CountAdd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RelationAction(ctx context.Context, req *user.RelationActionRequest) (r *user.RelationActionResponse, err error) {
	var _args user.ApiServiceRelationActionArgs
	_args.Req = req
	var _result user.ApiServiceRelationActionResult
	if err = p.c.Call(ctx, "RelationAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowList(ctx context.Context, req *user.FollowListRequest) (r *user.FollowListResponse, err error) {
	var _args user.ApiServiceFollowListArgs
	_args.Req = req
	var _result user.ApiServiceFollowListResult
	if err = p.c.Call(ctx, "FollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowerList(ctx context.Context, req *user.FollowerListRequest) (r *user.FollowerListResponse, err error) {
	var _args user.ApiServiceFollowerListArgs
	_args.Req = req
	var _result user.ApiServiceFollowerListResult
	if err = p.c.Call(ctx, "FollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FriendList(ctx context.Context, req *user.FriendListRequest) (r *user.FriendListResponse, err error) {
	var _args user.ApiServiceFriendListArgs
	_args.Req = req
	var _result user.ApiServiceFriendListResult
	if err = p.c.Call(ctx, "FriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}