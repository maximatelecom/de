package config

import "github.com/spf13/viper"

type Crypto struct {
	Key        []byte
	InitVector []byte
}

func NewCrypto(v *viper.Viper) *Crypto {
	return &Crypto{
		Key:        []byte(v.GetString("crypto.key")),
		InitVector: []byte(v.GetString("crypto.init_vector")),
	}
}
