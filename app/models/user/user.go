package user

import (
	"goer/app/models"
)

type User struct {
	models.BaseModel

	Uid          string `json:"uid"`
	CountryId    int64  `json:"country_id"`
	Name         string `json:"name"`
	Username     string `json:"username" gorm:"default:null"`
	Email        string `json:"email" gorm:"default:null"`
	Phone        string `json:"phone" gorm:"default:null"`
	Password     string `json:"-"`
	PayPassword  string `json:"-"`
	GoogleKey    string `json:"-"`
	GoogleStatus string `json:"google_status"`
	Sso          string `json:"-"`
	Pid          uint64 `json:"pid"`
	InviteCount  uint64 `json:"invite_count"`
	Depth        int64  `json:"depth"`
	IsValid      bool   `json:"is_valid"`
	KycStatus    int64  `json:"kyc_status"`
	Age          int64  `json:"age"`
	Gender       string `json:"gender"`
	Avatar       string `json:"avatar"`

	models.TimestampsField
}

// Google status: 'unbind','disabled','enabled'
type GoogleStatus string

// Gender: 'secret', 'male', 'female'
type Gender string

const (
	GoogleStatusUnbind   GoogleStatus = "unbind"
	GoogleStatusDisabled GoogleStatus = "disabled"
	GoogleStatusEnabled  GoogleStatus = "enabled"

	GenderSecret Gender = "secret"
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)
