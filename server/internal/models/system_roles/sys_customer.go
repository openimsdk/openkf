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

// SysCustomer system customer model.
type SysCustomer struct {
	base.UserBase

	Device     string `gorm:"column:device;type:varchar(255);not null;default:'';comment:'Customer device'" json:"device"`
	IPAddress  string `gorm:"column:ip_address;type:varchar(255);not null;default:'';comment:'Customer IP address'" json:"ip_address"`
	Source     string `gorm:"column:source;type:varchar(255);not null;default:'';comment:'Visitor source'" json:"source"`
	SourceType int    `gorm:"column:source_type;type:int(11);not null;default:0;comment:'Visitor type, 0:web, 1: slack'" json:"source_type"`
}

// TableName table name.
func (SysCustomer) TableName() string {
	return "sys_customer"
}
