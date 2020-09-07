package mocked

import (
	"errors"
	"goproposal/internal/models"
	"time"
)

var (
	NotFoundErr = errors.New("Record not found")
)

func (m Mocked) GetPerson(id int64) (*models.Person, error) {

	person := &models.Person{}
	for _, p := range m.data.persons {
		if p.Id != id {
			continue
		}

		person.Id = p.Id
		person.Name = p.Name
		person.Age = p.Age
		person.DateCreated = p.DateCreated
	}

	if person.Id <= 0 {
		return nil, NotFoundErr
	}

	return person, nil
}

func (m Mocked) GetAllPersons() ([]models.Person, error) {
	if m.data.persons == nil {
		m.data.persons = []models.Person{}
	}

	return m.data.persons, nil
}

func (m Mocked) SavePerson(name string, age int32) error {
	id := len(m.data.persons) + 1
	p := models.Person{
		Id:          int64(id),
		Name:        name,
		Age:         age,
		DateCreated: time.Now(),
	}

	m.data.persons = append(m.data.persons, p)
	return nil
}
