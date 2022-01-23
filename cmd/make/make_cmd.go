package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1),
}

func runMakeCMD(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])
	if model.Directory == "" {
		model.PackageName = "cmd"
	}

	dir := fmt.Sprintf("cmd/%s", model.Directory)

	// mkdir -p, 0777
	_ = os.MkdirAll(dir, os.ModePerm)

	createFileFromStub(dir+model.VariableName+".go", "cmd", model)
}
