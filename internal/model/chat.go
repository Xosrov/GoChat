package model

import (
	"time"

	"gorm.io/gorm"
)

type ChatType uint8

const (
	ChatTypePrivate ChatType = iota
	ChatTypeGroup
)

type Chat struct {
	ID             int32          `json:"id" gorm:"primaryKey"`
	Uuid           string         `json:"uuid" gorm:"unique;not null;type:varchar(100);comment:'unique manditory group name'"`
	Type           ChatType       `json:"type" gorm:"not null;comment:'message type; private, group, etc'"`
	Title          string         `json:"title,omitempty" gorm:"type:varchar(50)"`
	Bio            string         `json:"bio,omitempty" gorm:"type:varchar(50)"`
	InvitationLink string         `json:"invitation_link,omitempty" gorm:"type:varchar(50)"`
	UserList       User           `json:"user_list" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at,omitempty"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
