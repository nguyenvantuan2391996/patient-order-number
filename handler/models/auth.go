package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/utils"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/auth/models"
)

type LoginRequest struct {
	UserName string `form:"user_name"`
	Password string `form:"password"`
}

func (r *LoginRequest) ToLoginInput() *models.LoginInput {
	out := &models.LoginInput{}
	if r == nil {
		return out
	}

	out.UserName = r.UserName
	out.Password = utils.EncodePasswordSHA1(r.Password + r.UserName)

	return out
}

func (r *LoginRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.UserName, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}
