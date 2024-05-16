package db

type DbInterface interface {
	Connect(string) error
	Get(string) ([]byte, error)
	Create(interface{}) ([]byte, error)
	Patch(string, interface{}) ([]byte, error)
	Table(string) DbInterface
	Unmarshal([]byte, interface{}) error
}

var Client DbInterface = &Mongo
