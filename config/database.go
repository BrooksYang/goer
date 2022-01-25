package config

type Database struct {
	Connection string `mapstructure:"connection" json:"connection"`
	Mysql      Mysql  `mapstructure:"mysql" json:"mysql"`
	Sqlite     Sqlite `mapstructure:"sqlite" json:"sqlite"`
}
