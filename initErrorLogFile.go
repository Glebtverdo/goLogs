package logs

import (
	"log"
	"os"
)

func initErrorLogger(fileName string) {
	errorLogFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open error log file: %v", err)
	}
	if errorLogger != nil {
		errorLogger.SetOutput(errorLogFile)
	} else {
		errorLogger = log.New(errorLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}
