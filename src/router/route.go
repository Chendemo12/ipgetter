package router

import (
	"github.com/Chendemo12/fastapi"
	"github.com/Chendemo12/functools/logger"
	"github.com/gofiber/fiber/v2"
)

type IPModel struct {
	fastapi.BaseModel
	IP     string `json:"ip" description:"IPv4地址"`
	Detail struct {
		IPv4     string `json:"IPv4" description:"IPv4地址"`
		IPv4Full string `json:"IPv4_full" description:"带端口的IPv4地址"`
		Ipv6     string `json:"IPv6" description:"IPv6地址"`
	} `json:"detail" description:"详细信息"`
}

func (m *IPModel) SchemaDesc() string { return "IP信息" }

type IpRouter struct {
	fastapi.BaseGroupRouter
}

func (r *IpRouter) Tags() []string { return []string{"IP GETTER"} }

func (r *IpRouter) Prefix() string { return "/api" }

func (r *IpRouter) Summary() map[string]string {
	return map[string]string{
		"GetIp": "返回当前请求的来源IP地址",
	}
}

func (r *IpRouter) GetIp(c *fastapi.Context) (*IPModel, error) {
	clientIP := c.MuxContext().ClientIP()
	info := &IPModel{}
	info.IP = clientIP
	info.Detail.IPv4 = clientIP
	info.Detail.IPv4Full = clientIP

	fiberCtx, ok := c.MX().(*fiber.Ctx)
	if ok {
		info.Detail.IPv4Full = fiberCtx.Context().RemoteAddr().String()
	}

	logger.Debugf("client address: %s, full: %s", clientIP, info.Detail.IPv4Full)

	return info, nil
}
