package migrate

import (
	"fmt"
	"io/ioutil"

	"goer/global"
	"goer/pkg/database"
	"goer/pkg/file"

	"github.com/goer-project/goer-utils/console"
	"github.com/mgutz/ansi"
	"gorm.io/gorm"
)

type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

func NewMigrator() *Migrator {
	migrator := &Migrator{
		Folder:   global.MigrationsFolder + "/",
		DB:       global.DB,
		Migrator: global.DB.Migrator(),
	}

	migrator.createMigrationsTable()

	return migrator
}

func (migrator *Migrator) createMigrationsTable() {

	migration := Migration{}

	if !migrator.Migrator.HasTable(&migration) {
		_ = migrator.Migrator.CreateTable(&migration)
	}
}

func (migrator *Migrator) Up() {
	// Read migration files
	migrateFiles := migrator.readAllMigrationFiles()

	// Get batch
	batch := migrator.getBatch()

	var migrations []Migration
	migrator.DB.Find(&migrations)

	ran := false
	for _, migrationFile := range migrateFiles {
		if migrationFile.isNotMigrated(migrations) {
			migrator.runUpMigration(migrationFile, batch)
			ran = true
		}
	}

	if !ran {
		console.Info("Nothing to migrate.")
	}
}

func (migrator *Migrator) Rollback() {
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	var migrations []Migration
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)

	if !migrator.rollbackMigrations(migrations) {
		console.Info("Nothing to rollback.")
	}
}

func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {
	ran := false

	for _, _migration := range migrations {
		fmt.Printf("%s %s\n", ansi.Color("Rolling back:", "yellow"), _migration.Migration)

		migrationFile := getMigrationFile(_migration.Migration)
		if migrationFile.Down != nil {
			migrationFile.Down(global.DB.Migrator())
		}

		ran = true

		migrator.DB.Delete(&_migration)

		fmt.Printf("%s  %s\n", ansi.Color("Rolled back:", "green"), migrationFile.FileName)
	}

	return ran
}

func (migrator *Migrator) getBatch() int {
	batch := 1

	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}

	return batch
}

func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	files, err := ioutil.ReadDir(migrator.Folder)
	console.ExitIf(err)

	var migrateFiles []MigrationFile
	for _, f := range files {
		// Get filename
		fileName := file.GetFileNameWithoutExtension(f.Name())

		migrationFile := getMigrationFile(fileName)

		if len(migrationFile.FileName) > 0 {
			migrateFiles = append(migrateFiles, migrationFile)
		}
	}

	return migrateFiles
}

func (migrator *Migrator) runUpMigration(migrationFile MigrationFile, batch int) {
	if migrationFile.Up != nil {
		fmt.Printf("%s %s\n", ansi.Color("Migrating:", "yellow"), migrationFile.FileName)

		migrationFile.Up(global.DB.Migrator())

		fmt.Printf("%s  %s\n", ansi.Color("Migrated:", "green"), migrationFile.FileName)
	}

	err := migrator.DB.Create(&Migration{Migration: migrationFile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}

func (migrator *Migrator) Reset() {
	var migrations []Migration

	migrator.DB.Order("id DESC").Find(&migrations)

	if !migrator.rollbackMigrations(migrations) {
		console.Info("Nothing to rollback.")
	}
}

func (migrator *Migrator) Refresh() {
	migrator.Reset()

	migrator.Up()
}

func (migrator *Migrator) Fresh() {
	// Delete all tables
	err := database.DeleteAllTables()
	console.ExitIf(err)
	console.Info("Dropped all tables successfully.")

	// Migrate
	migrator.createMigrationsTable()
	console.Info("Migration table created successfully.")

	migrator.Up()
}
