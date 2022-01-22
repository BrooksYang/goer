package config

type Database struct {
	Connection string `mapstructure:"connection" json:"connection" yaml:"connection"`
	Mysql      Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Sqlite     Sqlite `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
}
