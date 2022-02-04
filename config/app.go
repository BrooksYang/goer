package config

import (
	"time"
)

type App struct {
	Name     string `mapstructure:"name" json:"name"`
	Env      string `mapstructure:"env" json:"env"`
	Debug    bool   `mapstructure:"debug" json:"debug"`
	Port     uint   `mapstructure:"port" json:"port"`
	Timezone string `mapstructure:"timezone" json:"timezone"`
}

func (a App) SetTimezone() {
	time.Local, _ = time.LoadLocation(a.Timezone)
}

func (a App) IsLocal() bool {
	return a.Env == "local"
}

func (a App) IsProduction() bool {
	return a.Env == "production"
}
