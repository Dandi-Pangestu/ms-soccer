package models

import "ms-soccer/service/shared/domains"

type Team struct {
	domains.Entity
	Name string `gorm:"type:varchar(255)" json:"user_id"`
}

func (m *Team) TableName() string {
	return "teams"
}
