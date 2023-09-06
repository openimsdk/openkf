/**
 * Copyright © 2023 OpenKF & OpenIM open source community. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { CommunityInfo } from './communityModel';

export interface UserInfo {
    avatar: string;
    email: string;
    nickname: string;
    password: string;
}

export interface RegisterAdminParam {
    code: string;
    community_info: CommunityInfo;
    user_info: UserInfo;
}

export interface RegisterStaffParam {
    community_id: number;
    user_info: UserInfo;
}

export interface AccountLoginParam {
    email: string;
    password: string;
}

export interface GetUserInfoParam {
    uuid: string;
}

export interface CreateUserParam {
    user_info: UserInfo;
}

export interface UpdateUserEnableStatusParam {
    uuid: string;
    is_enable: boolean;
}

export interface DeleteUserParam {
    uuid: string;
}

export interface UpdateUserInfoParam {
    avatar: string;
    email: string;
    nickname: string;
    description: string;
}