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

package constant

// Platform types.
const (
	PLATFORMID_IOS        = 1
	PLATFORMID_ANDROID    = 2
	PLATFORMID_WINDOWS    = 3
	PLATFORMID_OSX        = 4
	PLATFORMID_WEB        = 5
	PLATFORMID_MINIWEB    = 6
	PLATFORMID_LINUX      = 7
	PLATFORMID_ANDROIDPAD = 8
	PLATFORMID_IPAD       = 9
)

// Msg types.
const (
	CONTENT_TYPE_TEXT  = 101
	CONTENT_TYPE_IMAGE = 102
)

// Session types.
const (
	SESSION_TYPE_SINGLE_CHAT                = 1
	SESSION_TYPE_GROUP_CHAT                 = 2
	SESSION_TYPE_LARGE_GROUP_READ_DIFFUSION = 3
	SESSION_TYPE_NOTIFY                     = 4
)
