package migrate

import (
	"gorm.io/gorm"
)

type migrationFunc func(gorm.Migrator)

var migrationFiles []MigrationFile

type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}

func getMigrationFile(name string) MigrationFile {
	for _, migrationFile := range migrationFiles {
		if name == migrationFile.FileName {
			return migrationFile
		}
	}

	return MigrationFile{}
}

func (migrationFile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == migrationFile.FileName {
			return false
		}
	}

	return true
}
