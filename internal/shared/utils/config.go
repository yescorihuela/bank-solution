package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppHTTPPort      int    `mapstructure:"APP_HTTP_PORT"`
	BlueSoftURLDB    string `mapstructure:"BLUESOFT_URL_DB"`
	MigrationsPath   string `mapstructure:"MIGRATIONS_PATH"`
	MaxDBConnections int    `mapstructure:"MAX_DB_CONNECTIONS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
