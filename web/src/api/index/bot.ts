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
    CreateBotParam, 
    UpdateBotParam,
    DeleteBotParam } from '../request/botModel';
import { PageParam, PAGE } from '../request/commonModel';
import { GetBotInfoResponse, GetBotInfoListResponse } from '../response/botModel';
import { request } from '@/utils/request';

const API = {
    BotCreate: '/admin/bot/create',
    BotDelete: '/admin/bot/delete',
    BotUpdate: '/admin/bot/update',
    BotList: '/admin/bot/list',
};

// Create bot
export function createBot(data: CreateBotParam) {
    return request.post<any>({
        url: API.BotCreate,
        params: data,
    });
}

// Update bot
export function updateBot(data: UpdateBotParam) {
    return request.post<any>({
        url: API.BotUpdate,
        params: data,
    });
}

// Delete bot
export function deleteBot(data: DeleteBotParam) {
    return request.post<any>({
        url: API.BotDelete,
        params: data,
    });
}

// Get bot list info with default page
export function getBotInfoListDefault(pageNum = 1) {
    const page:PageParam = PAGE;
    page.page = pageNum;
    return getBotList(PAGE);
}

// Get bot list
export function getBotList(data: PageParam) {
    return request.post<GetBotInfoListResponse>({
        url: API.BotList,
        params: data,
    });
}

    