package config

type Aes struct {
	Key string `mapstructure:"key" json:"key"`
	Iv  string `mapstructure:"iv" json:"iv"`
}
