package util

import (
	"github.com/spf13/viper"
	"time"
)

// Config stores all configuration items
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	Port                string        `mapstructure:"PORT"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

const (
	LayoutISODOB = "2006-01-02"
	Layout3      = "2015-09-15T14:00:12-00:00"
)

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

//ProcessDateTime process time input as string
func ProcessDateTime(layout, input string) (*time.Time, error) {
	switch layout {
	case "dob":
		time, err := time.Parse(LayoutISODOB, input)
		if err != nil {
			return nil, err
		}
		return &time, nil
	case "rfc":
		time, err := time.Parse(time.RFC3339, input)
		if err != nil {
			return nil, err
		}
		return &time, nil
	default:
		time, err := time.Parse(time.Kitchen, input)
		if err != nil {
			return nil, err
		}
		return &time, nil
	}
}
