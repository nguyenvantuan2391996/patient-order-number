package patient

import "github.com/nguyenvantuan2391996/patient-order-number/internal/domains/repository"

type Patient struct {
	accountRepo repository.IAccountRepositoryInterface
}

func NewPatientService(accountRepo repository.IAccountRepositoryInterface) *Patient {
	return &Patient{
		accountRepo: accountRepo,
	}
}
