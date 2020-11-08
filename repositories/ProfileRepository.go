package repositories

import (
	"fmt"
	"log"

	"github.com/stud-crud/base"
	"github.com/stud-crud/models"
)

func GetAllProfiles() [] models.Profile{
	log.Println("GetAllProfiles()")
	db := base.DB
	var profiles []models.Profile
	db.Find(&profiles)
	return profiles
}

func GetProfileById(id int) models.Profile {
	log.Println("GetProfileById()")
	db := base.DB
	var profile models.Profile
	db.Find(&profile)
	return profile
}

func InsertProfile(profile models.Profile) models.Profile {
	log.Println("InsertProfile()")
	db := base.DB
	fmt.Printf("profiles: %v", profile)
	db.Create(&profile)
	return profile
}
