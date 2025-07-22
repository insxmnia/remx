package config

import (
	"remx/pkg/slogger"

	"github.com/spf13/viper"
)

func ReadConfig(params ConfigParams) (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigName(params.Name)
	config.SetConfigType("yaml")
	for _, value := range params.Path {
		config.AddConfigPath(value)
	}
	return config, config.ReadInConfig()
}

func MapConfig(viper *viper.Viper) *Configuration {
	var config Configuration
	if err := viper.Unmarshal(&config); err != nil {
		slogger.Fatal("configuration", "failed to unmarshal configuration into object", "error", err)
		return nil
	}

	return &config
}
