package main

import (
	"control-kit/config"
	"control-kit/internal"
	_logger "control-kit/logger"
	"flag"
	"go.uber.org/zap"
	"os"
)

var configFile string

func init() {

	flag.StringVar(&configFile, "c", "config.yaml", "choose config file.")
	flag.Parse()

}
func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	_logger.Ins = logger.Sugar()

	zap.ReplaceGlobals(logger)

	cf, err := config.LoadConfig(configFile)
	if err != nil {
		_logger.Ins.Fatal(err)
		os.Exit(1)
	}
	_, err = internal.SetUp(cf)
	if err != nil {
		_logger.Ins.Fatal(err)
		os.Exit(1)
	}
}
