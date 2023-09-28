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

package bot_test

import (
	"testing"

	"github.com/openimsdk/openkf/server/pkg/openim/param/request"
	"github.com/openimsdk/openkf/server/pkg/openim/sdk/bot"
)

// TestAskBot test ask bot info
func TestAskBot(t *testing.T) {
	// test case
	testData := []struct {
		query string
	}{
		{
			query: "What is openim",
		},
	}

	// range test case
	for _, data := range testData {
		resp, err := bot.AskBot(&request.BotQuery{
			Query: data.query,
		},
			"123123123123",
			"http://localhost:10011")
		if err != nil {
			t.Error(err)
		}
		t.Error(resp)
	}
}
