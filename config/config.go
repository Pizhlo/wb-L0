package config

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

type Config struct {
	DBAddress string `mapstructure:"DB_ADDRESS"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	conf := Config{}

	_, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)
	defer cancel()

	err := viper.ReadInConfig()
	if err != nil {
		return conf, err
	}

	err = viper.Unmarshal(&conf)

	return conf, nil
}
