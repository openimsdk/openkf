<!--
 Copyright © 2023 OpenKF & OpenIM open source community. All rights reserved.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

<script setup lang="ts">
import LoginHeader from './components/LoginHeader.vue';
import LoginForm from './components/LoginForm.vue';
import RegisterForm from './components/RegisterForm.vue';
import { ref } from 'vue';

const type = ref('login');
const isDark = ref(false);

const switchType = (val: string) => {
    type.value = val;
};
const changeMode = (value: boolean) => {
    isDark.value = value;
    if (isDark.value) {
        document.documentElement.setAttribute('theme-mode', 'dark');
    } else {
        document.documentElement.setAttribute('theme-mode', 'light');
    }
};
</script>

<template>
    <div :class="[isDark ? 'dark' : 'light', 'login-wrapper']">
        <login-header @change-mode="changeMode" />

        <div class="login-container">
            <div class="title-container">
                <h2 class="welcome margin-no">Welcome!</h2>
                <h1 class="title">
                    {{ type == 'register' ? 'Register' : 'Login' }} to OpenKF
                </h1>
                <div class="sub-title">
                    <p
                        class="tip"
                        @click="
                            switchType(
                                type == 'register' ? 'login' : 'register',
                            )
                        "
                    >
                        {{
                            type == 'register'
                                ? 'Jump to login'
                                : 'Jump to Register'
                        }}
                    </p>
                </div>
            </div>

            <login-form v-if="type === 'login'" />
            <register-form v-else />
        </div>

        <footer class="copyright">
            Copyright © 2023 OpenIMSDK open source community. All rights
            reserved.
        </footer>
    </div>
</template>

<style lang="less" scoped>
@import url('./index.less');
</style>
