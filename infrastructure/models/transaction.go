package models

type Transaction struct {
	Id              int32
	TransactionType TransactionType
	Amount          float32
	CategoryId      int32
}
