package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Create a new model file",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1),
}

func runMakeModel(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])
	if model.Directory == "" {
		model.PackageName = "models"
	}

	dir := fmt.Sprintf("app/models/%s", model.Directory)

	// mkdir -p, 0777
	_ = os.MkdirAll(dir, os.ModePerm)

	// Create file
	createFileFromStub(dir+model.VariableName+".go", "model", model)
}
