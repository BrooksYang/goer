package config

type Config struct {
	Aes      Aes      `mapstructure:"aes" json:"aes"`
	App      App      `mapstructure:"app" json:"app"`
	Common   Common   `mapstructure:"common" json:"common"`
	Database Database `mapstructure:"database" json:"database"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt"`
	Paging   Paging   `mapstructure:"paging" json:"paging"`
}
