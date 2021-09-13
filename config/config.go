package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	Port        int           `mapstructure:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"userSrv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}
