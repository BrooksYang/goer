package migrate

import (
	"github.com/spf13/cobra"
)

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run the database migrations",
	Args:  cobra.NoArgs,
	Run:   runMigrateUp,
}

func runMigrateUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}
