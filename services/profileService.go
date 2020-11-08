package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stud-crud/models"
	"github.com/stud-crud/repositories"
)

func GetProfileById(c *fiber.Ctx) error {
	profileId := c.Params("id")
	profile := repositories.ReadProfileById(profileId)
	buds := repositories.ReadStudyBudsByProfileId(profileId)
	profile.Buds = formatBudsForProfiles(buds)
	go repositories.UpdateProfileBuds(profileId, profile.Buds)
	return c.JSON(profile)
}

func formatBudsForProfiles(buds []models.Buds) []string {
	var budSlice []string
	for _, bud := range buds {
		budFromSession := bud.Buds
		for _, b := range budFromSession {
			budSlice = append(budSlice, b)
		}
	}
	return budSlice

}

func GetAllProfiles(c *fiber.Ctx) error {
	profiles := repositories.ReadAllProfiles()
	return c.JSON(profiles)
}

func CreateProfile(c *fiber.Ctx) error {

	params := new(struct {
		FirstName       string   `json:"first_name"`
		LastName        string   `json:"last_name"`
		PictureUrl      string   `json:"picture_url"`
		AuthenticatorId string   `json:"authenticator_id"`
		Major           string   `json:"major"`
		School          string   `json:"school"`
		Studs           []string `json:"studs"`
	})
	c.BodyParser(&params)
	var profile models.Profile
	profile.FirstName = params.FirstName
	profile.LastName = params.LastName
	profile.PictureUrl = params.PictureUrl
	profile.AuthenticatorId = params.AuthenticatorId
	profile.Major = params.Major
	profile.School = params.School
	profile.Studs = params.Studs
	repositories.InsertProfile(profile)
	return c.JSON(profile)
}
