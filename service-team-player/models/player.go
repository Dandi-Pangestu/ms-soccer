package models

import "ms-soccer/service/shared/domains"

type Player struct {
	domains.Entity
	TeamID string `gorm:"type:varchar(100)"`
	Name   string `gorm:"type:varchar(255)" json:"name"`
	Number int
}

func (m *Player) TableName() string {
	return "players"
}
