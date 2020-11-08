package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/stud-crud/models"
	"github.com/stud-crud/repositories"
)

func CreateStudySession(c *fiber.Ctx) error {
	var studySession models.StudySession
	params := new(struct {
		Title         string   `json:"title"`
		Description   string   `json:"description"`
		Subject       string   `json:"subject"`
		Status        string   `json:"status"`
		Lattitude     float64  `json:"lattitude"`
		Longitude     float64  `json:"longitude"`
		Location      string   `json:"location"`
		Buds          []string `json:"buds"`
		Duration      string   `json:"duration"`
		ScheduledDate string   `json:"scheduledDate"`
		IsPrivate     bool     `json:"is_private"`
		ProfileId     int      `json:"profile_id"`
	})
	c.BodyParser(&params)
	studySession.Title = params.Title
	studySession.Description = params.Description
	studySession.Subject = params.Subject
	studySession.Status = params.Status
	studySession.Lattitude = params.Lattitude
	studySession.Longitude = params.Longitude
	studySession.Location = params.Location
	studySession.Buds = params.Buds
	studySession.Duration = params.Duration
	studySession.ScheduledDate = params.ScheduledDate
	studySession.IsPrivate = params.IsPrivate
	studySession.ProfileId = params.ProfileId

	studySession = repositories.InsertStudySession(studySession)
	return c.JSON(studySession)
}

func GetAllStudySessionsByProfileID(c *fiber.Ctx) error {
	log.Println("GetAllStudySessionsByProfileID()")
	profileId := c.Params("id")
	studySessions := repositories.ReadStudySessionsByProfileId(profileId)
	return c.JSON(studySessions)
}

func GetAllStudySessions(c *fiber.Ctx) error {
	log.Println("GetAllStudySessions()")
	allStudySessions := repositories.ReadAllStudySessions()
	return c.JSON(allStudySessions)
}
