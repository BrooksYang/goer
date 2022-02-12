package config

type Open struct {
	Enabled   bool     `mapstructure:"enabled" json:"enabled"`
	ApiKey    string   `mapstructure:"api_key" json:"api_key"`
	ApiSecret string   `mapstructure:"api_secret" json:"api_secret"`
	Ip        []string `mapstructure:"ip" json:"ip"`
	TTL       int64    `mapstructure:"ttl" json:"ttl"`
}
