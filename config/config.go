package config

import (
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Kafka KafkaConfig
}

type KafkaConfig struct {
	Brokers  string
	Topic    string
	ClientId string
}

func GetConfig() *AppConfig {
	var appConfig AppConfig

	v := viper.New()
	v.SetConfigFile("resources/config.yml")

	if err := v.ReadInConfig(); err != nil {
		panic("error while reading configuration")
	}

	env := os.Getenv("ENV")

	if env == "" {
		env = "local"
	}

	v = v.Sub(env)

	err := v.Unmarshal(&appConfig)
	if err != nil {
		panic("error while marshalling configuration")
	}

	if env == "stage" {
		appConfig.Kafka.Brokers = os.Getenv("KAFKA_BROKERS")
	}

	return &appConfig
}
