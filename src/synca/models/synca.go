package models

import (
	"time"

	"gorm.io/gorm"
)

type Synca struct {
	Name         string    `json:"name,omitempty"`
	DatabaseA    string    `json:"databaseA,omitempty"`
	DatabaseAUrl string    `json:"databaseaurl,omitempty"`
	DatabaseB    string    `json:"databaseB,omitempty"`
	DatabaseBUrl string    `json:"databaseurl,omitempty"`
	Dated        time.Time `json:"dated,omitempty"`
	Start        time.Time `json:"start,omitempty"`
	Ending       string    `json:"ending,omitempty"`
	Message      string    `json:"message,omitempty"`
	Status       string    `json:"status,omitempty"`
	Items        int       `json:"items,omitempty"`
	gorm.Model
}
