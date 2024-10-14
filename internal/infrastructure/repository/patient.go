package repository

import (
	"context"

	"github.com/nguyenvantuan2391996/patient-order-number/base_common/database/entities"
	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (pr *PatientRepository) Create(ctx context.Context, record *entities.Patient) error {
	return pr.db.WithContext(ctx).Create(&record).Error
}

func (pr *PatientRepository) GetByQueries(ctx context.Context,
	queries map[string]interface{}) (*entities.Patient, error) {
	var record *entities.Patient
	err := pr.db.WithContext(ctx).Where(queries).First(&record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (pr *PatientRepository) List(ctx context.Context, queries map[string]interface{}, page, limit int,
	conditions ...string) ([]*entities.Patient, error) {
	records := make([]*entities.Patient, 0)

	query := pr.db.WithContext(ctx).Model(&entities.Patient{})
	if len(queries) > 0 {
		query = query.Where(queries)
	}

	for _, cond := range conditions {
		query = query.Where(cond)
	}

	query.Offset((page - 1) * limit).Limit(limit)
	err := query.Find(&records).Error
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (pr *PatientRepository) Total(ctx context.Context, queries map[string]interface{}, conditions ...string) (int64, error) {
	total := int64(0)

	query := pr.db.WithContext(ctx).Model(&entities.Patient{}).Where(queries)
	if len(queries) > 0 {
		query = query.Where(queries)
	}

	for _, cond := range conditions {
		query = query.Where(cond)
	}

	err := query.Count(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (pr *PatientRepository) UpdateWithMap(ctx context.Context, record *entities.Patient,
	params map[string]interface{}) error {
	return pr.db.WithContext(ctx).Model(record).Updates(params).Error
}

func (pr *PatientRepository) Delete(ctx context.Context, record *entities.Patient) error {
	return pr.db.WithContext(ctx).Delete(record).Error
}
