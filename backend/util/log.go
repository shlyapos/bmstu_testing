package util

import (
	"log"
	"os"
)

type InfoLogger struct {
	logFileInfo *os.File
	logger      *log.Logger
}

func NewInfoLogger() *InfoLogger {
	infoLogger := new(InfoLogger)

	f, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}

	infoLogger.logFileInfo = f
	infoLogger.logger = log.New(f, "[INFO]\t", log.Ldate|log.Ltime)

	return infoLogger
}

func (l *InfoLogger) WriteInfoLog(status string, className string, method string) {
	l.logger.Printf("\t[STATUS] %v  \t[CLASS] %v\t[METHOD] %v", status, className, method)
}

func (l *InfoLogger) CloseAllLogFiles() {
	l.logFileInfo.Close()
}

var Logger = NewInfoLogger()
