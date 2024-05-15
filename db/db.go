package db

type DbInterface interface {
	Get(string) ([]byte, error)
	Create(any) ([]byte, error)
	Table(string) DbInterface
}

var Client DbInterface = &Mongo
