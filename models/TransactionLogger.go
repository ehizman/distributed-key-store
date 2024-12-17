package models

type TransactionLogger interface {
	WriteDelete (key string)
	WritePut(key, value string)
}