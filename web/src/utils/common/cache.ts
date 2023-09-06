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

enum CacheType {
    Local,
    Session,
}

class CacheCls {
    storage: Storage;
    constructor(type: CacheType) {
        this.storage = type === CacheType.Local ? localStorage : sessionStorage;
    }
    setCache(key: string, value: any) {
        if (value !== null && value !== undefined) {
            this.storage.setItem(key, JSON.stringify(value));
        }
    }

    getCache(key: string) {
        const value = this.storage.getItem(key);
        if (value) {
            return JSON.parse(value);
        }
    }

    removeCache(key: string) {
        this.storage.removeItem(key);
    }

    clear() {
        this.storage.clear();
    }
}

const localCache = new CacheCls(CacheType.Local);
const sessionCache = new CacheCls(CacheType.Session);

export { localCache, sessionCache };
