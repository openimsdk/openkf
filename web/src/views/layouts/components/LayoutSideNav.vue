<script setup lang="ts">
import { ref } from 'vue';
import { computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import useMenuStore from '@/store/menu';
import { OpenIM } from '@/api/openim';
import { MessagePlugin } from 'tdesign-vue-next';

const router = useRouter();
const menuStore = useMenuStore();
const menuList = computed(() => menuStore.menu_routes);
const collapsed = ref(true);
const iconUrl = ref(
    'https://github.com/OpenIMSDK/OpenKF/assets/47499836/1cccc6f6-6baf-4849-b3f9-5901d683207c',
);

const changeCollapsed = () => {
    collapsed.value = !collapsed.value;
    iconUrl.value = collapsed.value
        ? 'https://github.com/OpenIMSDK/OpenKF/assets/47499836/1cccc6f6-6baf-4849-b3f9-5901d683207c'
        : 'https://github.com/OpenIMSDK/OpenKF/assets/47499836/a5384675-325e-478c-8228-2410f2329872';
};

const changeHandler = (active:string) => {
    console.log('change', active);
};


const goHome = async () => {
    try {
        await OpenIM.logout();
    } catch (e) {
        MessagePlugin.error('IM logout failed!');
        console.log(e);
    } finally {
        MessagePlugin.success('Logout successfully!');
        router.push('/home/dashboard');
    }
};

</script>

<template>
    <t-menu
        theme="light"
        default-value="dashboard"
        :collapsed="collapsed"
        @change="changeHandler"
        height="100vh"
    >
        <template #logo>
            <img :width="collapsed ? 50 : 200" :src="iconUrl" alt="logo" @click="goHome" />
        </template>
        <t-menu-item v-for="item in menuList" :key="item.path" :value="item.path" :to="item.path">
            <template #icon>
                <t-icon :name="item.icon" />
            </template>
            {{ item.name }}
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

<style lang="less" scoped>
</style>
