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

import { 
    CreateUserParam, 
    UpdateUserEnableStatusParam,
    DeleteUserParam } from '../request/userModel';
import { GetUserInfoResponse, GetUserInfoListResponse } from '../response/userModel';
import { request } from '@/utils/request';

const API = {
    StaffCreate: '/admin/staff/create',
    StaffDelete: '/admin/staff/delete',
    StaffUpdate: '/admin/staff/update',
};

// Create user
export function createStaff(data: CreateUserParam) {
    return request.post<any>({
        url: API.StaffCreate,
        params: data,
    });
}

// Update user enable status
export function updateStaffEnableStatus(data: UpdateUserEnableStatusParam) {
    return request.post<any>({
        url: API.StaffUpdate,
        params: data,
    });
}

// Delete user
export function deleteStaff(data: DeleteUserParam) {
    return request.post<any>({
        url: API.StaffDelete,
        params: data,
    });
}