package config


type UserServiceConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`

}
type RedisConfig struct {
	HOST         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}
type ServerConfig struct {
	Name string  `mapstructure:"name"`
	Port int `mapstructure:"port"`
	UserServiceInfo UserServiceConfig `mapstructure:"user_srv"`
	Redis RedisConfig `mapstructure:"redis"`
}