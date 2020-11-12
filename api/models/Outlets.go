package models

import (
	"encoding/json"
	"time"
)

type outlets struct {
	ID                  uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Createdtm           time.Time       `json:"create_dtm"`
	Userid              string          `json:"user_id"`
	Outletid            string          `json:"outlet_id"`
	Nama                string          `json:"nama"`
	Address             string          `json:"address"`
	Phone               string          `json:"phone"`
	Businesscategory    string          `json:"business_category"`
	Isactive            string          `json:"is_active"`
	Accounts            json.RawMessage `json:"accounts"`
	Images              json.RawMessage `json:"images"`
	Miniwebsiteurl      string          `json:"mini_website_url"`
	Isonlinestoreactive string          `json:"is_online_store_active"`
}
