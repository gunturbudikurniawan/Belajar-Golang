package models

import (
	"encoding/json"
	"time"
)

type Sales struct {
	ID             uint32          `gorm:"primary_key;auto_increment" json:"id"`
	CreateDTM      time.Time       `json:"create_dtm"`
	SalesID        string          `json:"sales_id"`
	UserID         string          `json:"user_id"`
	OutletID       string          `json:"outlet_id"`
	SalesType      string          `json:"sales_type"`
	CustomerID     string          `json:"customer_id"`
	Products       json.RawMessage `json:"products"`
	Subtotal       int             `json:"subtotal"`
	TotalDiskon    int             `json:"total_diskon"`
	TotalTax       json.RawMessage `json:"total_tax"`
	TotalBill      int             `json:"total_bill"`
	PaymentMethod  string          `json:"payment_method"`
	PaymentDueDate string          `json:"payment_due_date"`
	TotalPayment   int             `json:"total_payment"`
	Exchange       int             `json:"exchange"`
	Notes          string          `json:"notes"`
	TotalBuyCost   int             `json:"total_buy_cost"`
	PaymentDate    string          `json:"payment_date"`
	RewardID       string          `json:"Reward_id"`
	PointsRedeem   int             `json:"points_redeem"`
}
