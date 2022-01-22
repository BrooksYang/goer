package config

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitConfig(path string, data interface{}) *viper.Viper {
	// Default config file
	cfgFile := "config.yaml"

	// Use config file from the flag.
	if path != "" {
		cfgFile = path
	}

	// Read config file
	v := viper.New()
	v.SetConfigFile(cfgFile)

	// Config type，support JSON, TOML, YAML, HCL, env file and Java properties config files
	v.SetConfigType("yaml")

	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in working directory & home directory
	v.AddConfigPath(".")
	v.AddConfigPath(home)

	v.SetEnvPrefix("goer")

	// Find and read the config file
	err = v.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
		return nil
	}

	// log
	log.Printf("read config file: %v\n", v.ConfigFileUsed())

	// Watch config file
	v.WatchConfig()

	// Log change
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed:", e.Name)
		if err := v.Unmarshal(data); err != nil {
			log.Println(err)
		}
	})
	if err := v.Unmarshal(data); err != nil {
		log.Println(err)
	}

	return v
}
