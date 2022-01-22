package config

import "fmt"

type Mysql struct {
	Host               string `mapstructure:"host" json:"host" yaml:"host"`
	Port               string `mapstructure:"port" json:"port" yaml:"port"`
	Database           string `mapstructure:"database" json:"database" yaml:"database"`
	Username           string `mapstructure:"username" json:"username" yaml:"username"`
	Password           string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConnections int    `mapstructure:"max-idle-connections" json:"maxIdleConnections" yaml:"max-idle-connections"`
	MaxOpenConnection  int    `mapstructure:"max-open-connections" json:"maxOpenConnections" yaml:"max-open-connections"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4", m.Username, m.Password, m.Host, m.Port, m.Database)
}
