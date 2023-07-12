<script setup lang="ts">
import LoginHeader from './components/Header.vue'
import Login from './components/Login.vue'
import Register from './components/Register.vue'
import { ref } from 'vue';

const type = ref('login');
const switchType = (val: string) => {
  type.value = val;
};
</script>

<template>
    <div class="login-wrapper">
        <login-header/>
        
        <div class="login-container">
            <div class="title-container">
                <h1 class="title margin-no">Welcome!</h1>
                <h1 class="title">{{ type == 'register' ? 'Register' : 'Login' }} to OpenKF</h1>
                <div class="sub-title">
                    <p class="tip" @click="switchType(type == 'register' ? 'login' : 'register')">
                        {{ type == 'register' ? 'Jump to login' : 'Jump to Register' }}
                    </p>
                </div>
            </div>

            <login v-if="type === 'login'"/>
            <register v-else />
        </div>

        <footer class="copyright"> Copyright © 2023 OpenIMSDK open source community. All rights reserved.</footer>
    </div>
</template>

<style lang="less" scoped>
@import url('./index.less');
</style>