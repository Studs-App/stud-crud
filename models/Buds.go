package models

import "github.com/lib/pq"

type Buds struct {
	Buds pq.StringArray `gorm:"type:varchar(64)[]" json:"budsFromSessions"`
}
