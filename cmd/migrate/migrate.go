package migrate

import (
	"goapp/database/migrations"
	"goapp/global"
	"goapp/pkg/migrate"

	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run the database migrations",
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
		CmdMigrateRefresh,
		CmdMigrateReset,
		CmdMigrateFresh,
	)
}

func migrator() *migrate.Migrator {
	// Init migration files
	migrations.Initialize()

	return migrate.NewMigrator(global.DB, global.MigrationsFolder+"/")
}
