package config

import (
	"time"
)

type App struct {
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
	Env      string `mapstructure:"env" json:"env" yaml:"env"`
	Debug    bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	Port     uint   `mapstructure:"port" json:"port" yaml:"port"`
	Timezone string `mapstructure:"timezone" json:"timezone" yaml:"timezone"`
}

func (a App) SetTimezone() {
	time.Local, _ = time.LoadLocation(a.Timezone)
}
