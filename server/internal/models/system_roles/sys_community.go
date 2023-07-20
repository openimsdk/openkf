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
	"github.com/gofrs/uuid"

	"github.com/OpenIMSDK/OpenKF/server/internal/models/base"
)

// SysCommunity system community model.
type SysCommunity struct {
	base.Model

	UUID        uuid.UUID `gorm:"index;column:uuid;column:uuid;not null;unique;comment:'community uuid'"`
	Name        string    `gorm:"column:name;type:varchar(64);not null;comment:'community name'"`
	Email       string    `gorm:"column:email;type:varchar(64);not null;comment:'community email'"`
	Avatar      string    `gorm:"column:avatar;type:varchar(255);not null;comment:'community avatar'"`
	Description string    `gorm:"column:description;type:varchar(255);not null;comment:'community description'"`
}

// TableName table name.
func (SysCommunity) TableName() string {
	return "sys_community"
}
