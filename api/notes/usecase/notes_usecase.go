package usecase

import (
	"examplego/api/notes"
	"examplego/models"
)

type notesUsecase struct {
	repo notes.Repository
}

func NewNotesUsecase(n notes.Repository) notes.Usecase {
	return &notesUsecase{
		repo:  n,
	}
}

func (n *notesUsecase) Create(note *models.Note) (resp *models.Note) {
	return n.repo.Save(note)
}

func (notesUsecase) GetById(id int) (resp *models.Note, err error) {
	panic("implement me")
}

func (notesUsecase) List() (resp *[]models.Note, err error) {
	panic("implement me")
}

func (notesUsecase) Delete(id int) {
	panic("implement me")
}

func (notesUsecase) Update(id int) (resp *models.Note, err error) {
	panic("implement me")
}



