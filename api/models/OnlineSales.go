package models

import (
	"encoding/json"
	"time"
)

type onlinesales struct {
	ID             uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Createdtm      time.Time       `json:"create_dtm"`
	Salesid        string          `json:"sales_id"`
	Userid         string          `json:"user_id"`
	Outletid       string          `json:"outlet_id"`
	Customerid     string          `json:"customer_id"`
	Customer       json.RawMessage `json:"customer"`
	Products       json.RawMessage `json:"products"`
	Subtotal       int             `json:"subtotal"`
	Totaldiskon    int             `json:"total_diskon"`
	Totaltax       json.RawMessage `json:"total_tax"`
	Totalbill      int             `json:"total_bill"`
	Paymentmethod  string          `json:"payment_method"`
	Paymentaccount string          `json:"payment_account"`
	Paymentduedate string          `json:"payment_due_date"`
	Totalpayment   int             `json:"total_payment"`
	Expedition     string          `json:"expedition"`
	Service        string          `json:"service"`
	Weight         int             `json:"weight"`
	Deliverycost   int             `json:"delivery_cost"`
	Notes          string          `json:"notes"`
	Totalbuycost   int             `json:"total_buy_cost"`
	Paymentdate    string          `json:"payment_date"`
	Rewardid       string          `json:"reward_id"`
	Pointsredeem   int             `json:"points_redeem"`
	Orderstatus    string          `json:"order_status"`
	Shipmentnumber string          `json:"shipment_number"`
}
