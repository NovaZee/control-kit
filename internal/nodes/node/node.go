package node

import (
	"context"
	"control-kit/internal/nodes/util"
	_logger "control-kit/logger"
	"encoding/json"
	"github.com/NovaZee/control-kit/core"
	"github.com/NovaZee/control-kit/defs"
	"github.com/NovaZee/control-kit/informer"
	"github.com/NovaZee/control-kit/osutil"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"maps"
	"sync"
	"time"
)

type NodesWithMetrics struct {
	nodes map[string]map[string]*osutil.NodeMetrics
	etcdx *core.EtcdX
	i     *informer.Informer
	mu    sync.RWMutex
}

func NewMetricsHandler(ctx context.Context, etcdx *core.EtcdX) *NodesWithMetrics {
	ms := &NodesWithMetrics{
		nodes: make(map[string]map[string]*osutil.NodeMetrics),
		etcdx: etcdx,
	}
	go ms.startRefreshing(10)
	ms.i = informer.NewInformer(ctx, etcdx, ms, 10, defs.APPNodesMetricsPrefix)

	ms.i.Start()
	ms.i.Watch()

	return ms
}

func (m *NodesWithMetrics) OnAdd(key string, obj interface{}) {
	var sm osutil.NodeMetrics
	if err := json.Unmarshal(obj.([]byte), &sm); err != nil {
		_logger.Ins.Error("Failed to unmarshal newItem", zap.Error(err))
		return
	}
	m.add(key, &sm)
	_logger.Ins.Info("OnUpdate", zap.String("key", key), zap.Any("obj", sm))
}
func (m *NodesWithMetrics) OnUpdate(key string, oldItem, newItem interface{}) {
	var sm osutil.NodeMetrics
	if err := json.Unmarshal(newItem.([]byte), &sm); err != nil {
		_logger.Ins.Error("Failed to unmarshal newItem", zap.Error(err))
		return
	}
	m.add(key, &sm)
	_logger.Ins.Info("OnUpdate", zap.String("key", key), zap.Any("oldItem", sm), zap.Any("newItem", sm))
}
func (m *NodesWithMetrics) OnDelete(key string, res interface{}) {
	_logger.Ins.Info("OnDelete", zap.String("key", key), zap.Any("res", res))
}

func (m *NodesWithMetrics) add(key string, obj *osutil.NodeMetrics) {
	m.mu.Lock()
	defer m.mu.Unlock()
	key = util.FetchNodeId(key)
	ips := m.nodes[key]
	for s := range ips {
		ips[s] = obj
	}
}

func (m *NodesWithMetrics) delete(key string, obj *osutil.SystemMetrics) {
	m.mu.Lock()
	defer m.mu.Unlock()
	key = util.FetchNodeId(key)
	if _, ok := m.nodes[key]; ok {
		delete(m.nodes, key)
	}
}

func (m *NodesWithMetrics) listNodes() error {
	get, err := m.etcdx.Client.Get(context.Background(), defs.APPNodesOnlinePrefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	currentNodes := make(map[string]struct{})
	for _, kv := range get.Kvs {
		nodeId := util.FetchNodeId(string(kv.Key))
		currentNodes[nodeId] = struct{}{}
		if _, ok := m.nodes[nodeId]; ok {
			continue
		}
		var register defs.Register
		if err := json.Unmarshal(kv.Value, &register); err != nil {
			_logger.Ins.Error("Failed to unmarshal etcd data", zap.Error(err))
			continue
		}
		m2 := make(map[string]*osutil.NodeMetrics)
		for i := range register.BindIp {
			m2[register.BindIp[i]] = &osutil.NodeMetrics{}
		}
		m.nodes[nodeId] = m2
	}

	//剔除下线的
	for nodeId := range m.nodes {
		if _, ok := currentNodes[nodeId]; !ok {
			delete(m.nodes, nodeId)
		}
	}
	_logger.Ins.Info("Online nodes ", zap.Any("nodes", maps.Keys(currentNodes)))
	return nil
}

func (m *NodesWithMetrics) startRefreshing(sec int) {
	ticker := time.NewTicker(time.Duration(sec) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := m.listNodes()
			if err != nil {
				_logger.Ins.Error("Failed to list nodes", zap.Error(err))
				continue
			}
		}
	}
}
