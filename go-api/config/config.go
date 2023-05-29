package config


type UserServiceConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}
type ServerConfig struct {
	Name string  `mapstructure:"name"`
	Port int `mapstructure:"port"`
	UserServiceInfo UserServiceConfig `mapstructure:"user_srv"`
}