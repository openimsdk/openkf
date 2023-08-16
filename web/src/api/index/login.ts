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
    AccountLoginParam,
    RegisterAdminParam,
    RegisterStaffParam,
} from '../request/userModel';
import { UserLoginResponse } from '../response/userModel';
import { request } from '@/utils/request';

const API = {
    EmailCode: '/register/email/code',
    RegisterAdmin: '/register/admin',
    RegisterStaff: '/register/staff',
    AccountLogin: '/login/account',
    AccountExit: '/login/exit',
};

// Send email verification code
export function sendEmailCode(email: string) {
    return request.post<any>({
        url: API.EmailCode,
        data: {
            email,
        },
    });
}

// Register admin
export function registerAdmin(data: RegisterAdminParam) {
    return request.post<any>({
        url: API.RegisterAdmin,
        data,
    });
}

// Register admin
export function registerStaff(data: RegisterStaffParam) {
    return request.post<any>({
        url: API.RegisterStaff,
        data,
    });
}

// User login
export function accountLogin(data: AccountLoginParam) {
    return request.post<UserLoginResponse>({
        url: API.AccountLogin,
        data,
    });
}

// User logout
export function accountLogout() {
    return request.post<any>({
        url: API.AccountExit,
    });
}