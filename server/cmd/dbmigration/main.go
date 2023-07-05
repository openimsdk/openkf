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

package main

import (
	"flag"
	"reflect"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/internal/conn/db"
	"github.com/OpenIMSDK/OpenKF/server/internal/models"
	"github.com/OpenIMSDK/OpenKF/server/internal/utils"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

func init() {
	// arg
	configPath := flag.String("c", "../../config.yaml", "config file path")
	flag.Parse()

	config.ConfigInit(*configPath)
	utils.OpenKFBanner()
	log.InitLogger()
	db.InitMysqlDB()
}

// migrate table.
func main() {
	// get db instance
	db := db.GetMysqlDB()

	// tables
	tables := []interface{}{
		models.User{},
	}

	// migrate
	for _, table := range tables {
		err := db.AutoMigrate(&table)
		if err != nil {
			log.Panicf("OpenKF Table Migration", "Migrate table %v... failed", reflect.TypeOf(table))
		}
		log.Infof("OpenKF Table Migration", "Migrate table %v... ok", reflect.TypeOf(table))
	}
}
