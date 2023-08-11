import { defineStore } from 'pinia';
import usePermissionStore  from '@/store';
import type { MenuRoute }  from '@/types/interface'
import type { MenuInfoResponse } from '@/api/response/menuModel';
import store from "../index";

const MockMenuInfo: MenuInfoResponse[] = [
    {
        path: "/home/dashboard",
        name: "Dashboard",
        icon: "dashboard",
    },
    {
        path: "/home/session",
        name: "Session",
        icon: "chat",
    },
    {
        path: "/home/platform",
        name: "Platform",
        icon: "control-platform",
    },
    {
        path: "/home/health",
        name: "Health",
        icon: "chart",
    },
    {
        path: "/home/config",
        name: "Config",
        icon: "setting",
    },
    {
        path: "/login",
        name: "Exit",
        icon: "login",
    }
]

const useStore = defineStore('menu', {
    state: () => ({
        menu_routes: MockMenuInfo,
    }),
    getters: {},
    actions: {
        async StoreMenu() {
            // TODO: fetch menu info
        }
    }
});

export default function useMenuStore() {
  return useStore(store);
}
