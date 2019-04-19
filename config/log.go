package config

import (
	"github.com/spf13/viper"
	"gitlab.msk.vmet.ro/auth-backend/logutil"
)

func NewLog(v *viper.Viper) *logutil.Config {
	return &logutil.Config{
		Level: v.GetString("log.level"),
		Syslog: v.GetString("log.syslog"),
	}
}
