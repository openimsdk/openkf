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

package utils

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
)

func OpenKFBanner() {
	// logo
	logo := `
_______                        ______ ____________
__  __ \________ _____ _______ ___  //_/___  ____/
_  / / /___  __ \_  _ \__  __ \__  ,<   __  /_    
/ /_/ / __  /_/ //  __/_  / / /_  /| |  _  __/    
\____/  _  .___/ \___/ /_/ /_/ /_/ |_|  /_/       
	/_/                                  

APP Mode:
- Version: %s
- Debug: %v
- Log file: %s

Github repo: https://github.com/OpenIMSDK/OpenKF. 
Official website: https://https://www.openim.online/en
OpenKF Slack: https://join.slack.com/t/openimsdk/shared_invite/zt-1tmoj26uf-_FDy3dowVHBiGvLk9e5Xkg (OpenKF clannels)
🎉🎉 Welcome to your contribution :)

Copyright © 2023 OpenIM open source community. All rights reserved.
Licensed under the Apache License (the "License");
`
	content := fmt.Sprintf(logo, config.Config.App.Version, config.Config.App.Debug, config.Config.App.LogFile)

	Box := box.New(box.Config{
		Px:       5,
		Py:       0,
		Type:     "Round",
		Color:    "Cyan",
		TitlePos: "Top",
	})
	Box.Print("OpenIM Community", content)
}
