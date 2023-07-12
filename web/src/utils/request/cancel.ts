import type { AxiosRequestConfig, Canceler } from 'axios';
import axios from 'axios';
import isFunction from 'lodash/isFunction';

let pendingMap = new Map<string, Canceler>();

export const getPendingUrl = (config: AxiosRequestConfig) =>
    [config.method, config.url].join('&');

export class AxiosCanceler {
    addPending(config: AxiosRequestConfig) {
        this.removePending(config);
        const url = getPendingUrl(config);
        config.cancelToken =
            config.cancelToken ||
            new axios.CancelToken(cancel => {
                if (!pendingMap.has(url)) {
                    pendingMap.set(url, cancel);
                }
            });
    }

    removeAllPending() {
        pendingMap.forEach(cancel => {
            if (cancel && isFunction(cancel)) cancel();
        });
        pendingMap.clear();
    }

    removePending(config: AxiosRequestConfig) {
        const url = getPendingUrl(config);

        if (pendingMap.has(url)) {
            // If there is a current request identifier in pending,
            // the current request needs to be cancelled and removed
            const cancel = pendingMap.get(url);
            if (cancel) cancel(url);
            pendingMap.delete(url);
        }
    }

    reset() {
        pendingMap = new Map<string, Canceler>();
    }
}
