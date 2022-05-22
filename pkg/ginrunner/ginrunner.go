package ginrunner

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kanztu/goblog/pkg/logger"
	"github.com/kanztu/goblog/pkg/server_context"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func Run(listenAddress string, router http.Handler, shutdownTimeout time.Duration) error {
	log := logger.InitLogger(logger.LEVEL_INFO, "Gin")
	srv := &http.Server{
		Addr:    listenAddress,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("Server Shutdown: %v", err)
	}

	return nil
}

func ResponseJSON(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		server_context.SrvCtx.Logger.Debug(data)
		ctx.JSON(http.StatusOK, data)
	} else {
		ctx.JSON(400, &Response{
			Error: err.Error(),
		})
	}
}

func Return404(ctx *gin.Context) {
	ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
