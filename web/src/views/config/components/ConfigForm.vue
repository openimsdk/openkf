<script setup lang="ts">
import { ref } from 'vue';
import BotTable from './BotTable.vue';
import GroupTable from './GroupTable.vue';
import KnowledgeBaseTable from './KnowledgeBaseTable.vue';
import { MessagePlugin } from 'tdesign-vue-next';
import { Icon } from 'tdesign-icons-vue-next';
import { getMyInfo, updateUserInfo } from '@/api/index/user';
import { getMyCommunityInfo, updateCommunityInfo } from '@/api/index/community';
import { updateStaffEnableStatus, deleteStaff } from '@/api/index/admin';
import { UpdateUserInfoParam } from '@/api/request/userModel';
import { UpdateCommunityParam } from '@/api/request/communityModel';
import { onMounted, computed } from 'vue';
import { getAvatarString } from '@/utils/common/string';

const isDark = ref(false);
const myInfo = ref({
    avatar: '',
    nickname: '',
    email: '',
    description: '',
});
const communityInfo = ref({
    avatar: '',
    name: '',
    email: '',
    description: '',
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
const updateMyAndCommunityInfo = async () => {
    dataLoading.value = true;
    try {
    const myParams: UpdateUserInfoParam = {
        avatar: myInfo.value.avatar,
        nickname: myInfo.value.nickname,
        email: myInfo.value.email,
        description: myInfo.value.description,
    };
    const communityParams: UpdateCommunityParam = {
        avatar: communityInfo.value.avatar,
        name: communityInfo.value.name,
        email: communityInfo.value.email,
        description: communityInfo.value.description,
    };

    const res = await updateUserInfo(myParams);
    const res2 = await updateCommunityInfo(communityParams);
    MessagePlugin.success('Update my and community info successfully!');
  } catch (e) {
    console.log(e);
    MessagePlugin.error('Update error!');
  } finally {
    dataLoading.value = false;
  }
}


</script>

<template>
    <div :class="[isDark ? 'dark' : 'light', 'config-form-wrapper']">
        <div class="config-header ">
            <t-row align="middle">
                <t-col :span="11">
                    <h2 class="welcome">Config</h2>
                    <h1 class="sub-title">
                        All configs of your custom service.
                    </h1>
                </t-col>
            
                <t-col :span="1">
                    <!-- <t-button theme="default" shape="circle" variant="outline" @click="fetchData"><icon name="refresh" /></t-button>  -->
                    <t-tooltip content="Update config info"><t-button theme="default" variant="outline" @click="updateMyAndCommunityInfo">Update</t-button></t-tooltip>
                </t-col>
        </t-row>
        </div>

        <div class="config-content">

        <t-list :name="'System'" :split="true" class="t-list-wrapper">
        <t-list-item>
            <t-list-item-meta  title="Theme" description="Customize your system theme" />
            <template #action>
                <t-dropdown :options="[{ content: 'Light', value: 1 }, { content: 'Dark', value: 2 }]">
                    <t-button
                    shape="round"
                    variant="outline">Light</t-button>
                </t-dropdown>
            </template>
        </t-list-item>
        </t-list>
        <t-list :name="'Personal'" :split="true" class="t-list-wrapper">
        <t-list-item>
            <t-list-item-meta  title="Avatar" description="Your avatar" />
            <template #action>
                <t-avatar v-if="myInfo.avatar === ''">{{ getAvatarString(myInfo.nickname) }}</t-avatar>
                <t-avatar v-else :alt="getAvatarString(myInfo.nickname)" :hide-on-load-failed="true" :image="myInfo.avatar" />
            </template>
        </t-list-item>
        <t-list-item>
            <t-list-item-meta  title="Nickname" description="Your personal nickname" />
            <template #action>
                <t-input 
                    v-model="myInfo.nickname"
                    type="input"
                    class="kf-config-input"
                    clearable
                    autoWidth
                    align="center"
                    placeholder="nickname" />
            </template>
        </t-list-item>
        <t-list-item>
            <t-list-item-meta  title="Email" description="Your personal email address" />
            <template #action>
                <t-input 
                    v-model="myInfo.email"
                    type="text"
                    class="kf-config-input"
                    clearable
                    autoWidth
                    align="center"
                    placeholder="email"/>
            </template>
        </t-list-item>
        <!-- <t-list-item>
            <t-list-item-meta  title="Password" description="Your password" />
            <template #action>
                <t-input 
                    type="password"
                    class="kf-config-input"
                    clearable
                    autoWidth
                    align="center"
                    default-value="123123" />
            </template>
        </t-list-item> -->
        <t-list-item>
            <t-list-item-meta  title="Description" description="Your personal signature" />
            <template #action>
                <t-input 
                    v-model="myInfo.description"
                    type="text"
                    class="kf-config-input"
                    clearable
                    autoWidth
                    align="center"
                    placeholder="description"/>
            </template>
        </t-list-item>
        </t-list>

        <t-list :name="'Community'" :split="true" class="t-list-wrapper">
        <t-list-item>
        <t-list-item-meta  title="Avatar" description="Community avatar" />
            <template #action>
                <t-avatar v-if="communityInfo.avatar === ''">{{ getAvatarString(communityInfo.name) }}</t-avatar>
                <t-avatar v-else :alt="getAvatarString(communityInfo.name)" :hide-on-load-failed="true" :image="communityInfo.avatar" />
            </template>
        </t-list-item>
        <t-list-item>
            <t-list-item-meta  title="Name" description="Your community nickname" />
            <template #action>
                <t-input 
                    v-model="communityInfo.name"
                    type="input"
                    class="kf-config-input"
                    clearable
                    autoWidth
                    align="center"
                    placeholder="name" />
            </template>
        </t-list-item>
        <t-list-item>
            <t-list-item-meta  title="Email" description="Your community email address" />
            <template #action>
                <t-input 
                    v-model="communityInfo.email"
                    type="text"
                    class="kf-config-input"
                    clearable
                    autoWidth
                    align="center"
                    placeholder="email" />
            </template>
        </t-list-item>
        <t-list-item>
            <t-list-item-meta  title="Description" description="Your community signature" />
            <template #action>
                <t-input 
                    v-model="communityInfo.description"
                    type="text"
                    class="kf-config-input"
                    clearable
                    autoWidth
                    align="center"
                    placeholder="description" />
            </template>
        </t-list-item>
        </t-list>

        <t-list :name="'Group'" :split="true" class="t-list-wrapper">
            <group-table/>
        </t-list>

        <t-list :name="'Bot'" :split="true" class="t-list-wrapper">
            <bot-table/>
        </t-list>

        <t-list :name="'KnowledgeBase'" :split="true" class="t-list-wrapper">
            <knowledge-base-table />
        </t-list>

        </div>
    </div>
</template>

<script lang="ts">
export default {
  components: { BotTable, GroupTable, KnowledgeBaseTable },
  name: 'ConfigForm',
};
</script>

<style lang="less" scoped>
@import url('../index.less');
</style>
