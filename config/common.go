package config

type Common struct {
	UidLength       int    `mapstructure:"uid_length" json:"uid_length"`
	DefaultPassword string `mapstructure:"default_password" json:"default_password"`
}
