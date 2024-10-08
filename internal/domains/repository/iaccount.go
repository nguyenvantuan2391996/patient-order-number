package repository

//go:generate mockgen -package=repository -destination=iaccount_mock.go -source=iaccount.go

type IAccountRepositoryInterface interface {
}
