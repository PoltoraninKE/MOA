package database

type Repository interface {
	Create(model interface{}) error
	Get(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
}
