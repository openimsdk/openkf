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

import { get } from 'lodash';
import { PageParam, PAGE } from '../request/commonModel';
import { GetUserInfoParam, UpdateUserInfoParam } from '../request/userModel';
import { GetUserInfoResponse, GetUserInfoListResponse } from '../response/userModel';
import { request } from '@/utils/request';

const API = {
    UserMe: '/user/me',
    UserInfo: '/user/info',
    UserList: '/user/userlist',
    UserUpdate: '/user/update',
};

// Get my info
export function getMyInfo() {
    return request.get<GetUserInfoResponse>({
        url: API.UserMe,
    });
}

// Get user info
export function getUserInfo(data: GetUserInfoParam) {
    return request.post<GetUserInfoResponse>({
        url: API.UserInfo,
        params: data,
    });
}

// Get user list info with default page
export function getUserInfoListDefault(pageNum = 1) {
    const page:PageParam = PAGE;
    page.page = pageNum;
    return getUserInfoList(PAGE);
}

// Get user list info
export function getUserInfoList(data: PageParam) {
    return request.post<GetUserInfoListResponse>({
        url: API.UserList,
        params: data,
    });
}

// Update user info
export function updateUserInfo(data: UpdateUserInfoParam) {
    return request.post<any>({
        url: API.UserUpdate,
        params: data,
    });
}