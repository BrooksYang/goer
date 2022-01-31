package user

import (
	"goer/global"
	"goer/pkg/helpers"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Generate uid
func GenerateUid() string {
	length := global.Config.Common.UidLength

	return helpers.RandomNumber(length)
}

// Check password
func CheckPassword(user User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	}

	return true
}

// Check pay password
func CheckPayPassword(user User, payPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PayPassword), []byte(payPassword))
	if err != nil {
		return false
	}

	return true
}

// Check google code
func CheckGoogleCode(user User, googleCode string) bool {
	if user.GoogleStatus != string(GoogleStatusEnabled) {
		return true
	}

	// todo::check google code
	return true
}

// Check 2fa
func Check2FA(user User, payPassword string, googleCode string) bool {
	// Check pay password
	if !CheckPayPassword(user, payPassword) {
		return false
	}

	// Check google code
	return CheckGoogleCode(user, googleCode)
}

// Check if account exists
func AccountExists(user User) bool {
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

func GetChildrenIdSubQuery(user User, account string) *gorm.DB {
	subQuery := global.DB.Model(&User{}).
		Select("id").
		Where("pid", user.ID).
		Scopes(SearchAccount(account))

	return subQuery
}

func GetChild(user User, account string) User {
	var child User

	global.DB.Model(&User{}).
		Where("pid = ?", user.ID).
		Scopes(SearchAccount(account)).
		Limit(1).
		Find(&child)

	return child
}
