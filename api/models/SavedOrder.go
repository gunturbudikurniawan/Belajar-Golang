package models

import (
	"encoding/json"
	"time"
)

type savedorders struct {
	ID            uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Createdtm     time.Time       `json:"create_dtm"`
	Userid        string          `json:"user_id"`
	Outletid      string          `json:"outlet_id"`
	Savedordersid string          `json:"saved_orders_id"`
	Name          string          `json:"name"`
	Phone         string          `json:"phone"`
	Orders        json.RawMessage `json:"orders"`
	Tableid       string          `json:"table_id"`
}
