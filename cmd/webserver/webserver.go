package main

import (
	"flag"
	"time"

	"github.com/kanztu/goblog/internal/routers"
	"github.com/kanztu/goblog/pkg/config"
	"github.com/kanztu/goblog/pkg/db"
	"github.com/kanztu/goblog/pkg/ginrunner"
	"github.com/kanztu/goblog/pkg/server_context"
)

func main() {
	var (
		configFile string
		port       string
	)

	flag.StringVar(&port, "port", ":9097", "server listen and serve address, like 127.0.0.1:8888 or :8080")
	flag.StringVar(&configFile, "config", "./config/config.yaml", "the config file path")

	config.LoadGlobalConfig(configFile)
	server_context.InitServerContext()
	db.InitDB()

	r := routers.NewRouters()
	if err := ginrunner.Run(port, r, 15*time.Second); err != nil {
		server_context.SrvCtx.Logger.Fatalf("Run server failed: %v", err)
	}
}
