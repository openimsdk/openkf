import axios, {
    AxiosError,
    AxiosInstance,
    AxiosRequestConfig,
    AxiosRequestHeaders,
    AxiosResponse,
    InternalAxiosRequestConfig,
} from 'axios';
import cloneDeep from 'lodash/cloneDeep';
import debounce from 'lodash/debounce';
import isFunction from 'lodash/isFunction';
import throttle from 'lodash/throttle';
import { stringify } from 'qs';

import { ContentTypeEnum } from '@/constants';
import { AxiosRequestConfigRetry, RequestOptions, Result } from '@/types/axios';

import { AxiosCanceler } from './cancel';
import { CreateAxiosOptions } from './transform';

export class VAxios {
    private instance: AxiosInstance;
    private readonly options: CreateAxiosOptions;

    constructor(options: CreateAxiosOptions) {
        this.options = options;
        this.instance = axios.create(options);
        this.setupInterceptors();
    }

    private createAxios(config: CreateAxiosOptions): void {
        this.instance = axios.create(config);
    }

    private getTransform() {
        const { transform } = this.options;
        return transform;
    }

    getAxios(): AxiosInstance {
        return this.instance;
    }

    configAxios(config: CreateAxiosOptions) {
        if (!this.instance) return;
        this.createAxios(config);
    }

    setHeader(headers: Record<string, string>): void {
        if (!this.instance) return;
        Object.assign(this.instance.defaults.headers, headers);
    }

    private setupInterceptors() {
        const transform = this.getTransform();
        if (!transform) return;

        const {
            requestInterceptors,
            requestInterceptorsCatch,
            responseInterceptors,
            responseInterceptorsCatch,
        } = transform;
        const axiosCanceler = new AxiosCanceler();

        this.instance.interceptors.request.use(
            (config: InternalAxiosRequestConfig) => {
                // @ts-ignore
                const { ignoreCancelToken } = config.requestOptions;
                const ignoreCancel =
                    ignoreCancelToken ??
                    this.options.requestOptions?.ignoreCancelToken;
                if (!ignoreCancel) axiosCanceler.addPending(config);

                if (requestInterceptors && isFunction(requestInterceptors)) {
                    config = requestInterceptors(
                        config,
                        this.options,
                    ) as InternalAxiosRequestConfig;
                }

                return config;
            },
            undefined,
        );

        if (requestInterceptorsCatch && isFunction(requestInterceptorsCatch)) {
            this.instance.interceptors.request.use(
                undefined,
                requestInterceptorsCatch,
            );
        }

        this.instance.interceptors.response.use((res: AxiosResponse) => {
            if (res) axiosCanceler.removePending(res.config);
            if (responseInterceptors && isFunction(responseInterceptors)) {
                res = responseInterceptors(res);
            }
            return res;
        }, undefined);

        if (
            responseInterceptorsCatch &&
            isFunction(responseInterceptorsCatch)
        ) {
            this.instance.interceptors.response.use(undefined, error =>
                responseInterceptorsCatch(error, this.instance),
            );
        }
    }

    supportFormData(config: AxiosRequestConfig) {
        const headers =
            config.headers || (this.options.headers as AxiosRequestHeaders);
        const contentType =
            headers?.['Content-Type'] || headers?.['content-type'];

        if (
            contentType !== ContentTypeEnum.FormURLEncoded ||
            !Reflect.has(config, 'data') ||
            config.method?.toUpperCase() === 'GET'
        ) {
            return config;
        }

        return {
            ...config,
            data: stringify(config.data, { arrayFormat: 'brackets' }),
        };
    }

    supportParamsStringify(config: AxiosRequestConfig) {
        const headers = config.headers || this.options.headers;
        const contentType =
            headers?.['Content-Type'] || headers?.['content-type'];

        if (
            contentType === ContentTypeEnum.FormURLEncoded ||
            !Reflect.has(config, 'params')
        ) {
            return config;
        }

        return {
            ...config,
            paramsSerializer: (params: any) =>
                stringify(params, { arrayFormat: 'brackets' }),
        };
    }

    get<T = any>(
        config: AxiosRequestConfig,
        options?: RequestOptions,
    ): Promise<T> {
        return this.request({ ...config, method: 'GET' }, options);
    }

    post<T = any>(
        config: AxiosRequestConfig,
        options?: RequestOptions,
    ): Promise<T> {
        return this.request({ ...config, method: 'POST' }, options);
    }

    put<T = any>(
        config: AxiosRequestConfig,
        options?: RequestOptions,
    ): Promise<T> {
        return this.request({ ...config, method: 'PUT' }, options);
    }

    delete<T = any>(
        config: AxiosRequestConfig,
        options?: RequestOptions,
    ): Promise<T> {
        return this.request({ ...config, method: 'DELETE' }, options);
    }

    patch<T = any>(
        config: AxiosRequestConfig,
        options?: RequestOptions,
    ): Promise<T> {
        return this.request({ ...config, method: 'PATCH' }, options);
    }

    upload<T = any>(
        key: string,
        file: File,
        config: AxiosRequestConfig,
        options?: RequestOptions,
    ): Promise<T> {
        const params: FormData = config.params ?? new FormData();
        params.append(key, file);

        return this.request(
            {
                ...config,
                method: 'POST',
                headers: {
                    'Content-Type': ContentTypeEnum.FormData,
                },
                params,
            },
            options,
        );
    }

    request<T = any>(
        config: AxiosRequestConfigRetry,
        options?: RequestOptions,
    ): Promise<T> {
        const { requestOptions } = this.options;

        if (
            requestOptions!.throttle !== undefined &&
            requestOptions!.debounce !== undefined
        ) {
            throw new Error(
                'throttle and debounce cannot be set at the same time',
            );
        }

        if (requestOptions!.throttle && requestOptions!.throttle.delay !== 0) {
            return new Promise(resolve => {
                throttle(
                    () => resolve(this.synthesisRequest(config, options)),
                    requestOptions!.throttle!.delay,
                );
            });
        }

        if (requestOptions!.debounce && requestOptions!.debounce.delay !== 0) {
            return new Promise(resolve => {
                debounce(
                    () => resolve(this.synthesisRequest(config, options)),
                    requestOptions!.debounce!.delay,
                );
            });
        }

        return this.synthesisRequest(config, options);
    }

    private async synthesisRequest<T = any>(
        config: AxiosRequestConfigRetry,
        options?: RequestOptions,
    ): Promise<T> {
        let conf: CreateAxiosOptions = cloneDeep(config);
        const transform = this.getTransform();

        const { requestOptions } = this.options;

        const opt: RequestOptions = { ...requestOptions, ...options };

        const { beforeRequestHook, requestCatchHook, transformRequestHook } =
            transform || {};
        if (beforeRequestHook && isFunction(beforeRequestHook)) {
            conf = beforeRequestHook(conf, opt);
        }
        conf.requestOptions = opt;

        conf = this.supportFormData(conf);

        return new Promise((resolve, reject) => {
            this.instance
                .request<any, AxiosResponse<Result>>(
                    !config.retryCount ? conf : config,
                )
                .then((res: AxiosResponse<Result>) => {
                    if (
                        transformRequestHook &&
                        isFunction(transformRequestHook)
                    ) {
                        try {
                            const ret = transformRequestHook(res, opt);
                            resolve(ret);
                        } catch (err) {
                            reject(err || new Error('Request error!'));
                        }
                        return;
                    }
                    resolve(res as unknown as Promise<T>);
                })
                .catch((e: Error | AxiosError) => {
                    if (requestCatchHook && isFunction(requestCatchHook)) {
                        reject(requestCatchHook(e, opt));
                        return;
                    }
                    if (axios.isAxiosError(e)) {
                        // Add axios error code
                    }
                    reject(e);
                });
        });
    }
}
