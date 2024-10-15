package models

type AccountInput struct {
	UserName string
	Password string
}

type AccountUpdateInput struct {
	UserName string
	Password string
	UserID   int64
	Status   int
}

type DeleteAccountInput struct {
	UserID int64
}
