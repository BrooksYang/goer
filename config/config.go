package config

type Config struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
