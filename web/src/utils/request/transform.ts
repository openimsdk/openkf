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
