package main

import (
	"flag"
	"time"

	_ "github.com/kanztu/goblog/docs"
	"github.com/kanztu/goblog/internal/routers"
	"github.com/kanztu/goblog/pkg/config"
	"github.com/kanztu/goblog/pkg/db"
	"github.com/kanztu/goblog/pkg/ginrunner"
	"github.com/kanztu/goblog/pkg/server_context"
)

// @title           Goblog API
// @version         1.0
// @description     Backend api for the server

// @host      localhost:9098
// @BasePath  /api/v1

func main() {
	var (
		configFile string
		port       string
	)

	flag.StringVar(&port, "port", "localhost:9098", "server listen and serve address, like 127.0.0.1:8888 or :8080")
	flag.StringVar(&configFile, "config", "./config/config.yaml", "the config file path")

	config.LoadGlobalConfig(configFile)
	server_context.InitServerContext()
	db.InitDB()
	server_context.SrvCtx.Logger.Debugf("Run server in: %v", port)
	r := routers.NewRouters()
	if err := ginrunner.Run(port, r, 15*time.Second); err != nil {
		server_context.SrvCtx.Logger.Fatalf("Run server failed: %v", err)
	}
}
