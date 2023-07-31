package errons

import (
	_const "douyinv2/user/const"
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
	UserAleardyexist    = newErron(_const.UserAleardyexist, "The Username has existed")
	ServiceErr          = newErron(_const.ServiceErr, "Some Err happened in the server")
	AuthorizationFailed = newErron(_const.AuthorizationFailed, "the password or the username may be wrong")
	Successcode         = newErron(_const.Successcode, "success")
	paramcode           = newErron(_const.ParamErr, "the params may be wrong")
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
