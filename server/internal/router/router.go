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

package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// swagger docs.
	_ "github.com/OpenIMSDK/OpenKF/server/docs"

	"github.com/gin-gonic/gin"

	"github.com/OpenIMSDK/OpenKF/server/internal/api"
	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	urltrie "github.com/OpenIMSDK/OpenKF/server/internal/middleware/hooks/url_trie"
)

// InitRouter init router.
func InitRouter() *gin.Engine {
	if config.GetString("app.debug") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	// Set MaxMultipartMemory
	r.MaxMultipartMemory = int64(config.Config.Server.MaxFileSize) << 20

	// Enable Hooks
	r.Use(urltrie.RunHook())

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		// User register api
		register := apiv1.Group("/register")
		{
			// Register with email
			register.POST("/email/code", api.SendCode)
			register.POST("/admin", api.AdminRegister)
			// register.POST("/github", api.GithubRegister)
		}

		// User login api
		login := apiv1.Group("/login")
		{
			login.POST("/account", api.AccountLogin)
			login.POST("/exit", api.AccountExit)
			// user.POST("/email", api.GithubRegister)
			// user.POST("/github", api.GithubRegister)
			// user.POST("/wechat", api.GithubRegister)
		}

		common := apiv1.Group("/common")
		{
			common.POST("/file/upload", api.UploadFile)
		}

		admin := apiv1.Group("/admin")
		{
			staff := admin.Group("/staff")
			{
				staff.POST("/create", api.CreateStaff) // Register a new staff
				staff.POST("/delete", api.DeleteStaff) // Delete a staff
				staff.POST("/update", api.UpdateStaff) // Update staff info
			}
			bot := admin.Group("/bot")
			{
				bot.POST("/create", api.CreateBot)
				bot.POST("/list", api.ListBot)
				bot.POST("/delete", api.DeleteBot)
				bot.POST("/update", api.UpdateBot)
			}
		}

		user := apiv1.Group("/user")
		{
			user.GET("/me", api.GetMyInfo)
			user.POST("/info", api.GetUserInfo)
			user.POST("/userlist", api.GetCommunityUserList)
			user.POST("/update-password", api.UpdatePassword)
			user.POST("/update", api.UpdateInfo)
		}

		community := apiv1.Group("/community")
		{
			community.GET("/me", api.GetMyCommunityInfo)
			community.POST("/info", api.GetCommunityInfo)
			community.POST("/create", api.CreateCommunity)
			community.POST("/update", api.UpdateCommunity)
		}

		// OpenIM callback api
		command := apiv1.Group("/openim/callback")
		{
			command.POST("/", api.OpenIMCallback)
			command.POST("/callbackBeforeSendSingleMsgCommand", api.BeforeSendSingleMsg)
			command.POST("/callbackAfterSendSingleMsgCommand", api.AfterSendSingleMsg)
			command.POST("/callbackMsgModifyCommand", api.MsgModify)
			command.POST("/callbackUserOnlineCommand", api.UserOnline)
			command.POST("/callbackUserOfflineCommand", api.UserOffline)
			command.POST("/callbackOfflinePushCommand", api.OfflinePush)
			command.POST("/callbackOnlinePushCommand", api.OnlinePush)
		}

		// Platform callback api
		platform := apiv1.Group("/platform")
		{
			slack := platform.Group("/slack")
			{
				slack.GET("/config", api.SlackConfig)
			}
		}
	}

	return r
}
