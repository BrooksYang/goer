package make

import (
	"embed"
	"fmt"
	"strings"

	"goer/pkg/file"

	"github.com/goer-project/goer-utils/console"
	"github.com/goer-project/goer-utils/helpers"
	"github.com/iancoleman/strcase"
	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
)

type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	VariableNameSnake  string
	PackageName        string
	Directory          string
}

//go:embed stubs
var stubsFS embed.FS

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	CmdMake.AddCommand(
		CmdMakeCMD,
		CmdMakeController,
		CmdMakeRequest,
		CmdMakeMiddleware,
		CmdMakeMigration,
		CmdMakeModel,
	)
}

func makeModelFromString(path string) Model {
	arr := strings.Split(path, "/")
	name := arr[len(arr)-1]

	model := Model{}
	model.StructName = helpers.Singular(strcase.ToCamel(name))
	model.StructNamePlural = helpers.Plural(model.StructName)
	model.VariableNameSnake = helpers.Snake(model.StructName)
	model.TableName = helpers.Snake(model.StructNamePlural)
	model.VariableName = helpers.LowerCamel(model.StructName)
	model.PackageName = helpers.Snake(model.StructName)
	model.VariableNamePlural = helpers.LowerCamel(model.StructNamePlural)

	// Directory
	directorArr := arr[:len(arr)-1]
	model.Directory = strings.Join(directorArr, "/")
	if model.Directory != "" {
		model.Directory = model.Directory + "/"
		model.PackageName = helpers.Snake(directorArr[len(directorArr)-1])
	}

	// Table name
	model.TableName = strings.Replace(model.TableName, "create_", "", 1)
	model.TableName = strings.Replace(model.TableName, "_tables", "", 1)

	return model
}

func createFileFromStub(filePath string, stubName string, model Model, variables ...interface{}) {

	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)
	}

	if file.Exists(filePath) {
		console.Exit(filePath + " already exists!")
	}

	// Read stub
	modelData, _ := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	modelStub := string(modelData)

	// Replace
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{PackageName}}"] = model.PackageName
	replaces["{{TableName}}"] = model.TableName

	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}

	err := file.Put([]byte(modelStub), filePath)
	if err != nil {
		console.Exit(err.Error())
	}

	fmt.Printf("%s %s\n", ansi.Color("Created:", "green"), filePath)
}
