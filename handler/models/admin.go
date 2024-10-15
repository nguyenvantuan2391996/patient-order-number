package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/utils"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/admin/models"
)

type AccountRequest struct {
	UserName string `form:"user_name"`
	Password string `form:"password"`
}

type AccountUpdateRequest struct {
	UserName string `form:"user_name"`
	Password string `form:"password"`
	Status   int    `form:"status"`
}

type DeleteAccountRequest struct {
}

func (r *AccountRequest) ToAccountInput() *models.AccountInput {
	out := &models.AccountInput{}
	if r == nil {
		return out
	}

	out.UserName = r.UserName
	out.Password = utils.EncodePasswordSHA1(r.Password + r.UserName)

	return out
}

func (r *AccountUpdateRequest) ToAccountUpdateInput(userID int64) *models.AccountUpdateInput {
	out := &models.AccountUpdateInput{}
	if r == nil {
		return out
	}

	out.UserID = userID
	out.UserName = r.UserName
	out.Password = utils.EncodePasswordSHA1(r.Password + r.UserName)
	out.Status = r.Status

	return out
}

func (r *DeleteAccountRequest) ToDeleteAccountInput(userID int64) *models.DeleteAccountInput {
	out := &models.DeleteAccountInput{}
	if r == nil {
		return out
	}

	out.UserID = userID

	return out
}

func (r *AccountRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.UserName, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

func (r *AccountUpdateRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.UserName, validation.Required),
		validation.Field(&r.Password, validation.Required),
		validation.Field(&r.Status, validation.Required, validation.In(0, 1)),
	)
}
