package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	"github.com/stud-crud/base"
	"github.com/stud-crud/models"
	"github.com/stud-crud/services"
)

func initRoutes(app *fiber.App) {
	app.Get("/get/profile", services.GetProfileById)
	app.Get("/get/profile/all", services.GetAllProfiles)
	app.Post("/create/profile", services.CreateProfile)
	app.Post("/create/studySession", services.CreateStudySession)
	app.Get("/get/studySession/:id/all", services.GetAllStudySessionsByProfileID)
	app.Get("/get/studySession/all", services.GetAllStudySessions)
}

func initializeDB() {
	var err error
	base.DB, err = gorm.Open("postgres", base.DBUrl(base.BuildDBConfig()))
	if err != nil {
		panic(err)
	}
	log.Println("DB connection successful!")

	base.DB.AutoMigrate(&models.Profile{})
	base.DB.AutoMigrate(&models.StudySession{})
	log.Println("Database Migrated")

}

func main() {
	server := fiber.New()
	server.Use(cors.New())
	initializeDB()
	defer base.DB.Close()
	initRoutes(server)
	err := server.Listen(":3030")
	if err != nil {
		log.Println(err)
	}
}
