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

package customerroles

import "github.com/OpenIMSDK/OpenKF/server/internal/models/base"

// CustomerSlack customer from slack model.
type CustomerSlack struct {
	base.UserBase // UserID in uuid => type+userid. example: slackU05BL9CAN7N

	FirstName             string `json:"first_name"  gorm:"column:first_name;type:varchar(256);comment:'First name'"`
	LastName              string `json:"last_name"  gorm:"column:last_name;type:varchar(256);comment:'Last name'"`
	RealName              string `json:"real_name" gorm:"column:real_name;type:varchar(256);comment:'Real name'"`
	RealNameNormalized    string `json:"real_name_normalized" gorm:"column:real_name_normalized;type:varchar(256);comment:'Real name normalized'"`
	DisplayName           string `json:"display_name" gorm:"column:display_name;type:varchar(256);comment:'Display name'"`
	DisplayNameNormalized string `json:"display_name_normalized" gorm:"column:display_name_normalized;type:varchar(256);comment:'Display name normalized'"`
	Skype                 string `json:"skype" gorm:"column:skype;type:varchar(256);comment:'Skype'"`
	Phone                 string `json:"phone" gorm:"column:phone;type:varchar(256);comment:'Phone'"`
	Image24               string `json:"image_24" gorm:"column:image_24;type:varchar(256);comment:'Image 24'"`
	Image32               string `json:"image_32" gorm:"column:image_32;type:varchar(256);comment:'Image 32'"`
	Image48               string `json:"image_48" gorm:"column:image_48;type:varchar(256);comment:'Image 48'"`
	Image72               string `json:"image_72" gorm:"column:image_72;type:varchar(256);comment:'Image 72'"`
	Image192              string `json:"image_192" gorm:"column:image_192;type:varchar(256);comment:'Image 192'"`
	Image512              string `json:"image_512" gorm:"column:image_512;type:varchar(256);comment:'Image 512'"`
	ImageOriginal         string `json:"image_original" gorm:"column:image_original;type:varchar(256);comment:'Image original'"`
	Title                 string `json:"title" gorm:"column:title;type:varchar(256);comment:'Title'"`
	BotID                 string `json:"bot_id,omitempty" gorm:"column:bot_id;type:varchar(256);comment:'Bot id'"`
	ApiAppID              string `json:"api_app_id,omitempty" gorm:"column:api_app_id;type:varchar(256);comment:'Api app id'"`
	StatusText            string `json:"status_text,omitempty" gorm:"column:status_text;type:varchar(256);comment:'Status text'"`
	StatusEmoji           string `json:"status_emoji,omitempty" gorm:"column:status_emoji;type:varchar(256);comment:'Status emoji'"`
	StatusExpiration      int    `json:"status_expiration" gorm:"column:status_expiration;type:int(11);comment:'Status expiration'"`
	Team                  string `json:"team" gorm:"column:team;type:varchar(256);comment:'Team'"`
}

// TableName table name.
func (CustomerSlack) TableName() string {
	return "customer_slack"
}
