package service

import (
	"context"
	"crypto/md5"
	"douyinv2/user/dal/db"
	"douyinv2/user/dal/model"
	"douyinv2/user/errons"
	"douyinv2/user/kitex_gen/user"
	"fmt"
	"io"
)

type UserRegisterService struct {
	ctx context.Context
}

func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{ctx: ctx}
}
func (s *UserRegisterService) UserRegister(req *user.UserRegisterRequest) (id int64, reserr errons.Errorn) {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return -1, errons.ConverterrtoErr(err)
	}
	if len(users) != 0 {
		return -1, errons.UserAleardyexist
	}
	h := md5.New()
	_, err = io.WriteString(h, req.Username)
	if err != nil {
		return -1, errons.ConverterrtoErr(err)
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	err = db.CreateUser(s.ctx, []*model.Userlogin{{Username: req.Username, Password: password}})
	if err != nil {
		return -1, errons.ConverterrtoErr(err)
	}
	user, err := db.UserLogin(s.ctx, req.Username, password)
	if err != nil {
		return -1, errons.ConverterrtoErr(err)
	}
	return user.ID, errons.Successcode
}
