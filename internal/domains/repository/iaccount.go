package repository

import (
	"context"

	"github.com/nguyenvantuan2391996/patient-order-number/base_common/database/entities"
)

//go:generate mockgen -package=repository -destination=iaccount_mock.go -source=iaccount.go

type IAccountRepositoryInterface interface {
	GetByQueries(ctx context.Context, queries map[string]interface{}) (*entities.Account, error)
}
