package notes

import "examplego/models"

type Usecase interface {
	Create(note *models.Note) (resp *models.Note)
	GetById(id int) (resp *models.Note, err error)
	List() (resp *[]models.Note, err error)
	Delete(id int)
	Update(id int) (resp *models.Note, err error)
}

