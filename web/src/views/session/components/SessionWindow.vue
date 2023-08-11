<script setup lang="ts">
import { ErrorCodes, reactive, ref, watch, nextTick } from 'vue';
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
import { ConversationItem, MessageItem } from "@/utils/open-im-sdk-wasm/types/entity";
import { SendMsgParams, GetAdvancedHistoryMsgParams } from "@/utils/open-im-sdk-wasm/types/params";

import { getUserInfo } from '@/api/index/user';
import { GetUserInfoParam } from '@/api/request/userModel';
import { GetUserInfoResponse } from '@/api/response/userModel';
import { CbEvents } from "@/utils/open-im-sdk-wasm/constant";
import { WSEvent } from "@/utils/open-im-sdk-wasm/types/entity";
import useUserStore from '@/store/user';
import useMessageStore from '@/store/openim_message';

const InitUserInfo:GetUserInfoResponse = {
  uuid: '',
  email: '',
  nickname: '',
  avatar: '',
  description: '',
  is_enable: false,
  is_admin: false,
  created_at: '',
};

const isDark = ref(false);
const messageContent = ref('')
const messages = ref<MessageItem[]>()
const senderInfo = ref(useUserStore().userInfo)
const recvInfo = ref<GetUserInfoResponse>(InitUserInfo)
const messageStore = useMessageStore()

const props = defineProps({
  session: Object as () => ConversationItem,
})

const getRecvIdFromSessionId = (sessionId: string) => {
  const ids = sessionId.split('_');
  return ids[1] === senderInfo.value.uuid ? ids[2] : ids[1];
}

watch(() => props.session, (newVal, oldVal) => {
  console.log("session changed", newVal, oldVal)
  fetchRecvInfo();
  getMsgHistory();
  scrollToBottom();
});
watch(() => messageStore, (newVal, oldVal) => {
  console.log("messageStore changed", newVal, oldVal)
  fetchRecvInfo();
  getMsgHistory();
  scrollToBottom();
});

const fetchRecvInfo = async () => {
  try {
    const params: GetUserInfoParam = {
      uuid: getRecvIdFromSessionId(props.session?.conversationID || '')
    }

    const info = await getUserInfo(params);
    recvInfo.value = info;
  } catch (e) {
    console.log(e);
    MessagePlugin.error('Fetch recv info error!');
  }
}

const sendMessage = async () => {
  try {
    const msg = await OpenIM.createTextMessage<MessageItem>(messageContent.value)
    const options: SendMsgParams = {
      recvID:  recvInfo.value.uuid,
      groupID: "",
      message: msg.data,
    };

    try {
      const res = await OpenIM.sendMessage(options);
      messageStore.pushNewMessage(msg.data);
      clearMsgContent();
      console.log(res);
    } catch (e) {
      console.log(e);
    }

  } catch (e) {
    console.log(e);
    MessagePlugin.error('Send message error!');
  }
};

const clearMsgContent = () => {
  messageContent.value = '';
}

const scrollToBottom = () => {
  const child = document.querySelector(".session-body") // 需要滚动的元素
  nextTick(() => {
    child?.scrollTo({
      top: child.scrollHeight ,
    })
  })
};

// parseJSON like {textELem: {content: 'xxxx'}}
const parseTextELemInMsgJSON = (Msg:string) => {
  console.log("Msg", Msg)
  const MsgObj = JSON.parse(Msg);
  const textElem = MsgObj.textElem;
  return textElem.content;
}

const getMsgHistory = async () => {
  const options:GetAdvancedHistoryMsgParams = {
    conversationID: props.session?.conversationID as string,
    startClientMsgID: "",
    count: 100, // now only find 100 msgs
    lastMinSeq: 0,
  };

  try {
    const { messageIDList } = await messageStore.getHistoryMessageListFromReq(options);
    console.log("messageIDList", messageIDList);
    messages.value = messageStore.historyMessageList;
  } catch (e) {
    console.log(e);
  }
};

// TODO: Add loading animation
const dataLoading = ref(false);
</script>

<template>
  <div :class="[isDark ? 'dark' : 'light', 'session-window-wrapper']">
    <div class="session-header ">
      <t-row align="middle">
        <t-col :span="11">
          <h2 class="welcome">{{ recvInfo.nickname }}</h2>
          <h1 class="sub-title">
            UUID: {{ recvInfo.uuid }} IP: {{  'xxx' }} Device: {{ 'slack' }}
          </h1>
        </t-col>

        <t-col :span="1">
          <t-button theme="default" shape="circle" variant="outline"><icon name="refresh" /></t-button> 
          <t-tooltip content="Update config info"><t-button theme="default" variant="outline" shape="circle" ><icon name="close" /></t-button></t-tooltip>
        </t-col>
      </t-row>
    </div>

    <div class="session-body"  ref="sessionBody">
      <div v-for="(item, idx) in messages" :key="idx" :class="[item.sendID == senderInfo.uuid ? 'send' : 'recv', 'message-card']">
        <div v-if="item.sendID == senderInfo.uuid" class="session-avatar">
          <t-avatar v-if="senderInfo.avatar === ''">{{ getAvatarString(senderInfo.nickname!) }}</t-avatar>
          <t-avatar v-else :alt="getAvatarString(senderInfo.nickname!)" :hide-on-load-failed="true" :image="senderInfo.avatar" />
        </div>
        <div v-else class="session-avatar">
          <t-avatar v-if="recvInfo.avatar === ''">{{ getAvatarString(recvInfo.nickname!) }}</t-avatar>
          <t-avatar v-else :alt="getAvatarString(recvInfo.nickname!)" :hide-on-load-failed="true" :image="recvInfo.avatar" />
        </div>
        <div class="session-content">
          <div class="session-content-text">
            <span>{{ item.textElem.content }}</span>
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
