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
