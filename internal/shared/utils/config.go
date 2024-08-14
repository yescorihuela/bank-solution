package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppPort        string `mapstructure:"APP_PORT"`
	BlueSoftDBDSN  string `mapstructure:"BLUESOFT_DB_DSN"`
	MigrationsPath string `mapstructure:"MIGRATIONS_PATH"`
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
