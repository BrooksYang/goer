package config

type Code struct {
	Length int   `mapstructure:"length" json:"length"`
	TTL    int64 `mapstructure:"ttl" json:"ttl"`
}
