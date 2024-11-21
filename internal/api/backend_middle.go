package api

import (
	_logger "control-kit/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"time"
)

// Logger Logger中间件 集成到自己的日志库
func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 执行时间
		nowTime := time.Now()

		// 获取客户端真实 IP
		clientIP := c.GetHeader("X-Forwarded-For")
		if clientIP == "" {
			clientIP = c.GetHeader("X-Real-IP")
		}
		if clientIP == "" {
			clientIP = c.Request.RemoteAddr
		} else {
			// 如果 X-Forwarded-For 包含多个 IP 地址，取第一个
			clientIP = strings.Split(clientIP, ",")[0]
		}
		c.Next()
		_logger.Ins.Info(" [Gin Middle] http request", zap.Any(" request", c.Request.URL),
			zap.String("ip", clientIP),
			zap.Any("latency", fmt.Sprintf("%vms", time.Since(nowTime).Nanoseconds()/1e6)),
		)
	}
}

func CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 允许 Origin 字段中的域发送请求
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		// 设置预验请求有效期为 86400 秒
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 设置允许请求的方法
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		// 设置允许请求的 Header
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Api-Token,X-Data,X-Requested-With,X-Auth-Token")
		// 设置拿到除基本字段外的其他字段，如上面的Apitoken, 这里通过引用Access-Control-Expose-Headers，进行配置，效果是一样的。
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Headers")
		// 配置是否可以带认证信息
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// OPTIONS请求返回200
		if context.Request.Method == "OPTIONS" {
			fmt.Println(context.Request.Header)
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
	}
}
