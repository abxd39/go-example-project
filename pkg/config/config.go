package config

import (
	"example-project/internal/domain"

	"github.com/spf13/viper"
)

var cfg domain.Config

func Load() domain.Config {
	return cfg
}

func init() {
	cf, err := loadConfig()
	if err != nil {
		panic(err.Error())
	}
	cfg = cf
}

func loadConfig() (domain.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return domain.Config{}, err
	}

	var cfg domain.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return domain.Config{}, err
	}

	return cfg, nil
}
