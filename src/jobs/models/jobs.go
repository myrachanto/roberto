package models

import "gorm.io/gorm"

type Job struct {
	Name           string `json:"name,omitempty"`
	Source         string `json:"source,omitempty"`
	Destination    string `json:"destination,omitempty"`
	Database       string `json:"database,omitempty"`
	CollectionName string `json:"collection_name,omitempty"`
	gorm.Model
}
