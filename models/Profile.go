package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Studs []string

func (data Studs) Value() (driver.Value, error) {

	return json.Marshal(data)
}
func (data *Studs) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("[]byte assertion failed")
	}

	return json.Unmarshal(b, data)
}

type Profile struct {
	gorm.Model
	FirstName       string         `json:"first_name"`
	LastName        string           `json:"last_name"`
	AuthenticatorId string         `json:"authenticator_id"`
	Major           string         `json:"major"`
	School          string            `json:"school"`
	Studs           pq.StringArray `gorm:"type:varchar(64)[]" json:"stud"`
}
