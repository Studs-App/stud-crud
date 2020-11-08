package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Profile struct {
	gorm.Model
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	AuthenticatorId string         `json:"authenticator_id"`
	Major           string         `json:"major"`
	School          string         `json:"school"`
	Studs           pq.StringArray `gorm:"type:varchar(64)[]" json:"stud"`
}
