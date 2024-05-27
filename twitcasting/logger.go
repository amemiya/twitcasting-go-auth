package twitcasting

import (
	"log"
	"strings"
)

type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
}

// BasicLogger sample implementation of Logger
type BasicLogger struct {
	Logger *log.Logger
}

func (basicLogger *BasicLogger) Debug(v ...interface{}) {
	basicLogger.Logger.Printf(strings.Repeat("%v ", len(v)-1)+"%v", v...)
}

func (basicLogger *BasicLogger) Info(v ...interface{}) {
	basicLogger.Logger.Printf(strings.Repeat("%v ", len(v)-1)+"%v", v...)
}

func (basicLogger *BasicLogger) Warn(v ...interface{}) {
	basicLogger.Logger.Printf(strings.Repeat("%v ", len(v)-1)+"%v", v...)
}

func (basicLogger *BasicLogger) Error(v ...interface{}) {
	basicLogger.Logger.Printf(strings.Repeat("%v ", len(v)-1)+"%v", v...)
}
