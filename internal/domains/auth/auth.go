package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/comoutput"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/database/entities"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/auth/models"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/repository"
	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

type Auth struct {
	accountRepo repository.IAccountRepositoryInterface
}

func NewAuthService(accountRepo repository.IAccountRepositoryInterface) *Auth {
	return &Auth{
		accountRepo: accountRepo,
	}
}

func (as *Auth) Login(ctx context.Context, input *models.LoginInput) (*comoutput.BaseOutput, error) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginTask, "Login", input))

	account, err := as.accountRepo.GetByQueries(ctx, map[string]interface{}{
		"user_name": input.UserName,
		"password":  input.Password,
		"status":    1,
	})
	if err != nil {
		logrus.Errorf(constants.FormatGetEntityErr, "account", err)
		return nil, err
	}

	if account == nil {
		return nil, constants.AccountInvalid
	}

	token, err := as.generateToken(account)
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "generateToken", err)
		return nil, err
	}

	return &comoutput.BaseOutput{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"user_name": account.UserName,
			"token":     token,
			"expire":    constants.ExpiredTime * 3600,
		},
	}, nil
}

func (as *Auth) generateToken(account *entities.Account) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":        "account",
			"iat":        time.Now().Unix(),
			"account_id": account.ID,
			"user_name":  account.UserName,
			"status":     account.Status,
			"role":       account.Role,
			"exp":        time.Now().Add(time.Hour * constants.ExpiredTime).Unix(),
		})

	token, err := claims.SignedString([]byte(viper.GetString("PRIVATE_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil
}
