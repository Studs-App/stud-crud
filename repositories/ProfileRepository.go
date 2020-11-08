package repositories

import (
	"log"

	"github.com/lib/pq"
	"github.com/stud-crud/base"
	"github.com/stud-crud/models"
)

func ReadAllProfiles() []models.Profile {
	log.Println("ReadAllProfiles()")
	db := base.DB
	var profiles []models.Profile
	db.Find(&profiles)
	return profiles
}

func ReadProfileById(id string) models.Profile {
	log.Println("ReadProfileById()")
	db := base.DB
	var profile models.Profile
	db.Find(&profile, id)
	return profile
}

func InsertProfile(profile models.Profile) models.Profile {
	log.Println("InsertProfile()")
	db := base.DB
	db.Create(&profile)
	return profile
}

func UpdateProfileBuds(profileId string, buds pq.StringArray) {
	log.Println("UpdateProfileBuds()")
	db := base.DB
	var profile models.Profile
	db.Model(&profile).Where("id = ?", profileId).Update("buds", buds)
}
