package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID               int32          `json:"id" gorm:"primaryKey"`
	FromChat         Chat           `json:"from_chat" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	FromUser         User           `json:"from_user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ForwardedFrom    User           `json:"forwarder_from,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ReplyToMessage   *Message       `json:"reply_to,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ReplyToMessageID *uint          `json:"-" gorm:"index"`
	Text             string         `json:"text" gorm:"type:varchar(250);not null"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at,omitempty"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
