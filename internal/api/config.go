package api

import (
	"control-kit/internal/nodes/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfigHandler struct {
	Nc *conf.NodeConfig
}

func NewConfigHandler(nc *conf.NodeConfig) *ConfigHandler {
	return &ConfigHandler{Nc: nc}
}

func (c *ConfigHandler) Register(g gin.IRoutes, middle ...gin.HandlerFunc) {

	wd := g.Use(middle...)
	{
		wd.GET("/config", func(ctx *gin.Context) { c.getConfig(ctx) })
		wd.POST("/config", func(ctx *gin.Context) { c.editConfig(ctx) })
	}
	wd.Use()
}

func (c *ConfigHandler) getConfig(ctx *gin.Context) {
	nodeID := ctx.Query("nodeId")
	if nodeID == "" {
		ctx.JSON(http.StatusBadRequest, NewResponse(http.StatusBadRequest, false, nil, "nodeId is required"))
		return
	}
	// 可以根据 nodeID 处理逻辑，返回相应的配置信息
	// 此处假设 c.Nc 包含所需的信息
	ctx.JSON(http.StatusOK, NewResponse(http.StatusOK, true, c.Nc.GetTargetConfig(nodeID), "Config retrieved successfully"))
}

func (c *ConfigHandler) editConfig(ctx *gin.Context) {
	// 从表单中获取 node 和 config 字段
	node := ctx.PostForm("nodeId")
	config := ctx.PostForm("config")
	if node == "" || config == "" {
		ctx.JSON(http.StatusBadRequest, NewResponse(http.StatusBadRequest, false, nil, "nodeId and config are required"))
		return
	}
	targetConfig := c.Nc.EditTargetConfig(node, config)
	ctx.JSON(http.StatusOK, NewResponse(http.StatusOK, true, targetConfig, "Config updated successfully"))
}
