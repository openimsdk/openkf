<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { onMounted, computed } from 'vue';
import {
  ConversationItem,
  MessageItem,
} from "@/utils/open-im-sdk-wasm/types/entity";
import { getTimeFromDate } from '@/utils/common/time';
import useSessionStore from '@/store/openim_session'

const emit = defineEmits(['change'])
const isDark = ref(false);

const selectedSessionId = ref<string>("");
const sessionList = ref<ConversationItem[]>([]);
const sessionStore = useSessionStore();

onMounted(() => {
    getAllSession();
});

watch(() => sessionStore.conversationList, (newVal, oldVal) => {
  console.log("session changed", newVal, oldVal)
  getAllSession();
});

const getAllSession = async () => {
    sessionList.value = sessionStore.conversationList;
};

// parseJSON like {textELem: {content: 'xxxx'}}
const parseTextELemInLastMsgJSON = (lastMsg:string) => {
    const lastMsgObj = JSON.parse(lastMsg);
    const textElem = lastMsgObj.textElem;
    return textElem.content;
}

const changeSession = (conversation: ConversationItem) => {
    // change bind value
    selectedSessionId.value = conversation.conversationID;
    console.log("Change conversationID", conversation.conversationID)
    emit('change', conversation)
}
</script>

<template>
    <div :class="[isDark ? 'dark' : 'light', 'session-list-wrapper']">
        <div class="session-title">
            <h2 class="welcome">OpenKF session</h2>
            <h3 class="sub-title">Start your custom service session.</h3>
        </div>
        <div
            v-for="item in sessionList"
            :key="item.conversationID"
            :class="[selectedSessionId === item.conversationID ? 'kf-select-session' : '', 'kf-session-list-card']"
        >
            <t-card :title="item.showName" hover-shadow @click="changeSession(item)">
                <t-row justify="space-between">
                    <t-col :span="6" class="msg-ellipsis">
                        {{ parseTextELemInLastMsgJSON(item.latestMsg) }}
                    </t-col>
                    <t-col :span="3" class="msg-time">
                        {{ getTimeFromDate(item.latestMsgSendTime) }}
                    </t-col>
                </t-row>
                
                
                <template #actions>
                    <t-button
                        shape="circle"
                        theme="default"
                        size="small"
                        variant="outline"
                        ghost
                        ><t-icon name="close"
                    /></t-button>
                </template>
            </t-card>
        </div>
    </div>
</template>

<script lang="ts">
export default {
    name: 'SessionList',
};
</script>

<style lang="less" scoped>
@import url('../index.less');
</style>
