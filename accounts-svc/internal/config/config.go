package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Twilio OTP generation.....................
var TwilioAccountSid string
var TwilioAuthToken string
var TwilioServiceSid string

var EnvValues struct {
	Port  string `mapstructure:"PORT"`
	DbUrl string `mapstructure:"DB_URL"`
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
	// Twilio OTP generation.....................
	TwilioAccountSid = os.Getenv("TWILIO_ACCOUNT_SID")
	TwilioAuthToken = os.Getenv("TWILIO_AUTH_TOKEN")
	TwilioServiceSid = os.Getenv("TWILIO_SERVICE_SID")
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
