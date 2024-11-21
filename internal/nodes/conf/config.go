package conf

import (
	"context"
	"control-kit/internal/nodes/util"
	_logger "control-kit/logger"
	"github.com/NovaZee/control-kit/core"
	"github.com/NovaZee/control-kit/defs"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"sync"
	"time"
)

type NodeConfig struct {
	refreshTime time.Duration
	config      map[string]string
	etcdx       *core.EtcdX
	mu          sync.RWMutex
}

func NewNodeConfig(etcdx *core.EtcdX, refreshTime time.Duration) *NodeConfig {
	nc := &NodeConfig{
		refreshTime: refreshTime,
		config:      make(map[string]string),
		etcdx:       etcdx,
	}
	go nc.startRefreshing()

	return nc
}

func (nc *NodeConfig) startRefreshing() {
	ticker := time.NewTicker(nc.refreshTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			nc.refreshNodes()
		}
	}
}

func (nc *NodeConfig) refreshNodes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := nc.etcdx.Client.Get(ctx, defs.APPNodesConfigPrefix, clientv3.WithPrefix())
	if err != nil {
		_logger.Ins.Error("Failed to get data from etcd", zap.Error(err))
		return
	}

	configs := make(map[string]string)

	for _, kv := range resp.Kvs {
		configs[util.FetchNodeId(string(kv.Key))] = string(kv.Value)
	}
	nc.mu.Lock()
	defer nc.mu.Unlock()
	nc.config = configs
	_logger.Ins.Info("sync config success ! ")
}
