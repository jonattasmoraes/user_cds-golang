package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	return &Logger{
		debug:   log.New(writer, fmt.Sprintf("\033[34m%s\033[0m", "DEBUG: "), log.Ldate|log.Ltime),
		info:    log.New(writer, fmt.Sprintf("\033[32m%s\033[0m", "INFO: "), log.Ldate|log.Ltime),
		warning: log.New(writer, fmt.Sprintf("\033[33m%s\033[0m", "WARNING: "), log.Ldate|log.Ltime),
		err:     log.New(writer, fmt.Sprintf("\033[31m%s\033[0m", "ERROR: "), log.Ldate|log.Ltime),
		writer:  writer,
	}
}

func colorText(color, text string) string {
	return fmt.Sprintf("%s%s\033[0m", color, text)
}

// Create Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(colorText("\033[34m", fmt.Sprint(v...)))
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(colorText("\033[32m", fmt.Sprint(v...)))
}
func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(colorText("\033[33m", fmt.Sprint(v...)))
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(colorText("\033[31m", fmt.Sprint(v...)))
}

// Create Format Enabled Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(colorText("\033[34m", fmt.Sprintf(format, v...)))
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(colorText("\033[32m", fmt.Sprintf(format, v...)))
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(colorText("\033[33m", fmt.Sprintf(format, v...)))
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(colorText("\033[31m", fmt.Sprintf(format, v...)))
}
