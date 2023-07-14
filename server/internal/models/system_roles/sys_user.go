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

// SysUser system user model.
type SysUser struct {
	base.UserBase

	IsAdmin     bool         `json:"is_admin"     gorm:"index;column:is_admin;type:tinyint(1);not null;default:0;comment:'Is admin'"`
	Password    string       `json:"-"            gorm:"type:varchar(64);not null;comment:Password"`
	CommunityId uint         `json:"community_id" gorm:"column:community_id;type:int(11);not null;comment:'Community id'"`
	Community   SysCommunity `json:"community"    gorm:"foreignKey:CommunityId;"`
}

// TableName table name.
func (SysUser) TableName() string {
	return "sys_user"
}
