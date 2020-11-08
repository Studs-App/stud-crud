package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type StudySession struct {
	gorm.Model
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Subject       string         `json:"subject"`
	Status        string         `json:"status"`
	Lattitude     float64        `json:"lattitude"`
	Longitude     float64        `json:"longitude"`
	Location      string         `json:"location"`
	Buds          pq.StringArray `gorm:"type:varchar(64)[]" json:"buds"`
	Duration      string         `json:"duration"`
	ScheduledDate string         `json:"scheduledDate "`
	IsPrivate     bool           `json:"isPrivate"`
	ProfileId     int            `json:"profile_id"`
}
