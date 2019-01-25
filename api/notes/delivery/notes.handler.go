package delivery

import (
	"examplego/api/notes"
	"examplego/common"
	"examplego/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type Note = models.Note
type JSON = common.JSON

type NotesHandler struct {
	u notes.Usecase
}

func NewNotesHandler(u notes.Usecase) NotesHandler {
	return NotesHandler{
		u:  u,
	}
}

func (handler NotesHandler)create(c *gin.Context) {
	type RequestBody struct {
		NoteName string `json:"note_name" binding:"required"`
		NoteDetails string `json:"note_details" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}
	note := Note{NoteName: requestBody.NoteName, NoteDetails: requestBody.NoteDetails}
	handler.u.Create(&note)
	c.JSON(200, note.Serialize())
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var note Note
	if err := db.Where("id = ?", id).First(&note).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, note.Serialize())
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var notes []Note

	if err := db.Find(&notes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		length := len(notes)
		serialized := make([]JSON, length, length)
		for i := 0; i < length; i++ {
			serialized[i] = notes[i].Serialize()
		}

		c.JSON(200, serialized)
	}
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var note Note
	if err := db.Where("id = ?", id).First(&note).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	db.Delete(&note)
	c.Status(204)
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var note Note
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&note).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	note.UpdatedAt=time.Now()
	c.BindJSON(&note)

	db.Save(&note)
	c.JSON(200, note.Serialize())

}
