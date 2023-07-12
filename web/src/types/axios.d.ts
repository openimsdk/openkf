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
