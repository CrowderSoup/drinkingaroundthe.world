package cmd

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"PORT"`
}

func loadConfig(paths ...string) (config Config, err error) {
	viper.SetConfigName(".drink.env")
	viper.SetConfigType("env")

	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	err = viper.ReadInConfig()
	viper.AutomaticEnv()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
