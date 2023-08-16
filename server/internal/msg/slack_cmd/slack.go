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

package slackcmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shomali11/slacker"

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/client"
	requestparams "github.com/OpenIMSDK/OpenKF/server/internal/params/request"
	"github.com/OpenIMSDK/OpenKF/server/internal/service"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

// SLACK_PERFIX do not use separator.
const SLACK_PERFIX = "SLACK"

// BOT_CONTEXT_MAP bot context map in memory. TODO: use better way to store this map.
var BOT_CONTEXT_MAP map[string]slacker.BotContext

// InitSlackListen init slack socket listen.
func InitSlackListen() {
	bot := client.InitSlack()

	// init bot context map
	BOT_CONTEXT_MAP = make(map[string]slacker.BotContext)
	ctx := context.Background()
	svc := service.NewSlackService(ctx)

	// receive all @ message
	bot.Command("<question>", &slacker.CommandDefinition{
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			BOT_CONTEXT_MAP[bc.Event().ChannelID] = bc

			// find or create a customer
			messageEvent := bc.Event()
			slackUserId := fmt.Sprintf("%s%s", SLACK_PERFIX, messageEvent.UserID)
			senderId, _, err := svc.CreateCustomer(slackUserId, messageEvent.UserProfile)
			if err != nil {
				log.Errorf("Slack", "CreateCustomer failed: %s", err.Error())

				return
			}

			question := r.Param("question")
			// TODO: remove id from question
			words := strings.Split(question, " ")
			question = strings.Join(words[1:], " ")

			// send message to staff and cache pair
			err = svc.SendMsg(senderId, question, bc)
			if err != nil {
				log.Errorf("Slack", "SendMsg failed: %s", err.Error())
				_ = w.Reply("[OpenKF] Error in system.", slacker.WithThreadReply(true))
			}
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := bot.Listen(ctx); err != nil {
		log.Panicf("Slack Client", "Connection failed: %s", err.Error())
	}
}

// SendMsg to slack user.
func SendMsg(params *requestparams.MsgAbstract) {
	recvID := params.RecvID
	content := params.Content
	ctx := context.Background()

	// Parse content {"content": "xxx"}
	type Content struct {
		Content string `json:"content"`
	}
	c := &Content{}
	err := json.Unmarshal([]byte(content), &c)
	if err != nil {
		log.Errorf("Slack", "SendMsg json.Unmarshal failed: %s", err.Error())

		return
	}

	udSvc := service.NewUserDispatchService(ctx)
	sMap := udSvc.GetSlackMap(recvID)
	slackCtx, ok := BOT_CONTEXT_MAP[sMap.SlackChannelID]
	if !ok {
		log.Errorf("Slack", "Slack context not found")

		return
	}

	// Reply in thread and return response
	err = slacker.NewResponse(slackCtx).Reply(c.Content, slacker.WithThreadReply(true))
	if err != nil {
		log.Errorf("Slack", "SendToSlack failed: %s", err.Error())
	}
}
