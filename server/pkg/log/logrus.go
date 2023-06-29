// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// InitLogger init logger
func InitLogger() {
	_logger = loggerInit()
}

func loggerInit() *logrus.Logger {
	logger := logrus.New()

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

// NewLfsHook add fileline hook
func NewLfsHook(rotationTime time.Duration, maxRemainNum uint) logrus.Hook {
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: initRotateLogs(rotationTime, maxRemainNum),
		logrus.InfoLevel:  initRotateLogs(rotationTime, maxRemainNum),
		logrus.WarnLevel:  initRotateLogs(rotationTime, maxRemainNum),
		logrus.ErrorLevel: initRotateLogs(rotationTime, maxRemainNum),
	}, &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        false,
		FieldsOrder:     []string{"PID", "FilePath", "OperationID"},
	})

	return lfsHook
}

func initRotateLogs(rotationTime time.Duration, maxRemainNum uint) *rotatelogs.RotateLogs {
	writer, err := rotatelogs.New(
		config.Config.App.LogFile+".%Y%m%d%H%M",
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(maxRemainNum),
	)
	if err != nil {
		panic(err.Error())
	} else {
		return writer
	}
}

// GetLogger get logger instance
func GetLogger() *logrus.Logger {
	return _logger
}

// Info log info
func Info(Operation string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Infoln(args...)
}

// Error log error
func Error(Operation string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Errorln(args...)
}

// Debug log debug
func Debug(Operation string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Debugln(args...)
}

// Panic log panic
func Panic(Operation string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Panicln(args...)
}

// Infof log info with format
func Infof(Operation string, format string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Infof(format, args...)
}

// Errorf log error with format
func Errorf(Operation string, format string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Errorf(format, args...)
}

// Debugf log debug with format
func Debugf(Operation string, format string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Debugf(format, args...)
}

// Panicf log panic with format
func Panicf(Operation string, format string, args ...interface{}) {
	_logger.WithFields(logrus.Fields{
		"Operation": Operation,
	}).Panicf(format, args...)
}
