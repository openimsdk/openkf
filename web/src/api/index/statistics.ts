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

import { GetUserStatisticsParam } from '../request/userModel';
import { GetUserStatisticsResponse } from '../response/userModel';
import { request } from '@/utils/request';

const API = {
    UserStatistics: '/user/statistics',
};

// Get user statistics
export function getUserStatistics(data: GetUserStatisticsParam) {
    return request.post<GetUserStatisticsResponse[]>({
        url: API.UserStatistics,
        params: data,
    });
}