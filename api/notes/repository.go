package notes

import "examplego/models"

type Repository interface {
	Save(note *models.Note) (resp *models.Note)
	Read(id int) (resp *models.Note, err error)
	List() (resp *[]models.Note, err error)
	Remove(id int)
	Update(id int) (resp *models.Note, err error)
}
