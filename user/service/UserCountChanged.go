package service

import (
	"context"
	"douyinv2/user/dal/db"
	"douyinv2/user/errons"
	"douyinv2/user/kitex_gen/user"
)

func UserCountChanged(ctx context.Context, req *user.UserCountChangeRequest) errons.Errorn {
	err := db.CountChange(ctx, req.UserId, req.Filed, req.Ifadd)
	if err != nil {
		return errons.ConverterrtoErr(err)
	}
	return errons.Successcode
}
