package config

import "fmt"

type Mysql struct {
	Host               string `mapstructure:"host" json:"host"`
	Port               string `mapstructure:"port" json:"port"`
	Database           string `mapstructure:"database" json:"database"`
	Username           string `mapstructure:"username" json:"username"`
	Password           string `mapstructure:"password" json:"password"`
	MaxIdleConnections int    `mapstructure:"max-idle-connections" json:"maxIdleConnections"`
	MaxOpenConnection  int    `mapstructure:"max-open-connections" json:"maxOpenConnections"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4", m.Username, m.Password, m.Host, m.Port, m.Database)
}
