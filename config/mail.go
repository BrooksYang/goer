package config

type Mail struct {
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	FromAddress string `mapstructure:"from_address" json:"from_address"`
	FromName    string `mapstructure:"from_name" json:"from_name"`
}
