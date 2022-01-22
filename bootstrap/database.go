package bootstrap

import (
	"log"

	"goer/global"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConfig gorm.Dialector

// Gorm init gorm
func Gorm() *gorm.DB {
	switch global.Config.Database.Connection {
	case "mysql":
		GormMysqlConfig()
	case "sqlite":
		GormSqliteConfig()
	default:
		GormMysqlConfig()
	}

	db, err := gorm.Open(dbConfig, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Println(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.Config.Database.Mysql.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(global.Config.Database.Mysql.MaxOpenConnection)

	return db
}

func GormMysqlConfig() {
	dbConfig = mysql.New(mysql.Config{
		DSN: global.Config.Database.Mysql.Dsn(),
	})
}

func GormSqliteConfig() {
	dbConfig = sqlite.Open(global.Config.Database.Sqlite.Database)
}
