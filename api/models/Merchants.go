package models

import (
	"encoding/json"
	"time"
)

type subscribers struct {
	ID             uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Createdtm      time.Time       `json:"create_dtm"`
	Userid         string          `json:"user_id"`
	Email          string          `json:"email"`
	Ownername      string          `json:"owner_name"`
	Secretpassword string          `json:"secret_password"`
	Fcmtoken       string          `json:"fcm_token"`
	Idcardname     string          `json:"idcard_name"`
	Idcardnumber   string          `json:"idcard_number"`
	Bankholdername string          `json:"bank_holder_name"`
	Bankname       string          `json:"bank_name"`
	Bankaccount    string          `json:"bank_account"`
	Idcardimage    json.RawMessage `json:"idcard_image"`
	Referralcode   string          `json:"referral_code"`
}
