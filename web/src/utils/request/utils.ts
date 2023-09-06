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

import isObject from 'lodash/isObject';
import isString from 'lodash/isString';

const DATE_TIME_FORMAT = 'YYYY-MM-DD HH:mm:ss';

// join timestamp
export function joinTimestamp<T extends boolean>(
    join: boolean,
    restful: T,
): T extends true ? string : object;

export function joinTimestamp(join: boolean, restful = false): string | object {
    if (!join) {
        return restful ? '' : {};
    }
    const now = new Date().getTime();
    if (restful) {
        return `?_t=${now}`;
    }
    return { _t: now };
}

// format request params
export function formatRequestDate(params: Recordable) {
    if (Object.prototype.toString.call(params) !== '[object Object]') {
        return;
    }

    for (const key in params) {
        // eslint-disable-next-line no-underscore-dangle
        if (params[key] && params[key]._isAMomentObject) {
            params[key] = params[key].format(DATE_TIME_FORMAT);
        }
        if (isString(key)) {
            const value = params[key];
            if (value) {
                try {
                    params[key] = isString(value) ? value.trim() : value;
                } catch (error: any) {
                    throw new Error(error);
                }
            }
        }
        if (isObject(params[key])) {
            formatRequestDate(params[key]);
        }
    }
}

// set url params
export function setObjToUrlParams(
    baseUrl: string,
    obj: { [index: string]: any },
): string {
    let parameters = '';
    for (const key in obj) {
        parameters += `${key}=${encodeURIComponent(obj[key])}&`;
    }
    parameters = parameters.replace(/&$/, '');
    return /\?$/.test(baseUrl)
        ? baseUrl + parameters
        : baseUrl.replace(/\/?$/, '?') + parameters;
}
