package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Profile struct {
	gorm.Model
	FirstName       string         `json:"firstName"`
	LastName        string         `json:"lastName"`
	PictureUrl      string         `json:"pictureUrl"`
	AuthenticatorId string         `json:"authenticator_id"`
	Major           string         `json:"major"`
	School          string         `json:"school"`
	Studs           pq.StringArray `gorm:"type:varchar(64)[]" json:"studs"`
	Buds            pq.StringArray `gorm:"type:varchar(64)[]" json:"buds"`
}
