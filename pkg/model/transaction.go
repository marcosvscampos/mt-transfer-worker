package model

type Transaction struct {
	ID                       string
	Amount                   float64
	OriginUserID             string
	OriginAccountNumber      string
	DestinationUserID        string
	DestinationAccountNumber string
	Status                   string
}
