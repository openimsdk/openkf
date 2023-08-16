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

import { request } from '@/utils/request';
import { GetSlackConfigResponse, GetSlackCustomerResponse } from '@/api/response/platformModel';
import { GetUserInfoParam } from '@/api/request/userModel';


const API = {
    SlackConfig: '/platform/slack/config',
    SlackCustomer: '/platform/slack/customer',
};

// get slack info
export function getSlackConfig() {
    return request.get<GetSlackConfigResponse>({
        url: API.SlackConfig,
    });
}

// get slack customer info
export function getSlackCustomerInfo(data: GetUserInfoParam) {
    return request.post<GetSlackCustomerResponse>({
        url: API.SlackCustomer,
        params: data,
    });
}