package mylogger

import (
	"fmt"
	"time"
)

// 往终端写入日志

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewConsoleLogger 构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, arg ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, arg...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s %s %d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
	}
}

// Debug ...
func (c ConsoleLogger) Debug(format string, arg ...interface{}) {
	c.log(DEBUG, format, arg...)
}

// Trace ...
func (c ConsoleLogger) Trace(format string, arg ...interface{}) {
	c.log(TRACE, format, arg...)
}

// Info ...
func (c ConsoleLogger) Info(format string, arg ...interface{}) {
	c.log(INFO, format, arg...)
}

// Warning ...
func (c ConsoleLogger) Warning(format string, arg ...interface{}) {
	c.log(WARNING, format, arg...)
}

// Error ...
func (c ConsoleLogger) Error(format string, arg ...interface{}) {
	c.log(ERROR, format, arg...)
}

// Fatal ...
func (c ConsoleLogger) Fatal(format string, arg ...interface{}) {
	c.log(FATAL, format, arg...)
}
