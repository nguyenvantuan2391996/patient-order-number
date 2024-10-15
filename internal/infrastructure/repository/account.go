package repository

import (
	"context"

	"github.com/nguyenvantuan2391996/patient-order-number/base_common/database/entities"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (ar *AccountRepository) Create(ctx context.Context, record *entities.Account) error {
	return ar.db.WithContext(ctx).Create(&record).Error
}

func (ar *AccountRepository) GetByQueries(ctx context.Context, queries map[string]interface{}) (*entities.Account, error) {
	var record *entities.Account
	err := ar.db.WithContext(ctx).Where(queries).First(&record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (ar *AccountRepository) UpdateWithMap(ctx context.Context, record *entities.Account, params map[string]interface{}) error {
	return ar.db.WithContext(ctx).Model(record).Updates(params).Error
}

func (ar *AccountRepository) Delete(ctx context.Context, record *entities.Account) error {
	return ar.db.WithContext(ctx).Delete(record).Error
}
