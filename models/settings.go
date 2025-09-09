package models

import (
	"io"
)

// type logsFolders struct {
// 	Debug string
// 	Error string
// 	Warn  string
// 	Info  string
// }

// type logsFiles struct {
// 	DebugFile string
// 	ErrorFile string
// 	InfoFile  string
// 	WarnFile  string
// }

type LogsSetting struct {
	OutputType string
	Folder     string
	File       string
	Pattern    int // log.Ldate|log.Ltime|log.Lshortfile
}

type LogsSettings map[string]LogsSetting

// type LogsSettings struct {
// 	OutputType  string // file, stdout
// 	LogsDir     string
// 	LogsFolders map[string]string
// 	LogsFile    map[string]string
// }

type OutWriters map[string]io.Writer

// type OutWritters struct {
// 	DebugOut io.Writer
// 	ErrorOut io.Writer
// 	InfoOut  io.Writer
// 	WarnOut  io.Writer
// }
