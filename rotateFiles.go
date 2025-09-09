package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/Glebtverdo/goLogs/models"
	"github.com/robfig/cron/v3"
)

func rotateLogs(config models.LogsSettings, writers models.OutWriters) error {
	for level, setting := range config {
		if setting.OutputType != "file" {
			continue
		}
		current := fmt.Sprintf("%s/%s", setting.Folder, setting.File)
		backup := fmt.Sprintf("%s/%s_%s.log", setting.Folder, level, logTimestamp())

		if f, ok := writers[level].(*os.File); ok {
			err := f.Close()
			Error(err.Error())
		}

		if err := copyFile(current, backup); err != nil {
			Error(err.Error())
			continue
		}

		if err := os.Truncate(current, 0); err != nil {
			Error(err.Error())
			continue
		}

		f, err := os.OpenFile(current, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			Error(err.Error())
			continue
		}
		writers[level] = f
	}
	return nil
}

func logTimestamp() string {
	return fmt.Sprintf("%02d_%02d_%d_%02d_%02d",
		time.Now().Day(), time.Now().Month(), time.Now().Year(),
		time.Now().Hour(), time.Now().Minute())
}

func copyFile(src, dst string) error {
	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}

func initLogRotationCron(config models.LogsSettings, writers models.OutWriters) {
	c := cron.New()

	for level, setting := range config {
		if setting.OutputType != "file" {
			continue
		}

		levelCopy := level

		_, err := c.AddFunc("@daily", func() {
			err := rotateLogs(config, writers)
			if err != nil {
				log.Printf("rotation failed for %s: %v", levelCopy, err)
			} else {
				log.Printf("rotated log for %s successfully", levelCopy)
			}
		})
		if err != nil {
			log.Printf("failed to schedule rotation for %s: %v", levelCopy, err)
		}
	}

	c.Start()
}
