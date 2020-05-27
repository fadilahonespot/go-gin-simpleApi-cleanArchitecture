package person

import "simpleApic/model"

type PersonRepo interface {
	Create(person *model.Person) (*model.Person, error)
	ReadAll() (*[]model.Person, error)
	ReadById(id int)(*model.Person, error)
	Update(id int, person *model.Person) (*model.Person, error)
	Delete(id int) error
}