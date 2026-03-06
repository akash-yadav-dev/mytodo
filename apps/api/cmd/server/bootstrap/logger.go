package bootstrap

import (
	"log"
	"os"
)

// Package bootstrap handles application initialization.
//
// This file provides logging configuration and initialization.
//
// In production-grade applications, logger setup typically includes:
// - Structured logging (JSON format for production)
// - Log levels (debug, info, warn, error, fatal)
// - Log rotation and retention policies
// - Integration with logging aggregation (ELK, Splunk, Datadog)
// - Request correlation IDs
// - Performance considerations (async logging)
// - Sensitive data masking
//
// Example interface:
//   type Logger interface {
//       Debug(msg string, fields ...Field)
//       Info(msg string, fields ...Field)
//       Warn(msg string, fields ...Field)
//       Error(msg string, err error, fields ...Field)
//       Fatal(msg string, err error)
//       With(fields ...Field) Logger
//   }
//
// Example usage:
//   logger := NewLogger(config)
//   logger.Info("Server started", Field("port", 8080))
//   // Output: {"level":"info","msg":"Server started","port":8080,"time":"2026-03-06T..."}
//
//   logger.Error("Database connection failed", err, Field("host", "postgres"))
//   // Output: {"level":"error","msg":"Database connection failed","error":"...","host":"postgres"}

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
	fatalLog *log.Logger
}

func NewLogger() Logger {
	return &LoggerConfig{
		infoLog:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLog: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLog:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		fatalLog: log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
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
