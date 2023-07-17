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
