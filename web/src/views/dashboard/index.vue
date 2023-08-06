<script setup lang="ts">
import { MessagePlugin } from 'tdesign-vue-next';
import { ref, computed, nextTick, onMounted, onUnmounted, watch } from 'vue';
import { getTimeOfDay, getNowDiffDays } from '@/utils/common/time';
import { toUpperCase, getAvatarString } from '@/utils/common/string';
import { getMyInfo } from '@/api/index/user';
import { getMyCommunityInfo } from '@/api/index/community';

import MemberCard from './components/MemberCard.vue';
import StatisticCard from './components/StatisticCard.vue';

const navToGitHub = () => {
    window.open('https://github.com/OpenIMSDK/OpenKF')
}
const navToDoc = () => {
    window.open('https://openkf.nsddd.top/')
}
const navToSlack = () => {
    window.open('https://join.slack.com/t/openimsdk/shared_invite/zt-1tmoj26uf-_FDy3dowVHBiGvLk9e5Xkg')
}

const isDark = ref(false);

const myInfo = ref({
    uuid: '',
    email: '',
    nickname: '',
    avatar: '',
    description: '',
    is_enable: false,
    is_admin: false,
    created_at: '',    
});
const communityInfo = ref({
    uuid: '',
    name: '',
    email: '',
    avatar: '',
    description: '',
});
const filteredInfo = computed(() => {
    return {
        my_uuid: myInfo.value.uuid,
        my_email: myInfo.value.email,
        my_description: myInfo.value.description,
        my_joined_time: myInfo.value.created_at,
        community_uuid: communityInfo.value.uuid,
        community_name: communityInfo.value.name,
        community_email: communityInfo.value.email,
        community_description: communityInfo.value.description,
    };
});

// TODO: Add loading animation
const dataLoading = ref(false);
onMounted(() => {
    fetchMyData();
    fetchCommunityData();
});
const fetchMyData = async () => {
    dataLoading.value = true;
    try {
    const res = await getMyInfo();
    console.log(res);
    myInfo.value = res;
  } catch (e) {
    console.log(e);
    MessagePlugin.error('Fetch my info error!');
  } finally {
    dataLoading.value = false;
  }
}
const fetchCommunityData = async () => {
    dataLoading.value = true;
    try {
    const res = await getMyCommunityInfo();
    communityInfo.value = res;
  } catch (e) {
    console.log(e);
    MessagePlugin.error('Fetch community info error!');
  } finally {
    dataLoading.value = false;
  }
}


</script>

<template>
    <div :class="[isDark ? 'dark' : 'light', 'dashboard-wrapper']">
        <div class="platform-header ">
        <t-row :gutter="[24, 24]">
        <t-col :flex="3">
            <div class="user-left-greeting">
                <div>
                    Welcome, {{ myInfo.nickname }}
                    <span class="regular">
                        Good {{ getTimeOfDay() }}, it's your {{ getNowDiffDays(myInfo.created_at) }} days to join {{ communityInfo.name }} community...
                    </span>
                </div>
                <img src="https://github.com/OpenIMSDK/OpenKF/assets/47499836/a5384675-325e-478c-8228-2410f2329872" class="logo" />
            </div>

            <t-card class="user-info-list" title="Account Info" :bordered="false">
                <template #actions>
                    <t-button theme="default" shape="square" variant="text">
                        <t-icon name="ellipsis" />
                    </t-button>
                </template>
                <t-row class="content" justify="space-between">
                    <t-col v-for="(item, index) in filteredInfo" :key="index" class="contract" span="3">
                        <div class="contract-title">
                            {{ toUpperCase(index).replaceAll('_', ' ') }}
                        </div>
                        <div class="contract-detail">
                            {{ item }}
                        </div>
                    </t-col>
                </t-row>
            </t-card>

            <statistic-card />
        </t-col>

        <t-col :flex="1">
            <t-card class="user-intro" :bordered="false">
                <t-avatar size="80px" v-if="myInfo.avatar === ''">{{ getAvatarString(myInfo.nickname) }}</t-avatar>
                <t-avatar size="80px" v-else :alt="getAvatarString(myInfo.nickname)" hide-on-load-failed="true" :image="myInfo.avatar" />

                <div class="name">{{ myInfo.nickname }}</div>
                <div class="position">{{ myInfo.description }}</div>
            </t-card>

            <member-card />

            <t-card title="Reference" class="openkf-reference-container" :bordered="false">
                <template #actions>
                    <t-button theme="default" shape="square" variant="text">
                        <t-icon name="ellipsis" />
                    </t-button>
                </template>
                <t-row class="content" :gutters="0">
                    <t-col span="3">
                        <t-button size="large" theme="default" shape="circle" variant="outline" @click="navToGitHub">
                            <t-icon name="logo-github" />
                        </t-button>
                    </t-col>
                    <t-col span="3">
                        <t-button size="large" theme="default" shape="circle" variant="outline" @click="navToDoc">
                            <t-icon name="root-list" />
                        </t-button>
                    </t-col>
                    <t-col span="3">
                        <t-button size="large" theme="default" shape="circle" variant="outline" @click="navToSlack">
                            <t-icon name="user-talk" />
                        </t-button>
                    </t-col>
                </t-row>
            </t-card>
        </t-col>
    </t-row>
    </div>
    </div>
</template>

<script lang="ts">
export default {
  components: { StatisticCard },
    name: 'UserIndex',
};
</script>
  
<style lang="less" scoped>
@import url('./index.less');

.pagination-container {
    margin-top: var(--td-comp-margin-xl);
}
</style>
