package main

import (
	"examplego/api/notes/delivery"
	"examplego/api/notes/repository"
	"examplego/api/notes/usecase"
	"examplego/api/pingdom"
	"examplego/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initializes database
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app
	app.Use(database.Inject(db))

	// Notes api's
	notesRepo := repository.NewMysqlNotesRepository(db)
	notesUseCase := usecase.NewNotesUsecase(notesRepo)
	notesHandler :=delivery.NewNotesHandler(notesUseCase)
	delivery.ApplyRoutes(app,notesHandler)

	// Pingdom api's
	pingdom.ApplyRoutes(app)
	app.Run(":" + port)  // listen to given port
}
