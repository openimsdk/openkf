<!--
 Copyright © 2023 OpenKF & OpenIM open source community. All rights reserved.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

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
