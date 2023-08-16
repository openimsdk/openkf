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

// Platform type
export enum PlatformType {
    Slack = 'SLACK',
    Web = 'WEB',
}

// Define support platforms
export const PLATFORMS = [
    {
        order: 1,
        avatar: 'https://github.com/OpenIMSDK/OpenKF/assets/47499836/73b94766-9968-4b66-b0b6-cc7a6ebfda69',
        name: PlatformType.Slack,
        description: 'Slack is a new way to communicate with your customer. It\'s faster, better organized, and more secure than email.',
        is_enable: true,
        tags: ['Slack', 'LLM'], 
    },
    {
        order: 2,
        avatar: 'https://github.com/OpenIMSDK/OpenKF/assets/47499836/13292e53-68df-46f8-948c-1296ba3bf330',
        name: PlatformType.Web,
        description: 'Web is a basic way to communicate with your customer. It\'s light and easy to intergrate with your products.',
        is_enable: false,
        tags: ['Web', 'AI'],
    }
];