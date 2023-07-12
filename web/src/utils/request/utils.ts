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
