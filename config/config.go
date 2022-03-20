package config

type Config struct {
	Aes      Aes      `mapstructure:"aes" json:"aes"`
	App      App      `mapstructure:"app" json:"app"`
	Common   Common   `mapstructure:"common" json:"common"`
	Database Database `mapstructure:"database" json:"database"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt"`
	Mail     Mail     `mapstructure:"mail" json:"mail"`
	Swag     Swag     `mapstructure:"swag" json:"swag"`
	Open     Open     `mapstructure:"open" json:"open"`
	Code     Code     `mapstructure:"code" json:"code"`
}
