package utils

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type ModelBase struct {
	ID        uuid.UUID  `gorm:"type:varchar(75);primary_key;"json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt";sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *ModelBase) BeforeCreate(scope *gorm.Scope) (err error) {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}
