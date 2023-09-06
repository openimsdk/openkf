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
import { ref, reactive, computed } from 'vue';
import { CreateBotParam, UpdateBotParam } from '@/api/request/botModel';
import { createBot, updateBot } from '@/api/index/bot';
import { GetBotInfoResponse } from '@/api/response/botModel';
import { getAvatarString } from '@/utils/common/string';
import { MessagePlugin } from 'tdesign-vue-next';

const emit = defineEmits(['close', 'create'])
const props = defineProps({
    botDialogVisible: Boolean,
    createBotValue: Boolean,
    currBotValue: Object as () => GetBotInfoResponse,
    // currBotValue: Object,
});
const onCancel = () => {
    console.log('cancel');
    emit('close', false);
};
const onSubmit = async () => {
    if (props.createBotValue == false && props.currBotValue != undefined) {
      // update
      const param: UpdateBotParam = {
        uuid: props.currBotValue.uuid,
        avatar: botInfo.value.avatar,
        bot_addr: botInfo.value.bot_addr,
        bot_port: Number(botInfo.value.bot_port),
        bot_token: botInfo.value.bot_token,
        description: botInfo.value.description,
        nickname: botInfo.value.nickname,
      };

      try {
        const res = await updateBot(param);
        MessagePlugin.success('Update custom service bot successfully!');
        emit('close', false);
        emit('create', true);
      } catch (e) {
          console.log(e);
          MessagePlugin.error('Update error!');
      } finally {
          onCancel();
      }
    } else {
      // create
      const param: CreateBotParam = {
        avatar: botInfo.value.avatar,
        bot_addr: botInfo.value.bot_addr,
        bot_port: Number(botInfo.value.bot_port),
        bot_token: botInfo.value.bot_token,
        description: botInfo.value.description,
        nickname: botInfo.value.nickname,
      };
      try {
          const res = await createBot(param);
          MessagePlugin.success('Create a new custom service bot successfully!');
          emit('close', false);
          emit('create', true);
      } catch (e) {
          console.log(e);
          MessagePlugin.error('Create error!');
      } finally {
          onCancel();
      }
    }
};

const botInfo = computed(() => {
  if (props.createBotValue) {
    return reactive({
      bot_addr: '',
      bot_port: '',
      bot_token: '',
      nickname: '',
      avatar: '',
      description: '',
    });
  }

  if (props.currBotValue != undefined) {
    return reactive({
      bot_addr: props.currBotValue.bot_addr,
      bot_port: props.currBotValue.bot_port,
      bot_token: props.currBotValue.bot_token,
      nickname: props.currBotValue.nickname,
      avatar: props.currBotValue.avatar,
      description: props.currBotValue.description,
    });
  }
  return reactive({
    bot_addr: '',
    bot_port: '',
    bot_token: '',
    nickname: '',
    avatar: '',
    description: '',
  });
});
</script>

<template>
      <t-dialog 
        width="50%"
        destroyOnClose
        :visible="botDialogVisible"
        :on-close="onCancel"
        :on-cancel="onCancel"
        confirm-btn="Submit"
        cancel-btn="Cancel"
        @confirm="onSubmit"
        >
        <template #header>
            <div class="config-title" style="padding: 0px;">
                <h2 class="welcome">Bot</h2>
                <h1 class="sub-title">
                    Add a new custom service LLM Bot, with your local config.
                </h1>
            </div>
        </template>
        <template #body>
            <t-list :name="'Bot'" :split="true" class="t-list-wrapper">
                <t-list-item>
                    <t-list-item-meta  title="Avatar" description="Custom service bot's avatar" />
                    <template #action>
                        <t-avatar> {{ getAvatarString(botInfo.nickname) }}</t-avatar>
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="Nickname" description="*Custom service bot's nickname" />
                    <template #action>
                        <t-input 
                            type="input"
                            class="kf-config-input"
                            placeholder="nickname"
                            clearable
                            autoWidth
                            align="center"
                            v-model="botInfo.nickname" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="BotAddr" description="*Custom service bot's address" />
                    <template #action>
                        <t-input 
                            type="text"
                            class="kf-config-input"
                            placeholder="bot_addr"
                            clearable
                            autoWidth
                            align="center"
                            v-model="botInfo.bot_addr" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="BotPort" description="*Custom service bot's port" />
                    <template #action>
                        <t-input 
                            type="text"
                            class="kf-config-input"
                            placeholder="bot_port"
                            clearable
                            autoWidth
                            align="center"
                            v-model="botInfo.bot_port" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="BotToken" description="*Custom service bot's token" />
                    <template #action>
                        <t-input 
                            type="text"
                            class="kf-config-input"
                            placeholder="bot_token"
                            clearable
                            autoWidth
                            align="center"
                            v-model="botInfo.bot_token" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="Description" description="Custom service bot's description" />
                    <template #action>
                        <t-input 
                            type="text"
                            class="kf-config-input"
                            placeholder="description"
                            clearable
                            autoWidth
                            align="center"
                            v-model="botInfo.description" />
                    </template>
                </t-list-item>
                </t-list>
        </template>
      </t-dialog>
</template>

<script lang="ts">
export default {
  name: 'BotDialog',
};
</script>

<style lang="less" scoped>
@import url('../index.less');
</style>
