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

	"github.com/openimsdk/openkf/server/internal/config"
	"github.com/openimsdk/openkf/server/internal/conn/db"
	customerroles "github.com/openimsdk/openkf/server/internal/models/customer_roles"
	systemroles "github.com/openimsdk/openkf/server/internal/models/system_roles"
	"github.com/openimsdk/openkf/server/internal/utils"
	"github.com/openimsdk/openkf/server/pkg/log"
)

var (
	configPath string
	isDrop     bool
)

func init() {
	// arg
	flag.StringVar(&configPath, "c", "../../config.yaml", "config file path")
	flag.BoolVar(&isDrop, "d", false, "drop table if exist")
	flag.Parse()

	config.ConfigInit(configPath)
	utils.OpenKFBanner()
	log.InitLogger()
	db.InitMysqlDB()
}

// migrate tables.
func main() {
	// get db instance.
	db := db.GetMysqlDB()

	// tables
	tables := []interface{}{
		systemroles.SysUser{},
		systemroles.SysCustomer{},
		systemroles.SysCommunity{},
		systemroles.SysBot{},

		customerroles.CustomerSlack{},
	}

	// drop tables if exist.
	if isDrop {
		for i := 0; i < len(tables); i++ {
			if db.Migrator().HasTable(&tables[i]) && db.Migrator().DropTable(&tables[i]) != nil {
				log.Panicf("OpenKF Table Migration", "Drop table %T... failed", tables[i])
			}
			log.Infof("OpenKF Table Migration", "Drop table %T... ok", tables[i])
		}
	}

	// migrate tables.
	for i := 0; i < len(tables); i++ {
		if db.AutoMigrate(&tables[i]) != nil {
			log.Panicf("OpenKF Table Migration", "Migrate table %T... failed", tables[i])
		}
		log.Infof("OpenKF Table Migration", "Migrate table %T... ok", tables[i])
	}
}
