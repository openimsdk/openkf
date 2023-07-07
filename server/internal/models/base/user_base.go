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

package base

import "github.com/gofrs/uuid"

// UserBase user base model.
type UserBase struct {
	Model

	UUID     uuid.UUID `json:"uuid" gorm:"index;type:varchar(36);not null;comment:UUID"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null;unique;comment:Email"`
	Nickname string    `json:"nickname" gorm:"type:varchar(20);not null;comment:Nickname"`
	Avatar   string    `json:"avatar" gorm:"type:varchar(255);not null;comment:Avatar"`
	IsEnable bool      `json:"enable" gorm:"type:tinyint(1);not null;default:1;comment:IsEnable,1:enable,0:disable"`
}
