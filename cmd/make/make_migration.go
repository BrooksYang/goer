package make

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a new migration file",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(1),
}

func runMakeMigration(cmd *cobra.Command, args []string) {
	timeStr := time.Now().Format("2006_01_02_150405")

	model := makeModelFromString(args[0])
	fileName := timeStr + "_" + model.PackageName
	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)

	createFileFromStub(filePath, "migration", model, map[string]string{"{{FileName}}": fileName})
}
