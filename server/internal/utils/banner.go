// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

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
🎉🎉 Welcome to your contribution :)

Copyright © 2023 OpenIMSDK open source community. All rights reserved.
Licensed under the MIT License (the "License");
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
