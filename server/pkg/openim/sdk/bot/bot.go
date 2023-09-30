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

package bot

import (
	"fmt"

	"github.com/openimsdk/openkf/server/pkg/openim/client"
	"github.com/openimsdk/openkf/server/pkg/openim/param/request"
	"github.com/openimsdk/openkf/server/pkg/openim/param/response"
)

const (
	// pathBotTask ask with bot.
	pathBotTask = "/api/v1/ask"
)

// AskBot ask bot.
func AskBot(param *request.BotQuery, operationID, host string) (*response.BotQueryResponse, error) {
	// host: http://ip:port
	url := fmt.Sprintf("%s%s", host, pathBotTask)

	r := &response.BotQueryResponse{}
	client := client.NewClient(url)
	resp, err := client.POST(operationID, "", param)
	if err != nil {
		return r, err
	}

	code, ok := resp["code"].(float64)
	if !ok {
		return r, fmt.Errorf("code is not float64")
	}
	r.Code = uint(code)

	r.Msg, ok = resp["msg"].(string)
	if !ok {
		return r, fmt.Errorf("msg is not string")
	}

	data, ok := resp["data"]
	if !ok {
		return r, fmt.Errorf("data is not exist")
	}

	newData, ok := data.(map[string]interface{})
	if !ok {
		return r, fmt.Errorf("data is not map[string]interface{}")
	}

	r.Data.Context, ok = newData["context"].(string)
	if !ok {
		return r, fmt.Errorf("context is not string")
	}

	r.Data.Question, ok = newData["question"].(string)
	if !ok {
		return r, fmt.Errorf("question is not string")
	}

	r.Data.Text, ok = newData["text"].(string)
	if !ok {
		return r, fmt.Errorf("text is not string")
	}

	return r, nil
}
