package logs

import "log"

var (
	infoLogger  *log.Logger
	debugLogger *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

func Info(format string, v ...any) {
	if infoLogger != nil {
		infoLogger.Printf(format, v...)
	}
}

func Debug(format string, v ...any) {
	if debugLogger != nil {
		debugLogger.Printf(format, v...)
	}
}

func Warn(format string, v ...any) {
	if warnLogger != nil {
		warnLogger.Printf(format, v...)
	}
}

func Error(format string, v ...any) {
	if errorLogger != nil {
		errorLogger.Printf(format, v...)
	}
}
