package models

import (
	"encoding/json"
	"time"
)

type admin struct {
	ID             uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Phone          string          `gorm:"size:100;" json:"phone"`
	Username       string          `gorm:"size:255;not null;unique" json:"username"`
	Createdtm      time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Email          string          `gorm:"size:100;not null;unique" json:"email"`
	Secretpassword string          `json:"secret_password"`
	Idcardimage    json.RawMessage `json:"idcard_image"`
}
