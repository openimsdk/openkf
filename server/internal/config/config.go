// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package config

var Config *config

func ConfigInit() {
	// init Configuration
	Config = &config{
		App: App{
			Version: GetString("app.version"),
			Debug:   GetBool("app.debug"),
			LogFile: GetString("app.log_file"),
		},
		JWT: JWT{
			Secret:     GetString("jwt.secret"),
			Issuer:     GetString("jwt.issuer"),
			ExpireDays: GetInt("jwt.expire_days"),
		},
		Server: Server{
			Ip:   GetString("server.ip"),
			Port: GetInt("server.port"),
		},
		Mysql: Mysql{
			Ip:       GetString("mysql.ip"),
			Port:     GetInt("mysql.port"),
			Username: GetString("mysql.username"),
			Password: GetString("mysql.password"),
			Database: GetString("mysql.database"),
		},
		Redis: Redis{
			Ip:       GetString("redis.ip"),
			Port:     GetInt("redis.port"),
			Password: GetString("redis.password"),
			Database: GetInt("redis.database"),
		},
		Minio: Minio{
			Ip:              GetString("minio.ip"),
			Port:            GetInt("minio.port"),
			AccessKeyId:     GetString("minio.access_key_id"),
			SecretAccessKey: GetString("minio.secret_access_key"),
			Bucket:          GetString("minio.bucket"),
			AppBucket:       GetString("minio.app_bucket"),
			Location:        GetString("minio.location"),
		},
	}
}

type config struct {
	App    App
	JWT    JWT
	Server Server
	Mysql  Mysql
	Redis  Redis
	Minio  Minio
}

type App struct {
	Version string
	Debug   bool
	LogFile string
}

type JWT struct {
	Secret     string
	Issuer     string
	ExpireDays int
}

type Server struct {
	Ip   string
	Port int
}

type Mysql struct {
	Ip       string
	Port     int
	Username string
	Password string
	Database string
}

type Redis struct {
	Ip       string
	Port     int
	Password string
	Database int
}

type Minio struct {
	Ip              string
	Port            int
	AccessKeyId     string
	SecretAccessKey string
	Bucket          string
	AppBucket       string
	Location        string
}
