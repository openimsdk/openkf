// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"flag"
	"fmt"

	"github.com/OpenIMSDK/OpenKF/server/internal/client"
	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/internal/db"
	"github.com/OpenIMSDK/OpenKF/server/internal/router"
	"github.com/OpenIMSDK/OpenKF/server/internal/utils"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
	"github.com/OpenIMSDK/OpenKF/server/pkg/server"
)

func init() {
	// arg
	configPath := flag.String("c", "./config.yaml", "config file path")
	flag.Parse()

	// init
	config.ConfigInit(*configPath)
	utils.OpenKFBanner()
	log.InitLogger()
	db.InitMysqlDB()
	db.InitRedisDB()
	client.InitMinio()
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
