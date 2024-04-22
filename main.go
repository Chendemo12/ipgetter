package main

import (
	"github.com/Chendemo12/functools/environ"
	"github.com/Chendemo12/ipgetter/src/router"
)

func main() {
	listenHost := environ.GetString("LISTEN_HOST", "0.0.0.0")
	listenPort := environ.GetString("LISTEN_PORT", "7290")

	app := router.New()
	app.Run(listenHost, listenPort) // 阻塞运行
}
