package make

import (
	"embed"
	"fmt"
	"strings"

	"goer/pkg/console"
	"goer/pkg/file"
	"goer/pkg/str"

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
	PackageName        string
}

//go:embed stubs
var stubsFS embed.FS

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	CmdMake.AddCommand(
		CmdMakeMigration,
	)
}

func makeModelFromString(name string) Model {
	model := Model{}
	model.StructName = str.Singular(strcase.ToCamel(name))
	model.StructNamePlural = str.Plural(model.StructName)
	model.TableName = str.Snake(model.StructNamePlural)
	model.VariableName = str.LowerCamel(model.StructName)
	model.PackageName = str.Snake(model.StructName)
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural)
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

	fmt.Printf("%s %s", ansi.Color("Created Migration: ", "green"), filePath)
}
