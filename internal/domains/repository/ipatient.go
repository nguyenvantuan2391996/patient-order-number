package repository

import (
	"context"

	"github.com/nguyenvantuan2391996/patient-order-number/base_common/database/entities"
)

//go:generate mockgen -package=repository -destination=ipatient_mock.go -source=ipatient.go

type IPatientRepositoryInterface interface {
	Create(ctx context.Context, record *entities.Patient) error
	GetByQueries(ctx context.Context, queries map[string]interface{}) (*entities.Patient, error)
	List(ctx context.Context, queries map[string]interface{}, page, limit int,
		conditions ...string) ([]*entities.Patient, error)
	Total(ctx context.Context, queries map[string]interface{}, conditions ...string) (int64, error)
	UpdateWithMap(ctx context.Context, record *entities.Patient, params map[string]interface{}) error
	Delete(ctx context.Context, record *entities.Patient) error
}
