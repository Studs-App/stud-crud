package repositories

import (
	"log"

	"github.com/stud-crud/base"
	"github.com/stud-crud/models"
)

func InsertStudySession(studySession models.StudySession) models.StudySession {
	log.Println("InsertStudySession()")
	db := base.DB
	db.Create(&studySession)
	return studySession
}
func ReadStudySessionsByProfileId(profileId string) []models.StudySession {
	log.Println("ReadStudySessionsByProfileId()")
	db := base.DB
	var allProfileStudySession []models.StudySession
	db.Where("profile_id = ?", profileId).Find(&allProfileStudySession)
	return allProfileStudySession
}

func ReadAllStudySessions() []models.StudySession {
	log.Println("ReadAllStudySessions()")
	db := base.DB
	var allStudySessions []models.StudySession
	db.Find(&allStudySessions)
	return allStudySessions
}

func ReadStudyBudsByProfileId(profileId string) []models.Buds {
	log.Println("ReadStudyBudsByProfileId")
	db := base.DB
	var budsSlice []models.Buds
	db.Raw("SELECT buds FROM study_sessions WHERE profile_id = ?", profileId).Scan(&budsSlice)
	return budsSlice
}
