package models

type TransactionType int8

const (
	Undefined TransactionType = iota
	Add       TransactionType = 1
	Substract TransactionType = 2
)
