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

import { CreateCommunityParam, UpdateCommunityParam } from '../request/communityModel';
import { GetCommunityInfoResponse } from '../response/communityModel';
import { request } from '@/utils/request';

const API = {
    CreateCommunity: '/community/create',
    CommunityMe: '/community/me',
    CommunityUpdate: '/community/update',
};

// Create community
export function createCommunity(data: CreateCommunityParam) {
    return request.post<CreateCommunityParam>({
        url: API.CreateCommunity,
        data,
    });
}

// Get my community info
export function getMyCommunityInfo() {
    return request.get<GetCommunityInfoResponse>({
        url: API.CommunityMe,
    });
}

// Get community info
export function getCommunityInfo(data: GetCommunityInfoResponse) {
    return request.get<GetCommunityInfoResponse>({
        url: API.CommunityMe,
        params: data,
    });
}

// Update community info
export function updateCommunityInfo(data: UpdateCommunityParam) {
    return request.post<any>({
        url: API.CommunityUpdate,
        params: data,
    });
}