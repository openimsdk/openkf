// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package models

type User struct {
	Model
	Username string `json:"username" gorm:"type:varchar(20);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(20);not null"`
}
