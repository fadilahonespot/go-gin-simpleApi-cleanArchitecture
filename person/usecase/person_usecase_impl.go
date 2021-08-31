package usecase

import (
	"simpleApic/person"
	"simpleApic/model"
)

type PersonUsecaseImpl struct {
	personRepo person.PersonRepo
}

func CreatePersonUsecase(personRepo person.PersonRepo) person.PersonUsecase {
	return &PersonUsecaseImpl{personRepo}
}

func (e *PersonUsecaseImpl) Create(person *model.Person) (*model.Person, error) {
	return e.personRepo.Create(person)
}

func (e *PersonUsecaseImpl) ReadAll() (*[]model.Person, error) {
	return e.personRepo.ReadAll()
}

func (e *PersonUsecaseImpl) ReadById(id int)(*model.Person, error) {
	return e.personRepo.ReadById(id)
}

func (e *PersonUsecaseImpl) Update(id int, person *model.Person) (*model.Person, error) {
	return e.personRepo.Update(id, person)
}

func (e *PersonUsecaseImpl) Delete(id int) error {
	return e.personRepo.Delete(id)
}