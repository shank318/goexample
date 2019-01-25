package repository

import (
	"examplego/api/notes"
	"examplego/models"
	"github.com/jinzhu/gorm"
)

type mysqlNotesRepository struct {
	db *gorm.DB
}

type Note = models.Note 

func NewMysqlNotesRepository(db *gorm.DB) notes.Repository {
	return &mysqlNotesRepository{db}
}

func (m *mysqlNotesRepository) Save(note *Note) (resp *Note) {
	m.db.NewRecord(note)
	m.db.Create(&note)
	return note
}

func (m *mysqlNotesRepository) Read(id int) (resp *Note, err error) {
	panic("implement me")
}

func (mysqlNotesRepository) List() (resp *[]Note, err error) {
	panic("implement me")
}

func (mysqlNotesRepository) Remove(id int) {
	panic("implement me")
}

func (mysqlNotesRepository) Update(id int) (resp *Note, err error) {
	panic("implement me")
}

