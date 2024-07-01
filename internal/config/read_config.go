package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	SERVER_PORT string
	LISTEN_ADDR string
}

func ReadConfig() AppConfig {
	var c AppConfig

	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./internal/config/")
	v.AddConfigPath(".")

	v.SetDefault("SERVER_PORT", "80")
	v.SetDefault("LISTEN_ADDR", "0.0.0.0")

	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = v.Unmarshal(&c)

	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %w", err))
	}

	return c
}
