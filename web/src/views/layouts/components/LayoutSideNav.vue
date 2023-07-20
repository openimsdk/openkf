<script setup lang="ts">
import { ref } from 'vue';

const collapsed = ref(true);
const iconUrl = ref(
    'https://github.com/OpenIMSDK/OpenKF/assets/47499836/1cccc6f6-6baf-4849-b3f9-5901d683207c',
);

const changeCollapsed = () => {
    collapsed.value = !collapsed.value;
    iconUrl.value = collapsed.value
        ? 'https://github.com/OpenIMSDK/OpenKF/assets/47499836/1cccc6f6-6baf-4849-b3f9-5901d683207c'
        : 'https://github.com/OpenIMSDK/OpenKF/assets/47499836/e4475b49-ccc0-4d1b-a308-b5fd150a594f';
};

const changeHandler = active => {
    console.log('change', active);
};
</script>

<template>
    <t-menu
        theme="light"
        default-value="dashboard"
        :collapsed="collapsed"
        @change="changeHandler"
    >
        <template #logo>
            <img :width="collapsed ? 35 : 136" :src="iconUrl" alt="logo" />
        </template>
        <t-menu-item value="dashboard" class="menu-item">
            <template #icon>
                <t-icon name="dashboard" />
            </template>
            Dashboard
        </t-menu-item>
        <t-menu-item value="client-session">
            <template #icon>
                <t-icon name="chat" />
            </template>
            Session
        </t-menu-item>
        <t-menu-item value="access-platform">
            <template #icon>
                <t-icon name="control-platform" />
            </template>
            Platform
        </t-menu-item>
        <t-menu-item value="system-monitor">
            <template #icon>
                <t-icon name="chart" />
            </template>
            Monitor
        </t-menu-item>
        <t-menu-item value="system-config">
            <template #icon>
                <t-icon name="setting" />
            </template>
            Config
        </t-menu-item>
        <t-menu-item value="exit">
            <template #icon>
                <t-icon name="login" />
            </template>
            Exit
        </t-menu-item>
        <template #operations>
            <t-button
                class="t-demo-collapse-btn"
                variant="text"
                shape="square"
                @click="changeCollapsed"
            >
                <template #icon><t-icon name="view-list" /></template>
            </t-button>
        </template>
    </t-menu>
</template>

<style lang="less" scoped></style>
