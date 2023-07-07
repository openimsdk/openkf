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

	"gorm.io/gen"

	systemroles "github.com/OpenIMSDK/OpenKF/server/internal/models/system_roles"
)

func main() {
	outpath := flag.String("path", "../../internal/dal/gen", "output path for generated dal files")
	flag.Parse()

	g := gen.NewGenerator(gen.Config{
		OutPath:       *outpath,
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// Add tables here to generate
	tables := []interface{}{
		systemroles.SysUser{},
		systemroles.SysCustomer{},
		systemroles.SysCommunity{},
		systemroles.SysBot{},
	}

	// Generate basic dao
	g.ApplyBasic(tables...)

	// Generate query interface with dynamic query
	// Ref: https://gorm.io/gen/
	// g.ApplyInterface()

	g.Execute()
}
