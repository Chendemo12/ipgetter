package main

import (
	"github.com/Chendemo12/fastapi"
	"github.com/Chendemo12/fastapi-tool/logger"
	"github.com/Chendemo12/functools/environ"
	"github.com/Chendemo12/functools/zaplog"
	"github.com/Chendemo12/ipgetter/src/router"
)

func main() {
	listenHost := environ.GetString("LISTEN_HOST", "0.0.0.0")
	listenPort := environ.GetString("LISTEN_PORT", "7290")
	debug := environ.GetBool("DEBUG", false)

	zc := &zaplog.Config{
		Filename:   AppName,
		Level:      zaplog.WARNING,
		Rotation:   10,
		Retention:  5,
		MaxBackups: 10,
		Compress:   false,
	}

	if debug {
		zc.Level = zaplog.DEBUG
	}

	fc := &router.Config{}
	svc := &router.ServiceContext{Conf: fc, Logger: logger.NewDefaultLogger()}
	fc.HTTP.Host = listenHost
	fc.HTTP.Port = listenPort

	app := fastapi.New(AppName, VERSION, true, svc)
	app.DisableMultipleProcess().
		SetLogger(svc.Logger).
		SetDescription("向请求者返回其当前的IP地址").
		SetShutdownTimeout(5).
		IncludeRouter(router.New())

	app.Run(listenHost, listenPort) // 阻塞运行
}
