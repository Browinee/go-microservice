package config


type UserServiceConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Name string `mapstructure:"name"`

}
type RedisConfig struct {
	HOST         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
type ServerConfig struct {
	Name string  `mapstructure:"name"`
	Port int `mapstructure:"port"`
	UserServiceInfo UserServiceConfig `mapstructure:"user_srv"`
	Redis RedisConfig `mapstructure:"redis"`
	Consul ConsulConfig `mapstructure:"consul" json:"consul"`
}