package validation

import requestmodels "github.com/ZiplEix/pixel-espion/request_models"

func NewSpy(req requestmodels.NewSpyRequest) error {
	return validate.Struct(req)
}
