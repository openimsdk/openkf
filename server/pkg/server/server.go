// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package server

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type Server interface {
	ListenAndServe() error
}

func InitServer(address string, r *gin.Engine) Server {
	server := endless.NewServer(address, r)

	// set server timeout
	server.ReadHeaderTimeout = 20 * time.Second
	server.WriteTimeout = 20 * time.Second
	server.MaxHeaderBytes = 1 << 20
	return server
}
