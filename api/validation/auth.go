package validation

import (
	requestmodels "github.com/ZiplEix/pixel-espion/request_models"
)

func Register(req requestmodels.RegisterReq) error {
	return validate.Struct(req)
}

func Login(req requestmodels.LoginReq) error {
	return validate.Struct(req)
}
