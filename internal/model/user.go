package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int32          `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"unique;not null;type:varchar(150)"`
	Uuid      string         `json:"uuid,omitempty" gorm:"unique;type:varchar(100);comment:'unique nickname'"`
	Bio       string         `json:"bio,omitempty" gorm:"type:varchar(50)"`
	FirstName string         `json:"first_name" gorm:"not null;type:varchar(150)"`
	LastName  string         `json:"last_name,omitempty" gorm:"type:varchar(150)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
