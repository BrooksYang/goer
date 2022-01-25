package config

type Config struct {
	App      App      `mapstructure:"app" json:"app"`
	Database Database `mapstructure:"database" json:"database"`
}
