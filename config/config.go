package config

import (
	"github.com/spf13/viper"
	"gitlab.msk.vmet.ro/auth-backend/logutil"
)

type Config struct {
	Crypto *Crypto
	DataEncrypterServer *DataEncrypterServer
	Log *logutil.Config
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		Crypto: NewCrypto(v),
		DataEncrypterServer: NewDataEncrypterServer(v),
		Log: NewLog(v),
	}
}
