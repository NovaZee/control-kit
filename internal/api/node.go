package api

import (
	"control-kit/internal/nodes/node"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NodeHandler struct {
	no *node.NodesWithMetrics
}

func NewNodeHandler(n *node.NodesWithMetrics) *NodeHandler {
	return &NodeHandler{no: n}
}

func (n *NodeHandler) Register(g gin.IRoutes, middle ...gin.HandlerFunc) {

	wd := g.Use(middle...)
	{
		wd.GET("/nodes/online", func(ctx *gin.Context) { n.ListOnlineNode(ctx) })
		wd.GET("/nodes/node", func(ctx *gin.Context) { n.GetNode(ctx) })
	}
	wd.Use()
}

func (n *NodeHandler) ListOnlineNode(ctx *gin.Context) {
	nodes := n.no.GetNodes()
	if nodes == nil || len(nodes) == 0 {
		ctx.JSON(http.StatusOK, NewResponse(http.StatusOK, false, nil, "No online nodes"))
		return
	}
	ctx.JSON(http.StatusOK, NewResponse(http.StatusOK, true, nodes, "Online nodes retrieved successfully"))
}

func (n *NodeHandler) GetNode(ctx *gin.Context) {
	// 从表单中获取 node 和 config 字段
	nodeId := ctx.Query("nodeId")
	if nodeId == "" {
		ctx.JSON(http.StatusBadRequest, NewResponse(http.StatusBadRequest, false, nil, "nodeId and config are required"))
		return
	}
	res := n.no.GetNode(nodeId)

	if res == nil {
		ctx.JSON(http.StatusOK, NewResponse(http.StatusOK, false, nil, "No online nodes"))
		return
	}
	ctx.JSON(http.StatusOK, NewResponse(http.StatusOK, true, res, "Online nodes retrieved successfully"))
}
