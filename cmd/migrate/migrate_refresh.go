package migrate

import (
	"github.com/spf13/cobra"
)

var CmdMigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migrations",
	Args:  cobra.NoArgs,
	Run:   runMigrateRefresh,
}

func runMigrateRefresh(cmd *cobra.Command, args []string) {
	migrator().Refresh()
}
