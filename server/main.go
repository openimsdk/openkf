// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"fmt"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/internal/db"
	"github.com/OpenIMSDK/OpenKF/server/internal/router"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
	"github.com/OpenIMSDK/OpenKF/server/pkg/server"
)

func init() {
	// init config
	config.ConfigInit()

	// logo
	_logo := `
_______                        ______ ____________
__  __ \________ _____ _______ ___  //_/___  ____/
_  / / /___  __ \_  _ \__  __ \__  ,<   __  /_    
/ /_/ / __  /_/ //  __/_  / / /_  /| |  _  __/    
\____/  _  .___/ \___/ /_/ /_/ /_/ |_|  /_/       
        /_/                                  
Copyright © 2023 OpenIMSDK open source community. All rights reserved.
Licensed under the MIT License (the "License");
`
	fmt.Printf("%s\n", _logo)
	fmt.Printf("OpenKF version: %s\n", config.Config.App.Version)

	// init database
	db.InitMysqlDB()
	db.InitRedisDB()
}

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	serverAddress := fmt.Sprintf("%s:%d", config.Config.Server.Ip, config.Config.Server.Port)

	r := router.InitRouter()
	s := server.InitServer(serverAddress, r)
	log.Error("server start error: %v", s.ListenAndServe().Error())
}
