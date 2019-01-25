package models

import (
	"examplego/common"
	"github.com/jinzhu/gorm"
)

// Note data model
type Note struct {
	gorm.Model
	NoteName string
	NoteDetails string
}

// Serialize serializes note data
func (n Note) Serialize() common.JSON {
	return common.JSON{
		"id":         n.ID,
		"note_name":       n.NoteName,
		"note_details":       n.NoteDetails,
		"created_at": n.CreatedAt,
		"updated_at": n.UpdatedAt,

	}
}

