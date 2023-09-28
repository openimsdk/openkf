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
	// PATH_BOT_ASK ask with bot.
	PATH_BOT_ASK = "/api/v1/ask"
)

// AskBot ask bot.
func AskBot(param *request.BotQuery, operationID, host string) (*response.BotQueryResponse, error) {
	// host: http://ip:port
	url := fmt.Sprintf("%s%s", host, PATH_BOT_ASK)

	r := &response.BotQueryResponse{}
	client := client.NewClient(url)
	resp, err := client.POST(operationID, "", param)
	if err != nil {
		return r, err
	}

	r.Code = uint(resp["code"].(float64))
	r.Msg = resp["msg"].(string)
	if data, ok := resp["data"]; ok {
		data := data.(map[string]interface{})
		r.Data.Context = data["context"].(string)
		r.Data.Question = data["question"].(string)
		r.Data.Text = data["text"].(string)
	}

	return r, nil
}
