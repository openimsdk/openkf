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

package common

import "github.com/pkg/errors"

// msg is a mapping of message.
var msg = map[int]string{
	SUCCESS:        "success",
	ERROR:          "error",
	INVALID_PARAMS: "request params error",
	UNAUTHORIZED:   "unauthorized",
	FORBIDDEN:      "forbidden",

	// OpenIM callback code
	OPENIM_SERVER_ALLOW_ACTION: "OpenIM allow action",
	OPENIM_SERVER_DENY_ACTION:  "OpenIM deny action",

	// KF service status
	KF_RECORD_NOT_FOUND: "kf: record not found",
	KF_FILE_SIZE_LIMIT:  "kf: file size limit",
	KF_INTERNAL_ERROR:   "kf: internal error",
}

// internalMsg is a mapping of internal service message.
var internalMsg = map[int]string{
	I_INVALID_PARAM: "invalid param",
}

// GetMsg get the message by code.
func GetMsg(code int) string {
	m, ok := msg[code]
	if ok {
		return m
	}

	return msg[ERROR]
}

// NewError return a new error.
func NewError(code int) error {
	return errors.Errorf("%d: %s", code, internalMsg[code])
}
