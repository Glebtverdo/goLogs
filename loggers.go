package logs

import "github.com/Glebtverdo/goLogs/initLoggers"

func Info(format string, v ...any) {
	initLoggers.GetLogger("info").Printf(format, v...)
}

func Debug(format string, v ...any) {
	initLoggers.GetLogger("debug").Printf(format, v...)
}

func Warn(format string, v ...any) {
	initLoggers.GetLogger("warn").Printf(format, v...)
}

func Error(format string, v ...any) {
	initLoggers.GetLogger("error").Printf(format, v...)
}

func Log(level string, format string, v ...any) {
	initLoggers.GetLogger(level).Printf(format, v...)
}
