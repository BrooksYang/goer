package make

import (
	"fmt"
	"os"
	"time"

	"goer/global"

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
	dir := global.MigrationsFolder
	filePath := fmt.Sprintf("%s/%s.go", dir, fileName)

	// mkdir -p, 0777
	_ = os.MkdirAll(dir, os.ModePerm)

	createFileFromStub(filePath, "migration", model, map[string]string{"{{FileName}}": fileName})
}
