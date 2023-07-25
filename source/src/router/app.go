package router

import (
	"github.com/Chendemo12/fastapi"
	"github.com/Chendemo12/fastapi-tool/logger"
)

// Config 配置文件类
type Config struct {
	HTTP struct {
		Host string `json:"host" yaml:"host"` // API host
		Port string `json:"port" yaml:"port"` // API port
	}
}

// ServiceContext 全局服务依赖
type ServiceContext struct {
	Conf   *Config
	Logger logger.Iface
}

func (c *ServiceContext) Config() any { return c.Conf }

// -------------------------------- 模型路由 --------------------------------

func New() *fastapi.Router {
	r := fastapi.APIRouter("/ip", []string{"IP GETTER"})
	{
		r.GET("", &IPModel{}, "返回当前请求的来源IP地址", GetAddress)
	}

	router := fastapi.APIRouter("/api", []string{})
	router.IncludeRouter(r)

	return router
}
