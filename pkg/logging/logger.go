// Package logging
// @Description: 自定义日志库
package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	InfoLevel  = "INFO"
	DebugLevel = "DEBUG"
	FatalLevel = "Fatal"
)

var (
	logger *logrus.Logger
)

func InitLogger() {
	logger = logrus.New()
	// 打印文件名，行数以及对应的方法名
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 设置时间戳格式
	})
	// 日志同时发送到文件以及控制台当中
	writer1 := os.Stdout
	writer2, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logger.WithFields(logrus.Fields{"err": err.Error()}).Fatal("open file error")
	}
	logger.SetOutput(io.MultiWriter(writer1, writer2))
}

func SetLevel(level string) {
	switch level {
	case DebugLevel:
		logger.SetLevel(logrus.DebugLevel)
		break
	case InfoLevel:
		logger.SetLevel(logrus.InfoLevel)
		break
	case FatalLevel:
		logger.SetLevel(logrus.FatalLevel)
		break
	}
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func InfoWithFields(fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Info(args)
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func DebugWithFields(fields map[string]interface{}, args ...interface{}) {
	logger.WithFields(fields).Debug(args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func ErrorWithFields(fields map[string]interface{}, args ...interface{}) {
	logger.WithFields(fields).Error(args)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

func FatalWithFields(fields map[string]interface{}, args ...interface{}) {
	logger.WithFields(fields).Fatal(args)
}
