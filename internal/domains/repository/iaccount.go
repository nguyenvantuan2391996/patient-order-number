package repository

import (
	"context"

	"github.com/nguyenvantuan2391996/patient-order-number/base_common/database/entities"
)

//go:generate mockgen -package=repository -destination=iaccount_mock.go -source=iaccount.go

type IAccountRepositoryInterface interface {
	Create(ctx context.Context, record *entities.Account) error
	GetByQueries(ctx context.Context, queries map[string]interface{}) (*entities.Account, error)
	UpdateWithMap(ctx context.Context, record *entities.Account, params map[string]interface{}) error
	Delete(ctx context.Context, record *entities.Account) error
}
