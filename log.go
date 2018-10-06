package main

import (
	"fmt"
	"log"
)

type Logger struct {
	Id string
}

func NewLogger(id string) *Logger {
	return &Logger{Id: id}
}

func (l *Logger) Format(level, format string) string {
	return fmt.Sprintf("%s [%s] %s", l.Id, level, format)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	log.Printf(l.Format("debug", format), args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	log.Printf(l.Format("warning", format), args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	log.Printf(l.Format("error", format), args...)
}
