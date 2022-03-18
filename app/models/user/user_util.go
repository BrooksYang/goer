package user

import (
	"goer/global"

	"github.com/goer-project/goer-utils/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GenerateUid() string {
	length := global.Config.Common.UidLength

	return helpers.RandomNumber(length)
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func (user *User) CheckPayPassword(payPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PayPassword), []byte(payPassword))
	if err != nil {
		return false
	}

	return true
}

func (user *User) HasPayPassword() bool {
	return !helpers.Empty(user.PayPassword)
}

func (user *User) CheckGoogleCode(googleCode string) bool {
	if user.GoogleStatus != string(GoogleStatusEnabled) {
		return true
	}

	// todo::check google code
	return true
}

func (user *User) Check2FA(payPassword string, googleCode string) bool {
	// Check pay password
	if !user.CheckPayPassword(payPassword) {
		return false
	}

	// Check google code
	return user.CheckGoogleCode(googleCode)
}

func (user *User) AccountExists() bool {
	var id int
	if user.Email != "" {
		global.DB.Model(&user).Where("email=?", user.Email).Select("ID").First(&id)
		if id > 0 {
			return true
		}
	}

	if user.Phone != "" {
		global.DB.Model(&user).Where("phone=?", user.Phone).Select("ID").First(&id)
	}

	return id > 0
}

func SearchAccount(account string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if account == "" {
			return db
		}

		return db.Where(
			db.Where("id = ?", account).
				Or("uid = ?", account).
				Or("email = ?", account).
				Or("phone = ?", account).
				Or("username = ?", account),
		)
	}
}

func SearchIsValid(isValid string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if isValid == "" {
			return db
		}

		return db.Where("is_valid = ?", isValid)
	}
}

func (user *User) GetChildrenIdSubQuery(account string) *gorm.DB {
	subQuery := global.DB.Model(&User{}).
		Select("id").
		Where("pid", user.ID).
		Scopes(SearchAccount(account))

	return subQuery
}

func (user *User) GetChild(account string) User {
	var child User

	global.DB.Model(&User{}).
		Where("pid = ?", user.ID).
		Scopes(SearchAccount(account)).
		Limit(1).
		Find(&child)

	return child
}
