// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package log

import (
	"bufio"
	"io"
	"os"
	"time"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var _logger *logrus.Logger

func InitLogger() {
	_logger = loggerInit()
}

func loggerInit() *logrus.Logger {
	var logger = logrus.New()

	// set log level
	if config.Config.App.Debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	// set output
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err.Error())
	}
	writer := bufio.NewWriter(src)

	// write to stdout and file
	logger.SetOutput(io.MultiWriter(os.Stdout, writer))

	// set formatter
	logger.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        false,
		FieldsOrder:     []string{"PID", "FilePath", "OperationID"},
	})

	// add print fileline hook
	logger.AddHook(newFilelineHook())

	// add rotate file hook, default is 24h
	logger.AddHook(NewLfsHook(24*time.Hour, 7))

	return logger
}

func NewLfsHook(rotationTime time.Duration, maxRemainNum uint) logrus.Hook {
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: initRotateLogs(rotationTime, maxRemainNum, "all"),
		logrus.InfoLevel:  initRotateLogs(rotationTime, maxRemainNum, "all"),
		logrus.WarnLevel:  initRotateLogs(rotationTime, maxRemainNum, "all"),
		logrus.ErrorLevel: initRotateLogs(rotationTime, maxRemainNum, "all"),
	}, &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        false,
		FieldsOrder:     []string{"PID", "FilePath", "OperationID"},
	})
	return lfsHook
}

func initRotateLogs(rotationTime time.Duration, maxRemainNum uint, level string) *rotatelogs.RotateLogs {
	writer, err := rotatelogs.New(
		config.Config.App.LogFile+level+".%Y%m%d%H%M",
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(maxRemainNum),
	)
	if err != nil {
		panic(err.Error())
	} else {
		return writer
	}
}

func GetLogger() *logrus.Logger {
	return _logger
}

func Info(Operation string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Infoln(args...)
}
func Error(Operation string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Errorln(args...)
}
func Debug(Operation string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Debugln(args...)
}
func Panic(Operation string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Panicln(args...)
}

func Infof(Operation string, format string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Infof(format, args...)
}
func Errorf(Operation string, format string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Errorf(format, args...)
}
func Debugf(Operation string, format string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Debugf(format, args...)
}
func Panicf(Operation string, format string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Panicf(format, args...)
}
