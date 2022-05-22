package logger

import (
	"fmt"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type LogLevel log.Level

type Logger struct {
	log.Entry
}

type Formatter struct {
	log.TextFormatter
}

const (
	LEVEL_DEBUG LogLevel = (LogLevel)(log.DebugLevel)
	LEVEL_WARN  LogLevel = (LogLevel)(log.WarnLevel)
	LEVEL_INFO  LogLevel = (LogLevel)(log.InfoLevel)
	LEVEL_ERROR LogLevel = (LogLevel)(log.ErrorLevel)
)

func (f *Formatter) Format(entry *log.Entry) ([]byte, error) {
	pc := make([]uintptr, 1)
	n := runtime.Callers(6, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	app := entry.Data["app"]
	msg := "time=%s level=%s app=%s fn=(%s:%d)(%s) msg='%s'\n"
	return []byte(fmt.Sprintf(
		msg,
		entry.Time.Format(f.TimestampFormat),
		entry.Level,
		app,
		frame.File,
		frame.Line,
		strings.Split(frame.Function, "/")[len(strings.Split(frame.Function, "/"))-1],
		entry.Message,
	)), nil

}

func InitLogger(l LogLevel, module string) *log.Entry {
	logger := log.New()
	logger.SetLevel((log.Level)(l))
	logger.SetFormatter(&Formatter{log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02-15:04:05.999999999TZ07",
	},
	})
	return logger.WithField("app", module)
}
