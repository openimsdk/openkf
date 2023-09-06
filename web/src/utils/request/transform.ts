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

import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import { AxiosError } from 'axios';

import type { RequestOptions, Result } from '@/types/axios';

export interface CreateAxiosOptions extends AxiosRequestConfig {
    authenticationScheme?: string;
    transform?: AxiosTransform;
    requestOptions?: RequestOptions;
}

export abstract class AxiosTransform {
    beforeRequestHook?: (
        config: AxiosRequestConfig,
        options: RequestOptions,
    ) => AxiosRequestConfig;

    transformRequestHook?: <T = any>(
        res: AxiosResponse<Result>,
        options: RequestOptions,
    ) => T;

    requestCatchHook?: <T = any>(
        e: Error | AxiosError,
        options: RequestOptions,
    ) => Promise<T>;

    requestInterceptors?: (
        config: AxiosRequestConfig,
        options: CreateAxiosOptions,
    ) => AxiosRequestConfig;

    responseInterceptors?: (res: AxiosResponse) => AxiosResponse;

    requestInterceptorsCatch?: (error: AxiosError) => void;

    responseInterceptorsCatch?: (
        error: AxiosError,
        instance: AxiosInstance,
    ) => void;
}
