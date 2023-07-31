package errons

import (
	_const "douyinv2/video/const"
	"fmt"
	"github.com/pkg/errors"
)

type Errorn struct {
	Errcode    int32
	Errmessage string
}

func (e Errorn) Error() string {
	return fmt.Sprintf("error code=%d,error_message=%s", e.Errcode, e.Errmessage)
}

func newErron(code int32, msg string) Errorn {
	return Errorn{
		Errcode:    code,
		Errmessage: msg,
	}
}

var (
	ServiceErr          = newErron(_const.ServiceErr, "Some Err happened in the server")
	AuthorizationFailed = newErron(_const.AuthorizationFailed, "the password or the username may be wrong")
	Successcode         = newErron(_const.Successcode, "success")
	Likeleardyexist     = newErron(_const.Likeexisted, "LikeAction have existed")
	Likenotexist        = newErron(_const.Likenotexist, "the moive did not likes")
)

func ConverterrtoErr(err error) Errorn {
	Err := Errorn{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.Errmessage = err.Error()
	return s
}
