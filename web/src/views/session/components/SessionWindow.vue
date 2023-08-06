<script setup lang="ts">
import { ErrorCodes, reactive, ref } from 'vue';
import { MessagePlugin } from 'tdesign-vue-next';
import { Icon } from 'tdesign-icons-vue-next';
import { getMyInfo, updateUserInfo } from '@/api/index/user';
import { getMyCommunityInfo, updateCommunityInfo } from '@/api/index/community';
import { updateStaffEnableStatus, deleteStaff } from '@/api/index/admin';
import { UpdateUserInfoParam } from '@/api/request/userModel';
import { UpdateCommunityParam } from '@/api/request/communityModel';
import { onMounted, computed } from 'vue';
import { getAvatarString } from '@/utils/common/string';
import { OpenIM } from '@/api/openim';
import { MessageItem } from "@/utils/open-im-sdk-wasm/types/entity";

const isDark = ref(false);
const messageContent = ref('');

const customInfo = reactive({
  uuid: '57daf388-1316-46ac-9523-737e82dc0d2b',
  email: '',
  nickname: '',
  avatar: '',
  description: '',
  is_enable: false,
  is_admin: false,
  created_at: '',
});

const sendMessage = async () => {

  const msg = await OpenIM.createTextMessage(messageContent.value)<MessageItem>

  console.log(msg.data);
  const content = msg.data

  const options = {
    recvID:  customInfo.uuid,
    groupID: "",
    message: msg.data,
  };

  OpenIM.sendMessage(options).then((res) => {
    console.log(res);
  }).catch(({ ErrorCodes, errMsg }) => {
    console.log(ErrorCodes, errMsg);
    MessagePlugin.error({
      content: errMsg,
    });
  }),
  console.log(messageContent.value);
};


const messages = ref([
  {
    isSender: true,
    content: 'Hello, I am a rob生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世时尚高腰裤高速度尬哭都要尬哭一段关于大股东挂上看到顾客御史大夫告诉卡给发u说ot.',
    nickname: 'Robot',
    avatar: '',
  },
  {
    isSender: true,
    content: 'Hello, I am a rob生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世时尚高腰裤高速度尬哭都要尬哭一段关于大股东挂上看到顾客御史大夫告诉卡给发u说ot.',
    nickname: 'Robot',
    avatar: '',
  },
  {
    isSender: false,
    content: 'Hello, I am a rob生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世时尚高腰裤高速度尬哭都要尬哭一段关于大股东挂上看到顾客御史大夫告诉卡给发u说ot.',
    nickname: 'Robot',
    avatar: '',
  },
  {
    isSender: true,
    content: 'Hello, I am a rob生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世时尚高腰裤高速度尬哭都要尬哭一段关于大股东挂上看到顾客御史大夫告诉卡给发u说ot.',
    nickname: 'Robot',
    avatar: '',
  },
  {
    isSender: false,
    content: 'Hello, I am a rob生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世生生世世时尚高腰裤高速度尬哭都要尬哭一段关于大股东挂上看到顾客御史大夫告诉卡给发u说ot.',
    nickname: 'Robot',
    avatar: '',
  },
]);

// TODO: Add loading animation
const dataLoading = ref(false);
</script>

<template>
  <div :class="[isDark ? 'dark' : 'light', 'session-window-wrapper']">
    <div class="session-header ">
      <t-row align="middle">
        <t-col span="11">
          <h2 class="welcome">{{ 'zzz' }}</h2>
          <h1 class="sub-title">
            UUID: {{ 'zzz' }} IP: {{  'xxx' }} Device: {{ 'slack' }}
          </h1>
        </t-col>

        <t-col span="1">
          <t-button theme="default" shape="circle" variant="outline"><icon name="refresh" /></t-button> 
          <t-tooltip content="Update config info"><t-button theme="default" variant="outline" shape="circle" ><icon name="close" /></t-button></t-tooltip>
        </t-col>
      </t-row>
    </div>

    <div class="session-body">
      <div v-for="(item, idx) in messages" :key="idx" :class="[item.isSender ? 'send' : 'recv', 'message-card']">
        <div class="session-avatar">
          <t-avatar v-if="item.avatar === ''">{{ getAvatarString(item.nickname!) }}</t-avatar>
          <t-avatar v-else :alt="getAvatarString(item.nickname!)" hide-on-load-failed="true" :image="item.avatar" />
        </div>
        <div class="session-content">
          <div class="session-content-text">
            <span>{{ item.content }}</span>
          </div>
        </div>
      </div>
    </div>

    <div class="session-input">
      <t-space size="24px" class="tools">  
        <t-button shape="round" variant="outline" size="small"><t-icon name="tools" /></t-button>
      </t-space>
      <div class="inner">
        <t-textarea v-model="messageContent" placeholder="Enter your text here..."  autofocus />
        <t-button class="send-button" @click="sendMessage">Send</t-button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: 'SessionWindow',
};
</script>

<style lang="less" scoped>
@import url('../index.less');
</style>
