package repo

import (
	"fmt"
	"simpleApic/model"
	"simpleApic/person"

	"github.com/jinzhu/gorm"
)

type PersonRepoImpl struct {
	DB *gorm.DB
}

func CreatePersonRepo(DB *gorm.DB) person.PersonRepo {
	return &PersonRepoImpl{DB}
}

func (e *PersonRepoImpl) Create(person *model.Person) (*model.Person, error) {
	err := e.DB.Save(&person).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return person, nil
}

func (e *PersonRepoImpl) ReadAll() (*[]model.Person, error) {
	var persons []model.Person
	err := e.DB.Find(&persons).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.ReadAll] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data")
	}
	return &persons, nil
}

func (e *PersonRepoImpl) ReadById(id int)(*model.Person, error) {
	var person = model.Person{}
	err := e.DB.Table("person").Where("id = ?", id).First(&person).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.ReadById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}
	return &person, nil
}

func (e *PersonRepoImpl) Update(id int, person *model.Person) (*model.Person, error) {
	var upPerson = model.Person{}
	err := e.DB.Table("person").Where("id = ?", id).First(&upPerson).Update(&person).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.Update] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data")
	}
	return &upPerson, nil
}

func (e *PersonRepoImpl) Delete(id int) error {
	var person = model.Person{}
	err := e.DB.Table("person").Where("id = ?", id).First(&person).Delete(&person).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.Delete] error execute query %v \n", err)
		return fmt.Errorf("id is not exsis")
	}
	return nil
}