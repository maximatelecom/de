package config

import "github.com/spf13/viper"

type DataEncrypterServer struct {
	ServerAddr string
}

func NewDataEncrypterServer(v *viper.Viper) *DataEncrypterServer {
	return &DataEncrypterServer{
		ServerAddr: v.GetString("grpc.server_addr"),
	}
}
