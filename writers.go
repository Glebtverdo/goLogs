package logs

import (
	"fmt"
	"os"

	"github.com/Glebtverdo/goLogs/models"
)

func initFileLogger(folder string, file string) (*os.File, error) {
	fullPath := fmt.Sprintf("%s/%s", folder, file)
	if err := os.MkdirAll(folder, 0755); err != nil {
		return nil, fmt.Errorf("failed to create dir %s: %w", folder, err)
	}
	f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", fullPath, err)
	}
	return f, err
}

func makeWriters(config models.LogsSettings) (models.OutWriters, error) {
	writers := make(models.OutWriters)
	var err error
	for level, setting := range config {
		switch setting.OutputType {
		case "stdout":
			writers[level] = os.Stdout
		case "stderr":
			writers[level] = os.Stderr
		case "file":
			writers[level], err = initFileLogger(setting.File, setting.Folder)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown output type %s for level %s", setting.OutputType, level)
		}
	}

	return writers, nil
}
