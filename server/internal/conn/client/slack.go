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

package client

import (
	"github.com/shomali11/slacker"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

// SlackBot slack client.
var SlackBot *slacker.Slacker

// InitSlack init slack client.
func InitSlack() *slacker.Slacker {
	botToken := config.Config.Slack.BotToken
	appToken := config.Config.Slack.AppToken
	debug := config.Config.App.Debug

	SlackBot = slacker.NewClient(botToken, appToken, slacker.WithDebug(debug))
	log.Info("Slack", "init config ok")

	return SlackBot
}
