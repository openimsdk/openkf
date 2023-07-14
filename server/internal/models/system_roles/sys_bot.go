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

package systemroles

import (
	"github.com/OpenIMSDK/OpenKF/server/internal/models/base"
)

// SysBot system bot model.
type SysBot struct {
	base.Model

	BotAddr     string       `gorm:"column:bot_addr;type:varchar(255);not null;comment:'AI bot address'"`
	BotPort     int          `gorm:"column:bot_port;type:int(11);not null;comment:'AI bot port'"`
	BotToken    string       `gorm:"column:bot_token;type:varchar(255);not null;comment:'AI bot token'"`
	Nickname    string       `gorm:"column:nickname;type:varchar(255);not null;comment:'AI bot nickname'"`
	Avatar      string       `gorm:"column:avatar;type:varchar(255);not null;comment:'AI bot avatar'"`
	CommunityId uint         `gorm:"index;column:community_id;type:int(11);not null;comment:'Community id'" json:"community_id"`
	Community   SysCommunity `gorm:"foreignKey:CommunityId;"                                                json:"community"`
}

// TableName table name.
func (SysBot) TableName() string {
	return "sys_bot"
}
