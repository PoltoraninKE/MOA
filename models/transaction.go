package models

type Transaction struct {
	Id              int32
	TransactionType TransactionType
	IsValid         bool
}
