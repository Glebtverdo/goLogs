package logs

import (
	"os"

	"github.com/Glebtverdo/goLogs/initLoggers"
	"github.com/Glebtverdo/goLogs/models"
)

var openFiles = map[string]*os.File{}

func InitLoggers(settings models.LogsSettings) error {
	writers, err := makeWriters(settings)
	if err != nil {
		return err
	}
	initLoggers.InitLoggers(writers, settings)
	initLogRotationCron(settings, writers)
	return nil
}
