package internal

import (
	"actuator/config"
	_logger "actuator/logger"
	"actuator/utils"
	"context"
	"github.com/NovaZee/control-kit/core"
	"github.com/NovaZee/control-kit/defs"
	"go.uber.org/zap"
)

type Actuator struct {
	etcdX  *core.EtcdX
	config *config.Config
}

func SetUp(config *config.Config) (*Actuator, error) {
	var err error
	//if config.ExternalIp == "" {
	//	ip, err := utils.LookUpLocalIP()
	//	if err != nil {
	//		return nil, err
	//	}
	//	config.ExternalIp = ip
	//}
	//defs.NodeId = fmt.Sprintf("node-%d", utils.IpToUint32(config.ExternalIp))
	defs.NodeId, err = utils.GetMACBasedIdentifier()
	if err != nil {
		return nil, err
	}

	_logger.Ins.Info("External IP address", zap.String("ip", config.ExternalIp), zap.String("nodeId", defs.NodeId))

	actuator := &Actuator{
		config: config,
	}
	err = actuator.loadingComponents()
	if err != nil {
		return nil, err
	}

	return actuator, nil
}

func (a *Actuator) loadingComponents() error {
	x, err := core.BuildEtcdX(a.config.Etcd.Endpoints)
	if err != nil {
		return err
	}
	a.etcdX = x
	vi, err := a.SearchInI(context.Background())
	if err != nil {
		return err
	}
	// 上线
	err = a.RegisterAndKeepAlive(context.Background(), defs.NodeId, vi)
	if err != nil {
		return err
	}
	if a.config.Metrics {
		go func() {
			err = a.ReportMetrics(context.Background())
			if err != nil {
				_logger.Ins.Error("Error reporting metrics", zap.Error(err))
			}
		}()
	}
	return nil
}
