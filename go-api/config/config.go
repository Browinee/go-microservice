package config


type UserServiceConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`

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
	UserConfig UserServiceConfig `mapstructure:"user_service"`
	Redis RedisConfig `mapstructure:"redis"`
	Consul ConsulConfig `mapstructure:"consul" json:"consul"`

}

type NacosConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port uint64 `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	User string`mapstructure:"user" json:"user"`
	Password string`mapstructure:"password" json:"password"`
	DataId string `mapstructure:"dataId" json:"dataId"`
	Group string `mapstructure:"group" json:"group"`
}