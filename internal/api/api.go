package api

import (
	"github.com/gin-gonic/gin"
	"io"
)

type Module interface {
	Register(g gin.IRoutes, middle ...gin.HandlerFunc)
}

func InitController(
	nodeHandler *NodeHandler,
	configHandler *ConfigHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	engine := gin.Default()
	engine.Use(Logger(), gin.Recovery(), CORS())

	api := engine.Group("/api/").Use()
	//b := engine.Group("/b").Use(internalMiddle.BasicAuth())
	nodeHandler.Register(api)
	configHandler.Register(api)
	return engine
}

type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResponse(code int, success bool, data interface{}, msg string) *Response {
	return &Response{
		Code:    code,
		Success: success,
		Data:    data,
		Message: msg,
	}
}
