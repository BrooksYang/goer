package config

type JWT struct {
	SecretKey string `mapstructure:"secret_key" json:"secret_key"`
	TTL       int64  `mapstructure:"ttl" json:"ttl"`
}
