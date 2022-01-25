package user

import (
	"goer/app/models"
)

type User struct {
	models.BaseModel

	Uid          string `json:"uid"`
	CountryId    int64  `json:"country_id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
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
