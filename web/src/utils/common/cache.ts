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
