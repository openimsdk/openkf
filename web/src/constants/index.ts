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

// Define OpenIM login config
export const OpenIMLoginConfig = {
    PlatformID: 5, // 5: web
    APIAddress: import.meta.env.VITE_OPENIM_API_ADDRESS,
    WSAddress: import.meta.env.VITE_OPENIM_WS_ADDRESS,
    LogLevel: import.meta.env.VITE_OPENIM_LOG_LEVEL,
};

// Define content type
export enum ContentTypeEnum {
    Json = 'application/json;charset=UTF-8',
    FormURLEncoded = 'application/x-www-form-urlencoded;charset=UTF-8',
    FormData = 'multipart/form-data;charset=UTF-8',
}
