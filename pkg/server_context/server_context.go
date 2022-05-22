package server_context

import (
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"

	log "github.com/kanztu/goblog/pkg/logger"
)

type ServerContext struct {
	Logger *logrus.Entry
	DB     *xorm.Engine
}

var SrvCtx ServerContext

func InitServerContext() {
	SrvCtx.Logger = log.InitLogger(log.LEVEL_DEBUG, "gin")
}
