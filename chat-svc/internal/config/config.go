package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvValues struct {
	Port     string `mapstructure:"PORT"`
	DbUrl    string `mapstructure:"DB_URL"`
	KafkaUrl string `mapstructure:"KAFKA_URL"`
}

func init() {
	loadConfig()
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./internal/config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("error occured while reading env variables, error:", err)
	}

	err = viper.Unmarshal(&EnvValues)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}
}
