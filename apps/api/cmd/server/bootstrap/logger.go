package bootstrap

import (
	"log"
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
	Warn(msg string)
	Fatal(msg string)
}

type LoggerConfig struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	debugLog *log.Logger
	warnLog  *log.Logger
}

func NewLogger() Logger {
	return &LoggerConfig{
		infoLog:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLog: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLog:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *LoggerConfig) Info(msg string) {
	l.infoLog.Println(msg)
}

func (l *LoggerConfig) Error(msg string) {
	l.errorLog.Println(msg)
}

func (l *LoggerConfig) Debug(msg string) {
	l.debugLog.Println(msg)
}

func (l *LoggerConfig) Warn(msg string) {
	l.warnLog.Println(msg)
}

func (l *LoggerConfig) Fatal(msg string) {
	l.errorLog.Fatal(msg)
}
