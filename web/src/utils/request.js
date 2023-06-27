/**
 * Copyright © 2023 OpenIM open source community. All rights reserved.
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

import axios from 'axios';
import { getCookie } from './token.js';
import { startLoading, stopLoading } from './loading.js';
import router from '../router';

const service = axios.create({
    baseURL: 'http://127.0.0.1:10010/',
    withCredentials: true, // add cookie
    timeout: 5000,
});

// request interceptor
service.interceptors.request.use(
    config => {
        // start loading
        startLoading();

        const token = getCookie('token');
        if (token) {
            config.headers.token = token;
        }
        return config;
    },
    error => {
        Promise.reject(error);
    },
);

// response interceptor
service.interceptors.response.use(
    ({ data }) => {
        // stop loading
        stopLoading();

        return data;
    },
    error => {
        if (error.response.status === 401) {
            // redirect to login page
            router.replace('/login').catch(err => {
                err;
            });
        }

        Promise.reject(error);
    },
);

// used for axios
export default (url, method, submitData) => {
    return service({
        url,
        method,
        // change the data to params when the method is get
        [method === 'get' ? 'params' : 'data']: submitData,
    });
};
