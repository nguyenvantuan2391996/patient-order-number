package admin

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/nguyenvantuan2391996/patient-order-number/base_common/comoutput"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/database/entities"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/admin/models"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/repository"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

type Admin struct {
	accountRepo repository.IAccountRepositoryInterface
}

func NewAdminService(accountRepo repository.IAccountRepositoryInterface) *Admin {
	return &Admin{
		accountRepo: accountRepo,
	}
}

func (as *Admin) CreateAccount(ctx context.Context, input *models.AccountInput) (*comoutput.BaseOutput, error) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginTask, "CreateAccount", input))

	account, err := as.accountRepo.GetByQueries(ctx, map[string]interface{}{
		"user_name": input.UserName,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf(constants.FormatGetEntityErr, "account", err)
		return nil, err
	}

	if account != nil {
		return nil, constants.AccountExisted
	}

	err = as.accountRepo.Create(ctx, &entities.Account{
		UserName: input.UserName,
		Password: input.Password,
		Status:   1,
		Role:     0,
	})
	if err != nil {
		logrus.Errorf(constants.FormatCreateEntityErr, "accounts", err)
		return nil, err
	}

	return &comoutput.BaseOutput{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"user_name": input.UserName,
		},
	}, nil
}

func (as *Admin) UpdateAccount(ctx context.Context, input *models.AccountUpdateInput) (*comoutput.BaseOutput, error) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginTask, "UpdateAccount", input))

	account, err := as.accountRepo.GetByQueries(ctx, map[string]interface{}{
		"id": input.UserID,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf(constants.FormatGetEntityErr, "account", err)
		return nil, err
	}

	if account == nil {
		return nil, constants.AccountNotExisted
	}

	err = as.accountRepo.UpdateWithMap(ctx, account, map[string]interface{}{
		"user_name": input.UserName,
		"password":  input.Password,
		"status":    input.Status,
	})
	if err != nil {
		logrus.Errorf(constants.FormatUpdateEntityErr, "accounts", err)
		return nil, err
	}

	return &comoutput.BaseOutput{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"user_id":   input.UserID,
			"user_name": input.UserName,
		},
	}, nil
}

func (as *Admin) DeleteAccount(ctx context.Context, input *models.DeleteAccountInput) (*comoutput.BaseOutput, error) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginTask, "DeleteAccount", input))

	account, err := as.accountRepo.GetByQueries(ctx, map[string]interface{}{
		"id": input.UserID,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf(constants.FormatGetEntityErr, "account", err)
		return nil, err
	}

	if account == nil {
		return nil, constants.AccountNotExisted
	}

	err = as.accountRepo.Delete(ctx, account)
	if err != nil {
		logrus.Errorf(constants.FormatDeleteEntityErr, "account", err)
		return nil, err
	}

	return &comoutput.BaseOutput{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"user_id": input.UserID,
		},
	}, nil
}
