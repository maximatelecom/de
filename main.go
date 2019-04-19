package main

import (
	"dataencrypter/config"
	"dataencrypter/dataencrypter"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.msk.vmet.ro/auth-backend/logutil"
	"google.golang.org/grpc"
	"net"
	"os"
	"path"
	"runtime"
)

// govvv variables
var (
	GitCommit string
	GitBranch string
	BuildDate string
	Version   string
	AppInfo   string
	AppName   string
)

var (
	configFilePath string
)

func init() {
	flag.StringVar(&configFilePath, "config", "./config.toml", "configuration file full path")

	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	AppName = fmt.Sprintf("%s/%s", path.Base(os.Args[0]), hostName)
	AppInfo = fmt.Sprintf(
		"%s (v%s %s %s-%s) %s",
		AppName,
		Version,
		BuildDate,
		GitCommit,
		GitBranch,
		runtime.Version(),
	)
}

func main() {
	flag.Parse()
	log.Warnf("Starting: %s", AppInfo)

	//viper
	v, err := config.NewViper(configFilePath)
	handleFatal(err, "unable to load config")

	config := config.NewConfig(v)
	// log
	handleFatal(logutil.InitLog(*config.Log), "unable to init logger")

	// crypto key and IV
	cryptoConf := config.Crypto
	if len(cryptoConf.Key) != 16 {
		log.Fatal("crypto.key length should be 16 symbols")
	}
	if len(cryptoConf.InitVector) != 12 {
		log.Fatal("crypto.init_vector length should be 12 symbols")
	}

	// grpc server
	grpcServerConfig := config.DataEncrypterServer
	lis, err := net.Listen("tcp", grpcServerConfig.ServerAddr)
	handleFatal(err, "unable to listen to tcp addr %s", grpcServerConfig.ServerAddr)
	defer lis.Close()
	grpcServer := grpc.NewServer()
	dataEncrypterServer := dataencrypter.NewServer(cryptoConf.Key, cryptoConf.InitVector)
	dataencrypter.RegisterDataEncrypterServer(grpcServer, dataEncrypterServer)
	handleFatal(grpcServer.Serve(lis), "unable to serve grpc")
}

// handleFatal handles fatal errors by logging them and exiting program
func handleFatal(err error, msg string, args ...interface{}) {
	if err != nil {
		log.
			WithError(err).
			Fatalf(msg, args...)
	}
}
