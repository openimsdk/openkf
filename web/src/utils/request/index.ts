import type { AxiosInstance } from 'axios';
import isString from 'lodash/isString';
import merge from 'lodash/merge';

import { ContentTypeEnum } from '@/constants';

import { VAxios } from './axios';
import type { AxiosTransform, CreateAxiosOptions } from './transform';
import { formatRequestDate, joinTimestamp, setObjToUrlParams } from './utils';
import useUserStore from '@/store/user';

// Default API url
const host = import.meta.env.VITE_API_URL;

const transform: AxiosTransform = {
    transformRequestHook: (res, options) => {
        const { isTransformResponse, isReturnNativeResponse } = options;

        // Return data without modification by default
        const method = res.config.method?.toLowerCase();
        if (
            res.status === 204 &&
            ['put', 'patch', 'delete'].includes(method ?? '')
        ) {
            return res;
        }

        if (isReturnNativeResponse) {
            return res;
        }

        if (!isTransformResponse) {
            return res.data;
        }

        const { data } = res;
        if (!data) {
            throw new Error('Response data does not exist');
        }

        const { code } = data;

        // Set success data to 200
        const hasSuccess = data && code === 200;
        if (hasSuccess) {
            return data.data;
        }

        throw new Error(`Request error, error code: ${code}`);
    },

    beforeRequestHook: (config, options) => {
        const {
            apiUrl,
            isJoinPrefix,
            urlPrefix,
            joinParamsToUrl,
            formatDate,
            joinTime = true,
        } = options;

        if (isJoinPrefix && urlPrefix && isString(urlPrefix)) {
            config.url = `${urlPrefix}${config.url}`;
        }

        if (apiUrl && isString(apiUrl)) {
            config.url = `${apiUrl}${config.url}`;
        }
        const params = config.params || {};
        const data = config.data || false;

        if (formatDate && data && !isString(data)) {
            formatRequestDate(data);
        }
        if (config.method?.toUpperCase() === 'GET') {
            if (!isString(params)) {
                config.params = Object.assign(
                    params || {},
                    joinTimestamp(joinTime, false),
                );
            } else {
                config.url = `${config.url + params}${joinTimestamp(
                    joinTime,
                    true,
                )}`;
                config.params = undefined;
            }
        } else if (!isString(params)) {
            if (formatDate) {
                formatRequestDate(params);
            }
            if (
                Reflect.has(config, 'data') &&
                config.data &&
                (Object.keys(config.data).length > 0 ||
                    data instanceof FormData)
            ) {
                config.data = data;
                config.params = params;
            } else {
                config.data = params;
                config.params = undefined;
            }
            if (joinParamsToUrl) {
                config.url = setObjToUrlParams(config.url as string, {
                    ...config.params,
                    ...config.data,
                });
            }
        } else {
            config.url += params;
            config.params = undefined;
        }
        return config;
    },

    requestInterceptors: (config, options) => {
        const userStore = useUserStore();
        const kf_token = userStore.kf_token.token;

        if (
            kf_token &&
            (config as Recordable)?.requestOptions?.withToken !== false
        ) {
            // Append Bearer token
            (config as Recordable).headers.Authorization =
                options.authenticationScheme
                    ? `${options.authenticationScheme} ${kf_token}`
                    : kf_token;
        }

        return config;
    },

    responseInterceptors: res => {
        return res;
    },

    responseInterceptorsCatch: (error: any, instance: AxiosInstance) => {
        const { config } = error;
        if (!config || !config.requestOptions.retry)
            return Promise.reject(error);

        config.retryCount = config.retryCount || 0;

        if (config.retryCount >= config.requestOptions.retry.count)
            return Promise.reject(error);

        config.retryCount += 1;

        const backoff = new Promise(resolve => {
            setTimeout(() => {
                resolve(config);
            }, config.requestOptions.retry.delay || 1);
        });
        config.headers = {
            ...config.headers,
            'Content-Type': ContentTypeEnum.Json,
        };
        return backoff.then(config => instance.request(config as any));
    },
};

function createAxios(opt?: Partial<CreateAxiosOptions>) {
    return new VAxios(
        merge(
            <CreateAxiosOptions>{
                // authenticationScheme: '',
                authenticationScheme: 'Bearer', // Use Bearer type token
                timeout: 10 * 1000,
                withCredentials: true,
                headers: { 'Content-Type': ContentTypeEnum.Json },
                transform,
                requestOptions: {
                    apiUrl: host,
                    isJoinPrefix: true,
                    urlPrefix: import.meta.env.VITE_API_URL_PREFIX,
                    isReturnNativeResponse: false,
                    isTransformResponse: true,
                    joinParamsToUrl: false,
                    formatDate: true,
                    joinTime: true,
                    ignoreCancelToken: true,
                    withToken: true,
                    retry: {
                        count: 3,
                        delay: 1000,
                    },
                },
            },
            opt || {},
        ),
    );
}
export const request = createAxios();
