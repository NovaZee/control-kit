package main

import (
	"actuator/config"
	"actuator/internal"
	_logger "actuator/logger"
	"flag"
	"fmt"

	"go.uber.org/zap"
	"os"
	"strings"
)

var cmd config.CMD

func init() {

	etcdPoints := flag.String("etcd", "", "etcd endpoints, separated by commas (required), e.g. http://127.0.0.1:2379,http://127.0.0.2:2379")
	externalIp := flag.String("externalIp", "", "external IP address")
	metrics := flag.Bool("metrics", true, "enable metrics")
	report := flag.Int64("reportTime", 30, "enable metrics")
	pname := flag.String("programName", "app", "program name")
	//保留 后续扩展
	cfg := flag.String("config", "", "config file path")
	flag.Parse()

	if *etcdPoints == "" && *cfg == "" {
		fmt.Println("Error: -etcd or config.yaml is required")
		flag.Usage()
		os.Exit(1)
	}

	cmd = config.CMD{
		EtcdPoints: strings.Split(*etcdPoints, ","),
		ExternalIp: *externalIp,
		Metrics:    *metrics,
		ReportTime: *report,
		PName:      *pname,
	}

}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	_logger.Ins = logger.Sugar()

	_logger.Ins.Info("actuator start", zap.Any("etcd", cmd.EtcdPoints), zap.String("externalIp", cmd.ExternalIp))

	conf := config.BuildConfig(cmd)
	_, err := internal.SetUp(conf)
	if err != nil {
		_logger.Ins.Fatal("actuator start failed", zap.Error(err))
		os.Exit(1)
	}
	select {}
}
