package api

import (
	"basic/config"
	"basic/log"

	"basic/api/middleware"

	"github.com/gin-gonic/gin"
)

var Routes = func(engine *gin.Engine) {}

// APIRun 启动 http api 服务
func APIRun(conf *config.Config) error {
	serveConf := conf.Server
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	api := gin.Default()

	// 通用中间件
	middleware.Common(api, conf)

	Routes(api)

	serverAddr := serveConf.Addr
	log.Infof("API Server Listening: %s", serverAddr)
	return api.Run(serverAddr)
}