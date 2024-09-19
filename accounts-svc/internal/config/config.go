package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var EnvValues struct {
	Port           string `mapstructure:"PORT"`
	DbUrl          string `mapstructure:"DB_URL"`
	AgenciesSvcUrl string `mapstructure:"AGENCIES_SVC_URL"`
	Environment    string `mapstructure:"ENVIRONMENT"`
}

var JWT struct {
	ExpTimeInMinutes int `mapstructure:"JwtTokenExpiryInMinutes"`
}

var Twilio struct {
	AccountSid string `mapstructure:"TWILIO_ACCOUNT_SID"`
	AuthToken  string `mapstructure:"TWILIO_AUTH_TOKEN"`
	ServiceSid string `mapstructure:"TWILIO_SERVICE_SID"`
}

var Emailing struct {
	FromEmail         string `mapstructure:"EMAIL_FROM"`
	AppPassword       string `mapstructure:"EMAIL_APP_PASSWORD"`
	SmtpServerAddress string `mapstructure:"SMTP_SERVER_ADDRESS"`
	SmtpsPort         string `mapstructure:"SMTPS_PORT"`
}

var DevMode struct {
	ByPassTwilio bool
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

	err = viper.Unmarshal(&Emailing)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}

	DevMode.ByPassTwilio = (viper.Get("DEVMODE_BYPASS_TWILIO") == "true")
	fmt.Println("DevMode.TwiliioBypass=", DevMode.ByPassTwilio)

}
