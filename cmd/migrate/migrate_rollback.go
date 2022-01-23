package migrate

import (
	"github.com/spf13/cobra"
)

var CmdMigrateRollback = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback the last database migration",
	Args:  cobra.NoArgs,
	Run:   runMigrateRollback,
}

func runMigrateRollback(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}
