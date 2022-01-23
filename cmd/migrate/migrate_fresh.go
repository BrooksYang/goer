package migrate

import (
	"github.com/spf13/cobra"
)

var CmdMigrateFresh = &cobra.Command{
	Use:   "fresh",
	Short: "Drop all tables and re-run all migrations",
	Args:  cobra.NoArgs,
	Run:   runMigrateFresh,
}

func runMigrateFresh(cmd *cobra.Command, args []string) {
	migrator().Fresh()
}
