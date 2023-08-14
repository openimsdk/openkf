<script setup lang="ts">
import { reactive, ref } from 'vue';
import SessionList from './components/SessionList.vue';
import SessionWindow from './components/SessionWindow.vue';
import {
  ConversationItem,
  MessageItem,
} from "@/utils/open-im-sdk-wasm/types/entity";

const session = ref<ConversationItem>();

const changeSession = (conversation: ConversationItem) => {
    session.value = conversation;
};
</script>

<template>
    <t-row class="session-container">
        <t-col :span="3" style="height: 100vh; overflow: auto;">
            <session-list @change="changeSession" />
        </t-col>
        <t-col :span="9">
            <session-window :session="session" style="box-shadow: 0px 1px 2px rgba(0, 0, 0, 0.1)" />
        </t-col>
    </t-row>
</template>

<script lang="ts">
export default {
    name: 'SessionIndex',
};
</script>

<style lang="less" scoped>
@import url('./index.less');
</style>
