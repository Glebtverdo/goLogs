package initLoggers

import (
	"io"
	"log"

	"github.com/Glebtverdo/goLogs/models"
)

var Loggers = make(map[string]*log.Logger)

func InitLoggers(writers models.OutWriters, config models.LogsSettings) {
	for level, writer := range writers {
		setting := config[level]
		Loggers[level] = log.New(writer, "["+level+"] ", setting.Pattern)
	}
}

func GetLogger(level string) *log.Logger {
	if logger, ok := Loggers[level]; ok {
		return logger
	}
	return log.New(io.Discard, "["+level+"] ", log.LstdFlags)
}
