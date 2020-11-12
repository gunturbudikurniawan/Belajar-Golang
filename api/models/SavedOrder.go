package models

import (
	"encoding/json"
	"time"
)

type Saved_orders struct {
	ID              uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Create_dtm      time.Time       `json:"create_dtm"`
	User_id         string          `json:"user_id"`
	Outlet_id       string          `json:"outlet_id"`
	Saved_orders_id string          `json:"saved_orders_id"`
	Name            string          `json:"name"`
	Phone           string          `json:"phone"`
	Orders          json.RawMessage `json:"orders"`
	Table_id        string          `json:"table_id"`
}
