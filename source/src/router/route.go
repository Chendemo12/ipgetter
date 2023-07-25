package router

import (
	"github.com/Chendemo12/fastapi"
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

func (m IPModel) SchemaDesc() string { return "IP信息" }

func GetAddress(c *fastapi.Context) *fastapi.Response {
	info := &IPModel{}
	info.Detail.IPv4Full = c.EngineCtx().Context().RemoteAddr().String()

	fiberIP := c.EngineCtx().IP()
	headerIP := c.EngineCtx().Get("X-Forwarded-For")

	if fiberIP == headerIP || headerIP == "" {
		info.IP = fiberIP
		info.Detail.IPv4 = fiberIP
	} else {
		info.IP = headerIP
		info.Detail.IPv4 = headerIP
	}

	c.Logger().Debug("fiber think: ", fiberIP, " X-Forwarded-For: ", headerIP)

	return c.OKResponse(info)
}
