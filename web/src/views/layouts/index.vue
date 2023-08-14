<script setup lang="ts">
import LayoutContent from './components/LayoutContent.vue';
import LayoutSideNav from './components/LayoutSideNav.vue';
import { useGlobalEvent } from './im';
import { onMounted, ref } from 'vue';
import { OpenIM } from '@/api/openim';
import  useUserStore  from '@/store/user';
import { IMLoginParam } from '@/api/request/openimModel';
import { OpenIMLoginConfig } from '@/constants';
import { useRouter } from 'vue-router';

useGlobalEvent();
const router = useRouter();
const userStore = useUserStore();

// login again
onMounted(() => {
    loginCheck();
});

const loginCheck = () => {
    OpenIM.getLoginStatus()
        .then(res => {
            if (res.data !== 3) {
                tryLogin();
            }
        })
        .catch(tryLogin);
};

const tryLogin = async () => {
    const IMToken = userStore.im_token;
    const KFToken = userStore.kf_token;

    const config: IMLoginParam = {
        userID: userStore.userInfo.uuid,
        token: userStore.im_token.token,
        platformID: OpenIMLoginConfig.PlatformID,
        apiAddr: OpenIMLoginConfig.APIAddress,
        wsAddr: OpenIMLoginConfig.WSAddress,
    };

    if (IMToken && KFToken) {
        OpenIM.login(config)
            .then(() => {
                console.log("login again success!")
            })
            .catch(err => {
                console.log("login again error", err);
                router.push('/login');
            });
    }
};
</script>

<template>
    <t-layout class="base">
        <t-aside><layout-side-nav /></t-aside>
        <t-layout>
            <t-content><layout-content /></t-content>
        </t-layout>
    </t-layout>
</template>

<style lang="less" scoped>
.base {
    min-height: 100vh;
}
</style>
