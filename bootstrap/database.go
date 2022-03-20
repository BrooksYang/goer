package bootstrap

import (
	"log"
	"os"
	"time"

	"goapp/global"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	// Log level
	logLevel := logger.Warn
	if global.Config.App.IsProduction() {
		logLevel = logger.Error
	}

	// Custom logger
	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             500 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:                  logLevel,               // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,                   // 禁用彩色打印
		},
	)

	// Open DB
	db, err := gorm.Open(dbConfig, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   customLogger,
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
