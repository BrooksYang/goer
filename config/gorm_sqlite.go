package config

type Sqlite struct {
	Database string `mapstructure:"database" json:"database"`
}
