package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var EnvValues struct {
	Port           string `mapstructure:"PORT"`
	AccountsSvcUrl string `mapstructure:"AUTH_SVC_URL"`
	AgenciesSvcUrl string `mapstructure:"AGENCIES_SVC_URL"`
	ChatSvcUrl     string `mapstructure:"CHAT_SVC_URL"`
	LogToFile      bool   `mapstructure:"LOG_TO_FILE"`
	PprofUrl       string `mapstructure:"PPROF_URL"`
	KafkaUrl       string `mapstructure:"KAFKA_URL"`
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
		log.Fatalln("error occured while writing env values onto variagbles, error:", err)
	}
	if EnvValues.LogToFile {
		err = os.Setenv("LOG_TO_FILE", "true")
		if err != nil {
			log.Fatalln("error occurred while setting LOG_TO_FILE environment variable, error:", err)
		}
	}
}
