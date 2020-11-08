package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stud-crud/models"
	"github.com/stud-crud/repositories"
)

func GetProfileById(c *fiber.Ctx) error {
	profile := repositories.ReadProfileById(1)
	return c.JSON(profile)
}

func GetAllProfiles(c *fiber.Ctx) error {
	profiles := repositories.ReadAllProfiles()
	return c.JSON(profiles)
}

func CreateProfile(c *fiber.Ctx) error {

	params := new(struct {
		FirstName       string   `json:"first_name"`
		LastName        string   `json:"last_name"`
		AuthenticatorId string   `json:"authenticator_id"`
		Major           string   `json:"major"`
		School          string   `json:"school"`
		Studs           []string `json:"studs"`
	})
	c.BodyParser(&params)
	var profile models.Profile
	profile.FirstName = params.FirstName
	profile.LastName = params.LastName
	profile.AuthenticatorId = params.AuthenticatorId
	profile.Major = params.Major
	profile.School = params.School
	profile.Studs = params.Studs
	repositories.InsertProfile(profile)
	return c.JSON(profile)
}
