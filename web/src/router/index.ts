/**
 * Copyright © 2023 OpenIMSDK open source community. All rights reserved.
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

import { createRouter, createWebHistory } from 'vue-router';

// router options
const routes = [
    {
        path: '/',
        redirect: '/home/dashboard',
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/login/index.vue'),
    },
    {
        path: '/home',
        name: 'Home',
        redirect: '/home/dashboard',
        component: () => import('@/views/layouts/index.vue'),
        children: [
            {
                path: 'config',
                name: 'Config',
                component: () => import('@/views/config/index.vue'),
            },
            {
                path: 'platform',
                name: 'Platform',
                component: () => import('@/views/platform/index.vue'),
            },
            {
                path: 'session',
                name: 'Session',
                component: () => import('@/views/session/index.vue'),
            },
            {
                path: 'dashboard',
                name: 'Dashboard',
                component: () => import('@/views/dashboard/index.vue'),
            },
            {
                path: 'health',
                name: 'Health',
                component: () => import('@/views/health/index.vue'),
            },
        ],
    },
];

// create router
const router = createRouter({
    history: createWebHistory(),
    routes: routes,
});

// router hook
router.beforeEach((to, from, next) => {
    console.log('beforeEach triggered');
    console.log('From:', from.path);
    console.log('To:', to.path);
    next();
});

router.afterEach((to, from) => {
    console.log('afterEach triggered');
    console.log('From:', from.path);
    console.log('To:', to.path);
});

// expose router
export default router;
