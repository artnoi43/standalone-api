package config

import (
	"github.com/artworkk/standalone-api/api/auth"
	"github.com/artworkk/standalone-api/lib/postgres"
	"github.com/spf13/viper"
)

type Config struct {
	Port     string          `mapstructure:"listen_address"`
	Auth     auth.Config     `mapstructure:"auth"`
	Postgres postgres.Config `mapstructure:"postgres"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return config, nil
}
