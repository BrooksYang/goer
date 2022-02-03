package config

type Redis struct {
	Host          string `mapstructure:"host" json:"host"`
	Port          string `mapstructure:"port" json:"port"`
	Password      string `mapstructure:"password" json:"password"`
	Database      int    `mapstructure:"database" json:"database"`
	DatabaseCache int    `mapstructure:"database_cache" json:"database_cache"`
}
