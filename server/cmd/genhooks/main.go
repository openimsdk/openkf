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
	"os"

	"github.com/OpenIMSDK/OpenKF/server/cmd/genhooks/pkg"
)

var (
	hookName   string
	urlPattern string
	prority    int64
	savePath   string
)

func init() {
	flag.StringVar(&hookName, "name", "", "hook name")
	flag.StringVar(&urlPattern, "pattern", "", "url pattern")
	flag.Int64Var(&prority, "prority", 0, "hook prority")
	flag.StringVar(&savePath, "path", "../../internal/middleware/hooks", "save path")
	flag.Parse()

	if hookName == "" || savePath == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	pkg.NewHookGenerator(hookName, urlPattern, savePath, prority).Generate().Format().Flush()
}
