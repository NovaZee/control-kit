package nodes

import (
	"context"
	"control-kit/config"
	"control-kit/internal/nodes/conf"
	"control-kit/internal/nodes/node"
	"github.com/NovaZee/control-kit/core"

	"time"
)

type Node struct {
	Nodes   map[string]struct{}
	Metrics *node.NodesWithMetrics
	Configs *conf.NodeConfig
	xetcd   *core.EtcdX
}

func BuildNode(ctx context.Context, xetcd *core.EtcdX, config *config.Config) *Node {
	n := Node{
		Nodes:   make(map[string]struct{}),
		xetcd:   xetcd,
		Metrics: node.NewMetricsHandler(ctx, xetcd),
		Configs: conf.NewNodeConfig(xetcd, time.Duration(config.RefreshTime)*time.Second),
	}

	return &n
}
