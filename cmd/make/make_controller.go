package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeController = &cobra.Command{
	Use:   "controller",
	Short: "Create a new controller",
	Run:   runMakeController,
	Args:  cobra.ExactArgs(1),
}

func runMakeController(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	if model.Directory == "" {
		model.PackageName = "controllers"
	}

	dir := fmt.Sprintf("app/http/controllers/%s", model.Directory)

	// mkdir -p, 0777
	_ = os.MkdirAll(dir, os.ModePerm)

	// Create file
	createFileFromStub(dir+model.VariableNameSnake+".go", "controller", model)
}
