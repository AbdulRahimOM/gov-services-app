package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvValues struct {
	Port           string `mapstructure:"PORT"`
	DbUrl          string `mapstructure:"DB_URL"`
	AgenciesSvcUrl string `mapstructure:"AGENCIES_SVC_URL"`
}

var JWT struct {
	ExpTimeInMinutes int `mapstructure:"JwtTokenExpiryInMinutes"`
}

var Twilio struct {
	AccountSid string `mapstructure:"TWILIO_ACCOUNT_SID"`
	AuthToken  string `mapstructure:"TWILIO_AUTH_TOKEN"`
	ServiceSid string `mapstructure:"TWILIO_SERVICE_SID"`
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

	err = viper.Unmarshal(&Twilio)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}

	err = viper.Unmarshal(&JWT)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}
}
