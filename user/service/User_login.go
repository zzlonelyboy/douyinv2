package service

import (
	"context"
	"crypto/md5"
	"douyinv2/user/dal/db"
	"douyinv2/user/errons"
	"douyinv2/user/kitex_gen/user"
	"fmt"
	"io"
)

type UserLoginService struct {
	ctx context.Context
}

func NewUserLoginService(ctx context.Context) *UserLoginService {
	return &UserLoginService{ctx: ctx}
}
func (s *UserLoginService) UserLogin(req *user.UserLoginRequest) (id int64, reerr errons.Errorn) {
	h := md5.New()
	_, err := io.WriteString(h, req.Username)
	if err != nil {
		return -1, errons.ConverterrtoErr(err)
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	user, err := db.UserLogin(s.ctx, req.Username, password)
	if err != nil {
		return -1, errons.ConverterrtoErr(err)
	}
	return user.ID, errons.Successcode
}
