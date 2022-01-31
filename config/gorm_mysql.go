package config

import "fmt"

type Mysql struct {
	Host               string `mapstructure:"host" json:"host"`
	Port               string `mapstructure:"port" json:"port"`
	Database           string `mapstructure:"database" json:"database"`
	Username           string `mapstructure:"username" json:"username"`
	Password           string `mapstructure:"password" json:"password"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections" json:"max_idle_connections"`
	MaxOpenConnection  int    `mapstructure:"max_open_connection" json:"max_open_connection"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Database)
}
