package config

type Config struct {
	App      App      `mapstructure:"app" json:"app"`
	Database Database `mapstructure:"database" json:"database"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt"`
	Paging   Paging   `mapstructure:"paging" json:"paging"`
	Common   Common   `mapstructure:"common" json:"common"`
}
