// https://vitejs.dev/guide/env-and-mode.html#env-files
// Set environment variables and modes
export interface ImportMetaEnv {
    readonly VITE_API_URL: string;
    readonly VITE_API_URL_PREFIX: string;

    readonly VITE_OPENIM_API_ADDRESS: string;
    readonly VITE_OPENIM_WS_ADDRESS: string;
    readonly VITE_OPENIM_LOG_LEVEL: string;
}
