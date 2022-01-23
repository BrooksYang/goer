package migrate

import (
	"github.com/spf13/cobra"
)

var CmdMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all database migrations",
	Args:  cobra.NoArgs,
	Run:   runMigrateReset,
}

func runMigrateReset(cmd *cobra.Command, args []string) {
	migrator().Reset()
}
