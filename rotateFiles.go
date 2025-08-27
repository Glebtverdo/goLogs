package logs

import (
	"fmt"
	"io"
	"log"
	"os"
)

func rotateLogFile(dir, currentFile, backupFile string, logFile **os.File, initFunc func(string)) {
	path := fmt.Sprintf("%s/%s", dir, currentFile)
	backup := fmt.Sprintf("%s/%s", dir, backupFile)

	if *logFile != nil {
		(*logFile).Close()
	}

	if err := copyFile(path, backup); err != nil {
		log.Fatalf("Failed to backup %s: %v", currentFile, err)
	}

	_ = os.Truncate(path, 0)
	initFunc(currentFile)
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
