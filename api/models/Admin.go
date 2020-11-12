package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gunturbudikurniawan/Belajar-Golang/api/security"
	"github.com/jinzhu/gorm"
)

//Admin
type Admin struct {
	ID              uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Phone           string          `gorm:"size:100;" json:"phone"`
	Username        string          `gorm:"size:255;not null;unique" json:"username"`
	Create_dtm      time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"create_dtm`
	Email           string          `gorm:"size:100;not null;unique" json:"email"`
	Secret_password string          `json:"secret_password"`
	Idcard_image    json.RawMessage `json:"idcard_image"`
	UpdatedAt       time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Admin) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Secret_password)
	if err != nil {
		return err
	}
	u.Secret_password = string(hashedPassword)
	return nil
}

func (u *Admin) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Create_dtm = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *Admin) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}

	case "login":
		if u.Secret_password == "" {
			err = errors.New("Required Password")
			errorMessages["Required_password"] = err.Error()
		}
		if u.Secret_password == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	case "forgotpassword":
		if u.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	default:
		if u.Username == "" {
			err = errors.New("Required Username")
			errorMessages["Required_username"] = err.Error()
		}
		if u.Secret_password == "" {
			err = errors.New("Required Password")
			errorMessages["Required_password"] = err.Error()
		}
		if u.Secret_password != "" && len(u.Secret_password) < 6 {
			err = errors.New("Password should be atleast 6 characters")
			errorMessages["Invalid_password"] = err.Error()
		}
		if u.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()

		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	}
	return errorMessages
}

func (u *Admin) SaveUser(db *gorm.DB) (*Admin, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Admin{}, err
	}
	return u, nil
}

func (u *Admin) FindAllUsers(db *gorm.DB) (*[]Admin, error) {
	var err error
	admins := []Admin{}
	err = db.Debug().Model(&Admin{}).Limit(100).Find(&admins).Error
	if err != nil {
		return &[]Admin{}, err
	}
	return &admins, err
}

func (u *Admin) FindUserByID(db *gorm.DB, uid uint32) (*Admin, error) {
	var err error
	err = db.Debug().Model(Admin{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Admin{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Admin{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *Admin) UpdateAUser(db *gorm.DB, uid uint32) (*Admin, error) {

	if u.Secret_password != "" {
		// To hash the password
		err := u.BeforeSave()
		if err != nil {
			log.Fatal(err)
		}

		db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).UpdateColumns(
			map[string]interface{}{
				"secret_password": u.Secret_password,
				"email":           u.Email,
				"update_at":       time.Now(),
			},
		)
	}

	db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).UpdateColumns(
		map[string]interface{}{
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Admin{}, db.Error
	}

	// This is the display the updated user
	err := db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Admin{}, err
	}
	return u, nil
}

func (u *Admin) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).Delete(&Admin{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *Admin) UpdatePassword(db *gorm.DB) error {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&Admin{}).Where("email = ?", u.Email).Take(&Admin{}).UpdateColumns(
		map[string]interface{}{
			"secret_password": u.Secret_password,
			"update_at":       time.Now(),
		},
	)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

type Data struct {
	UserID    string
	OwnerName string
	Email     string
	LastTrx   *time.Time
}

func Show(db *gorm.DB) (error, []Data) {
	var datas []Data

	query := `SELECT user_id, owner_name, email, Z.create_dtm as last_trx FROM (
		SELECT user_id,owner_name, email, (SELECT create_dtm FROM sales WHERE create_dtm > current_date-7 AND user_id = b.user_id ORDER BY id DESC LIMIT 1) FROM subscribers b
		UNION SELECT user_id, owner_name, email, (SELECT create_dtm FROM onlinesales WHERE create_dtm > current_date-7 AND user_id = b.user_id ORDER BY id DESC LIMIT 1) FROM subscribers b
		UNION SELECT user_id, owner_name, email, (SELECT create_dtm FROM saved_orders so WHERE create_dtm > current_date-7 AND user_id = b.user_id ORDER BY id DESC LIMIT 1) FROM subscribers b) AS Z`

	/** saya lebih suka seperti ini */
	err := db.Raw(query).Scan(&datas).Error
	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	// for rows.Next() {
	/**
	var (
		user_id    sql.NullString
		owner_name sql.NullString
		email      sql.NullString
		last_trx   sql.NullTime
	)

	err = rows.Scan(&user_id, &owner_name, &email, &last_trx)
	if err != nil {
		// handle this error
		fmt.Errorf("%v", err)
		return err, datas
	}

	datas = append(datas, Data{
		UserID:    user_id.String,
		OwnerName: owner_name.String,
		Email:     email.String,
		LastTrx:   last_trx.Time,
	})
	*/
	// }

	return nil, datas
}
