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

package requestparams

// CreateUserParams register params for user.
type CreateUserParams struct {
	Email    string  `json:"email"    binding:"required"`
	Nickname string  `json:"nickname" binding:"required"`
	Avatar   *string `json:"avatar"   binding:"required"` // Avatar is optional.
	Password string  `json:"password" binding:"required"`
}

// RegisterAdminParams register params for admin.
type RegisterAdminParams struct {
	UserInfo      CreateUserParams `json:"user_info"      binding:"required"`
	CommunityInfo CommunityParams  `json:"community_info" binding:"required"`
	Code          string           `json:"code"           binding:"required"`
}

// RegisterStaffParams register params for staff.
type RegisterStaffParams struct {
	UserInfo CreateUserParams `json:"user_info"    binding:"required"`
	// CommunityId uint             `json:"community_id" binding:"required"` // Get community id from token.
}

// LoginParamsWithAccount login params.
type LoginParamsWithAccount struct {
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

// GetUserInfoParams user info params.
type GetUserInfoParams struct {
	UUID string `json:"uuid"        binding:"required"`
}

// UpdateUserInfoParams update user info params.
type UpdateUserInfoParams struct {
	// Email can not be updated is this period.
	Email       *string `json:"email"`       // Email is optional.
	Nickname    *string `json:"nickname"`    // Nickname is optional.
	Description *string `json:"description"` // Description is optional.
	Avatar      *string `json:"avatar"`      // Avatar is optional.
}

// UpdateUserPasswordParams update user password params.
type UpdateUserPasswordParams struct {
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeat_password" binding:"required"`
}

// DeleteUserParams delete user params.
type DeleteUserParams struct {
	UUID string `json:"uuid" binding:"required"`
}
