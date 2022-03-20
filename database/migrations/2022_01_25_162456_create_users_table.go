package migrations

import (
	"goapp/app/models"

	"github.com/goer-project/goer/migrate"
	"gorm.io/gorm"
)

type CreateUsersTable struct {
	models.BaseModel

	Uid          string `gorm:"type:varchar(255);unique"`
	CountryId    int64  `gorm:"type:tinyint(4);default:0;index"`
	Name         string `gorm:"type:varchar(255)"`
	Username     string `gorm:"type:varchar(255);unique;default:null"`
	Email        string `gorm:"type:varchar(255);unique;default:null"`
	Phone        string `gorm:"type:varchar(255);unique;default:null"`
	Password     string `gorm:"type:varchar(255)"`
	PayPassword  string `gorm:"type:varchar(255)"`
	GoogleKey    string `gorm:"type:varchar(255)"`
	GoogleStatus string `gorm:"type:enum('unbind','disabled','enabled');not null;default:unbind;index"`
	Sso          string `gorm:"type:varchar(255)"`
	Pid          uint64 `gorm:"type:int(10);not null;default:0;index"`
	InviteCount  uint64 `gorm:"type:int(10);not null;default:0"`
	Depth        int64  `gorm:"type:int(10);not null;default:0"`
	IsValid      bool   `gorm:"type:tinyint(4);not null;default:true;index"`
	KycStatus    int64  `gorm:"type:tinyint(4);not null;default:0;index;comment:KYC status, 0-not verify, 1-pending, 2-verified, 3-failed"`
	Age          int64  `gorm:"type:tinyint(4);not null;default:0;index"`
	Gender       string `gorm:"type:enum('secret', 'male', 'female');not null;default:secret;index"`
	Avatar       string `gorm:"type:varchar(255);not null;'"`

	models.TimestampsField
}

func (CreateUsersTable) TableName() string {
	return "users"
}

func init() {
	up := func(migrator gorm.Migrator) {
		_ = migrator.AutoMigrate(&CreateUsersTable{})
	}

	down := func(migrator gorm.Migrator) {
		_ = migrator.DropTable(&CreateUsersTable{})
	}

	migrate.Add("2022_01_25_162456_create_users_table", up, down)
}
