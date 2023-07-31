package mw

import (
	"context"
	douyin "douyinv2/biz/model/douyinv2"
	"douyinv2/biz/rpc"
	"douyinv2/user/errons"
	"douyinv2/user/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	jwt2 "github.com/form3tech-oss/jwt-go"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

type Myclaims struct {
	Id int64 `json:"id"`
	jwt2.StandardClaims
}

var Jwtmiddleware *jwt.HertzJWTMiddleware

func InitJwtMiddleWare() {
	Jwtmiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "doushen",
		Key:         []byte("douyintest"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &user.UserLoginResponse{
				UserId: int64(claims["id"].(float64)),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			v, ok := data.(*user.UserLoginResponse)
			if ok {
				return jwt.MapClaims{
					"id": v.UserId,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req douyin.UserLoginRequest
			err = c.BindAndValidate(&req)
			if err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			user, err2 := rpc.UserLogin(ctx, &user.UserLoginRequest{
				Username: req.Username,
				Password: req.Password,
			})
			if err != nil {
				return nil, err2
			}
			c.Set("user_id", user.UserId)
			return user, err2
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, message string, time time.Time) {
			userid := c.Value("user_id").(int64)
			c.JSON(http.StatusOK, utils.H{
				"status_code": errons.Successcode.Errcode,
				"status_msg":  errons.Successcode.Errmessage,
				"user_id":     userid,
				"token":       message,
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": errons.Successcode.Errcode,
				"status_msg":  errons.Successcode.Errmessage,
				"user_id":     -1,
				"token":       message,
			})
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errons.Errorn:
				return t.Errmessage
			default:
				return t.Error()
			}
		},
	})
}
