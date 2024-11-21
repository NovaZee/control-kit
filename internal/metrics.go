package internal

import (
	_logger "actuator/logger"
	"context"
	"encoding/json"
	"fmt"
	"github.com/NovaZee/control-kit/defs"
	"github.com/NovaZee/control-kit/osutil"
	"go.uber.org/zap"
	"time"
)

func (a *Actuator) ReportMetrics(ctx context.Context) error {
	ticker := time.NewTicker(time.Duration(a.config.ReportTime) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			metrics, err := osutil.CollectMetrics(a.config.PName)
			if err != nil {
				_logger.Ins.Error("Error collecting metrics", zap.Error(err))
				continue
			}
			// 写入 Etcd
			node := osutil.NodeMetrics{
				NodeID:    defs.NodeId,
				Metrics:   metrics,
				Timestamp: time.Now().Unix(),
			}
			err = a.reportMetricsToEtcd(defs.NodeId, &node)
			if err != nil {
				_logger.Ins.Error("Error reporting metrics to Etcd", zap.Error(err))
			}
		}
	}
}

// 将数据写入 Etcd
func (a *Actuator) reportMetricsToEtcd(nodeID string, metrics *osutil.NodeMetrics) error {
	key := fmt.Sprintf("%s/%s", defs.APPNodesMetricsPrefix, nodeID)
	metricsJson, err := json.Marshal(metrics)
	if err != nil {
		return err
	}

	value := metricsJson

	_, err = a.etcdX.Client.Put(context.Background(), key, string(value))
	if err != nil {
		return fmt.Errorf("failed to write to etcd: %v", err)
	}
	_logger.Ins.Info("Reported metrics to Etcd", zap.String("key", key), zap.String("value", string(value)))
	return nil
}
