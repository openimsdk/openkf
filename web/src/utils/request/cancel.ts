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
