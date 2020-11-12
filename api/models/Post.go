package models

import (
	"time"
)

type post struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Phone     string    `gorm:"size:255;not null;" json:"phone"`
	Content   string    `gorm:"text;not null;" json:"content"`
	Author    admin     `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
