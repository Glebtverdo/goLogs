package logs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

// todo вынести в env
const (
	appDir       = "app_log"
	errorDir     = "error_log"
	appLogName   = "app.log"
	errorLogName = "error.log"
)

var (
	appLogFile   *os.File
	errorLogFile *os.File
)

func checkDir(dirName string) error {
	if _, err := os.Stat(dirName); err != nil {
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func InitLoggers(logsPath string) {
	if err := checkDir(fmt.Sprintf("%s/%s", logsPath, appDir)); err != nil {
		panic(fmt.Sprintf("can not createDir %s as %s", logsPath, err.Error()))
	}
	if err := checkDir(fmt.Sprintf("%s/%s", logsPath, errorDir)); err != nil {
		panic(fmt.Sprintf("can not createDir %s as %s", logsPath, err.Error()))
	}

	initAppLogger(fmt.Sprintf("%s/%s/%s", logsPath, appDir, appLogName))
	initErrorLogger(fmt.Sprintf("%s/%s/%s", logsPath, errorDir, errorLogName))

	c := cron.New()
	_, err := c.AddFunc("@daily", func() {
		timestamp := time.Now().Format("02_01_2006_15_04")
		rotateLogFile(appDir, appLogName, "app_"+timestamp+".log", &appLogFile, initAppLogger)
		rotateLogFile(errorDir, errorLogName, "error_"+timestamp+".log", &errorLogFile, initErrorLogger)
	})
	if err != nil {
		log.Fatalf("Could not schedule log rotation: %v", err)
	}
	c.Start()
}
