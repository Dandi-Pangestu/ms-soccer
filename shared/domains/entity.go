package domains

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Entity struct {
	ID        string     `gorm:"primary_key; type:varchar(100)" json:"id" bson:"id"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time `sql:"index" bson:"deleted_at"`
	CreatedBy *string    `gorm:"column:created_by; type:varchar(100);" json:"created_by" bson:"created_ay"`
	UpdatedBy *string    `gorm:"column:updated_by; type:varchar(100);" json:"updated_ay" bson:"updated_ay"`
}

func (e *Entity) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("id", uuid.NewV4().String())
}

func (e *Entity) SetCreatedBy(userID *string) {
	e.CreatedBy = userID
}

func (e *Entity) SetUpdatedBy(userID *string) {
	e.UpdatedBy = userID
}

func (e *Entity) SetCreatedAt(t time.Time) {
	e.CreatedAt = t
}

func (e Entity) SetUpdatedAt(t time.Time) {
	e.UpdatedAt = t
}
