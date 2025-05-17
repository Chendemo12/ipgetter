package router

import (
	"time"

	"github.com/Chendemo12/fastapi"
	"github.com/Chendemo12/fastapi/middleware/fiberWrapper"
	"github.com/Chendemo12/fastapi/middleware/routers"
)

func New() *fastapi.Wrapper {
	app := fastapi.New(fastapi.Config{
		Version:     "v1.0.0",
		Description: "向请求者返回其当前的IP地址",
		Title:       "IP-GETTER",
	})

	// 底层采用fiber
	app.SetMux(fiberWrapper.Default())
	app.UsePrevious(BeforeValidate)
	app.UseBeforeWrite(PrintRequestLog)

	// 创建路由
	app.IncludeRouter(&IpRouter{}).
		IncludeRouter(routers.NewInfoRouter(app.Config())) // 开启默认基础路由

	return app
}

func BeforeValidate(c *fastapi.Context) error {
	c.Set("before-validate", time.Now())

	return nil
}

func PrintRequestLog(c *fastapi.Context) {
	fastapi.Info("请求耗时: ", time.Since(c.GetTime("before-validate")))
	fastapi.Info("响应状态码: ", c.Response().StatusCode)
}
