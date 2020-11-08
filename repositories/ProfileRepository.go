package repositories

import (
	"fmt"
	"log"

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
	fmt.Printf("profiles: %v", profile)
	db.Create(&profile)
	return profile
}
