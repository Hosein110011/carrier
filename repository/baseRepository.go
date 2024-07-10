package repository

type baseRepository interface {
	Create(string, interface{}) error
	GetList() ([]interface{}, error)
	GetById(string) (interface{}, error)
	Update(string, interface{}) (interface{}, error)
	Delete(string) error
}
