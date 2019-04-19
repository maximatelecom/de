package config

import (
	"github.com/spf13/viper"
	"strings"
)

func NewViper(configFilePath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(configFilePath)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}
