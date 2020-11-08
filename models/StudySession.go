package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type StudySession struct {
	gorm.Model
	Title        string         `json:"title"`
	Subject      string         `json:"subject"`
	Status       string         `json:"status"`
	Lattitude    float64        `json:"lattitude"`
	Longitude    float64        `json:"longitude"`
	Location     string         `json:"location"`
	Buds         pq.StringArray `gorm:"type:varchar(64)[]" json:"buds"`
	Duration     string         `json:"duration"`
	ScheduledDate string         `json:"scheduled_date "`
	IsPrivate    bool           `json:"is_private"`
	ProfileId    int            `json:"profile_id"`
}
