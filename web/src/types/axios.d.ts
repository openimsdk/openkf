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

import type { AxiosRequestConfig } from 'axios';

export interface RequestOptions {
    apiUrl?: string;
    isJoinPrefix?: boolean;
    urlPrefix?: string;
    joinParamsToUrl?: boolean;
    formatDate?: boolean;
    isTransformResponse?: boolean;
    isReturnNativeResponse?: boolean;
    ignoreCancelToken?: boolean;
    joinTime?: boolean;
    withToken?: boolean;
    retry?: {
        count: number;
        delay: number;
    };
    throttle?: {
        delay: number;
    };
    debounce?: {
        delay: number;
    };
}

export interface Result<T = any> {
    code: number;
    msg: string;
    data: T;
}

export interface AxiosRequestConfigRetry extends AxiosRequestConfig {
    retryCount?: number;
}
