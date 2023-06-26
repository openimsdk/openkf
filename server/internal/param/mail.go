// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package param

type SendToParams struct {
	Email string `json:"email" binding:"required"`
}
