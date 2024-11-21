package internal

import (
	"context"
	"control-kit/config"
	"control-kit/internal/api"
	"control-kit/internal/nodes"
	"fmt"
	"github.com/NovaZee/control-kit/core"
	"github.com/gin-gonic/gin"
	"log"
)

type ControlPlane struct {
	etcdX  *core.EtcdX
	config *config.Config
	node   *nodes.Node
	router *gin.Engine
}

func SetUp(config *config.Config) (*ControlPlane, error) {
	var err error
	ctx := context.Background()
	var cp = &ControlPlane{}

	x, err := core.BuildEtcdX(config.Etcd.Endpoints)
	if err != nil {
		return nil, err
	}

	cp.etcdX = x
	cp.node = nodes.BuildNode(ctx, x, config)
	cp.config = config

	cp.router = cp.api()

	if err = cp.router.Run(fmt.Sprintf(":%d", config.Port)); err != nil {
		log.Fatalf("run web server error: %v", err)
	}

	return cp, nil
}

func (c *ControlPlane) api() *gin.Engine {
	return api.InitController(api.NewNodeHandler(c.node.Metrics), api.NewConfigHandler(c.node.Configs))
}
