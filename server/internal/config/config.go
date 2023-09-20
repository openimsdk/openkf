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

package config

// Config global config instance.
var Config *config

// ConfigInit init config.
func ConfigInit(configPath string) {
	// init viper
	initViper(configPath)

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
			Ip:          GetString("server.ip"),
			Port:        GetInt("server.port"),
			MaxFileSize: GetIntOrDefault("server.max_file_size", 10),
		},
		Mysql: Mysql{
			Ip:           GetString("mysql.ip"),
			Port:         GetInt("mysql.port"),
			Username:     GetString("mysql.username"),
			Password:     GetString("mysql.password"),
			Database:     GetString("mysql.database"),
			MaxLifetime:  GetIntOrDefault("mysql.max_lifetime", 120),
			MaxOpenConns: GetIntOrDefault("mysql.max_open_conns", 100),
			MaxIdleConns: GetIntOrDefault("mysql.max_idle_conns", 20),
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
		InfluxDB: InfluxDB{
			Ip:     GetString("influxdb.ip"),
			Port:   GetInt("influxdb.port"),
			Token:  GetString("influxdb.token"),
			Org:    GetString("influxdb.org"),
			Bucket: GetString("influxdb.bucket"),
		},
		Email: Email{
			Host:     GetString("email.host"),
			Port:     GetInt("email.port"),
			From:     GetString("email.from"),
			Nickname: GetString("email.nickname"),
			Password: GetString("email.password"),
			PoolSize: GetIntOrDefault("email.pool_size", 4),
		},
		OpenIM: OpenIM{
			Secret:     GetStringOrDefault("openim.secret", "openkf"),
			Ip:         GetStringOrDefault("openim.ip", "127.0.0.1"),
			ApiPort:    GetIntOrDefault("openim.api_port", 10002),
			PlatformID: GetIntOrDefault("openim.platform_id", 5),
			AdminID:    GetStringOrDefault("openim.admin_id", "openkf_admin"),
		},
		Slack: Slack{
			BotToken:          GetString("slack.bot_token"),
			AppToken:          GetString("slack.app_token"),
			AppID:             GetString("slack.app_id"),
			ClientID:          GetString("slack.client_id"),
			ClientSecret:      GetString("slack.client_secret"),
			SigningSecret:     GetString("slack.signing_secret"),
			VerificationToken: GetString("slack.verification_token"),
		},
	}
}

type config struct {
	App      App
	JWT      JWT
	Server   Server
	Mysql    Mysql
	Redis    Redis
	Minio    Minio
	InfluxDB InfluxDB
	Email    Email
	OpenIM   OpenIM
	Slack    Slack
}

// App config.
type App struct {
	Version string `mapstructure:"version"`
	Debug   bool   `mapstructure:"debug"`
	LogFile string `mapstructure:"log_file"`
}

// JWT config.
type JWT struct {
	Secret     string `mapstructure:"secret"`
	Issuer     string `mapstructure:"issuer"`
	ExpireDays int    `mapstructure:"expire_days"`
}

// Server config.
type Server struct {
	Ip          string `mapstructure:"ip"`
	Port        int    `mapstructure:"port"`
	MaxFileSize int    `mapstructure:"max_file_size"`
}

// Mysql config.
type Mysql struct {
	Ip           string `mapstructure:"ip"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	MaxLifetime  int    `mapstructure:"max_lifetime"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// Redis config.
type Redis struct {
	Ip       string `mapstructure:"ip"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

// Minio config.
type Minio struct {
	Ip              string `mapstructure:"ip"`
	Port            int    `mapstructure:"port"`
	AccessKeyId     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	Bucket          string `mapstructure:"bucket"`
	AppBucket       string `mapstructure:"app_bucket"`
	Location        string `mapstructure:"location"`
}

// InfluxDB config.
type InfluxDB struct {
	Ip     string `mapstructure:"ip"`
	Port   int    `mapstructure:"port"`
	Token  string `mapstructure:"token"`
	Org    string `mapstructure:"org"`
	Bucket string `mapstructure:"bucket"`
}

// Email config.
type Email struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	From     string `mapstructure:"from"`
	Nickname string `mapstructure:"nickname"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

// OpenIM config.
type OpenIM struct {
	Secret     string `mapstructure:"secret"`
	Ip         string `mapstructure:"ip"`
	ApiPort    int    `mapstructure:"api_port"`
	PlatformID int    `mapstructure:"platform_id"`
	AdminID    string `mapstructure:"admin_id"`
}

// Slack Bot config.
type Slack struct {
	BotToken          string `mapstructure:"bot_token"`
	AppToken          string `mapstructure:"app_token"`
	AppID             string `mapstructure:"app_id"`
	ClientID          string `mapstructure:"client_id"`
	ClientSecret      string `mapstructure:"client_secret"`
	SigningSecret     string `mapstructure:"signing_secret"`
	VerificationToken string `mapstructure:"verification_token"`
}
