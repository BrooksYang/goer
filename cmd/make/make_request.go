package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create a new form request",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1),
}

func runMakeRequest(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	if model.Directory == "" {
		model.PackageName = "requests"
	}

	dir := fmt.Sprintf("app/http/requests/%s", model.Directory)

	// mkdir -p, 0777
	_ = os.MkdirAll(dir, os.ModePerm)

	// Create file
	createFileFromStub(dir+model.VariableNameSnake+".go", "request", model)
}
