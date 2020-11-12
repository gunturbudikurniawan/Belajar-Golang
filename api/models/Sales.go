package models

import (
	"encoding/json"
	"time"
)

type sales struct {
	ID             uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Createdtm      time.Time       `json:"create_dtm"`
	Salesid        string          `json:"sales_id"`
	Userid         string          `json:"user_id"`
	Outletid       string          `json:"outlet_id"`
	Salestype      string          `json:"sales_type"`
	Customerid     string          `json:"customer_id"`
	Products       json.RawMessage `json:"products"`
	Subtotal       int             `json:"subtotal"`
	Totaldiskon    int             `json:"total_diskon"`
	Totaltax       json.RawMessage `json:"total_tax"`
	Totalbill      int             `json:"total_bill"`
	Paymentmethod  string          `json:"payment_method"`
	Paymentduedate string          `json:"payment_due_date"`
	Totalpayment   int             `json:"total_payment"`
	Exchange       int             `json:"exchange"`
	Notes          string          `json:"notes"`
	Totalbuycost   int             `json:"total_buy_cost"`
	Paymentdate    string          `json:"payment_date"`
	Rewardid       string          `json:"Reward_id"`
	Pointsredeem   int             `json:"points_redeem"`
}
