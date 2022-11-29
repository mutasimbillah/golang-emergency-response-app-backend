package util

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Env variables load failed ...")
		return
	}

	err = viper.Unmarshal(&config)
	return
}
