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

package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

var d *gorm.DB

type writer struct{}

// Write implement log writer interface.
func (w writer) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// InitMysqlDB init mysql connection.
func InitMysqlDB() {
	// try to use default database [mysql]
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.Username,
		config.Config.Mysql.Password,
		config.Config.Mysql.Ip,
		config.Config.Mysql.Port,
		"mysql", // sys database
	)

	// connect to mysql
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		// retry
		time.Sleep(time.Duration(10) * time.Second)
		db, err = gorm.Open(mysql.Open(dsn), nil)
		if err != nil {
			log.Panic("Mysql", err.Error(), " open failed ", dsn)
		}
	}

	// create database if not exists
	sql := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS %s default charset utf8 COLLATE utf8_general_ci;",
		config.Config.Mysql.Database,
	)
	err = db.Exec(sql).Error
	if err != nil {
		log.Panic("Mysql", err.Error(), " Exec failed ", sql)
	}

	// connect to openkf database
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.Username,
		config.Config.Mysql.Password,
		config.Config.Mysql.Ip,
		config.Config.Mysql.Port,
		config.Config.Mysql.Database, // sys database
	)

	logger := logger.New(
		writer{},
		logger.Config{
			IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
			Colorful:                  true, // Disable color
		},
	)
	// connect to openkf database
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		log.Panic("Mysql", err.Error(), " Open failed ", dsn)
	}

	// set mysql connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic("Mysql", err.Error(), " db.DB() failed ")
	}

	// set connect result
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.Config.Mysql.MaxLifetime))
	sqlDB.SetMaxOpenConns(config.Config.Mysql.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.Config.Mysql.MaxIdleConns)

	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	db.Set("gorm:table_options", "collation=utf8_unicode_ci")

	d = db
	log.Info("Mysql", "connect ok", dsn)
}

// GetMysqlDB get mysql connection.
func GetMysqlDB() *gorm.DB {
	return d
}

// CloseMysqlDB close mysql connection.
func CloseMysqlDB() {
	if d == nil {
		return
	}

	sqlDB, err := d.DB()
	if err != nil {
		log.Error("Mysql", err.Error(), " db.DB() failed ")
	} else {
		err = sqlDB.Close()
		if err != nil {
			log.Error("Mysql", err.Error(), " sqlDB.Close() failed ")
		}
	}
}
