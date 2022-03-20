package cmd

import (
	"goapp/bootstrap"
	"goapp/global"

	"github.com/goer-project/goer-core/config"
	"github.com/goer-project/goer/cmd/make"
	"github.com/goer-project/goer/cmd/migrate"
	"github.com/goer-project/goer/database"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goer",
	Short: "Api framework in Golang",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add sub command (only for development)
	rootCmd.AddCommand(
		make.CmdMake,
		migrate.CmdMigrate,
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config.InitConfig(cfgFile, &global.Config) // Init viper
	bootstrap.Logger()                         // Init logger
	global.Config.App.SetTimezone()            // Init timezone
	database.DB = database.Gorm()              // Init database
	global.DB = database.DB
	bootstrap.Redis() // Init redis
	bootstrap.Cache() // Init cache
}
