<script setup lang="ts">
import { reactive, onMounted } from 'vue';
import { MessagePlugin } from 'tdesign-vue-next';
import { PlatformType } from '@/constants';
import { getSlackConfig } from '@/api/index/platform';

const emit = defineEmits(['close'])
const props = defineProps({
    confirmVisible: Boolean,
    selectPlatformType: String,
});
const onCancel = () => {
    console.log('cancel');
    emit('close', false);
};

onMounted(async () => {
  console.log(props.selectPlatformType)
  if (props.selectPlatformType == PlatformType.Slack) {
      try {
        const res = await getSlackConfig();
        slackConfig.bot_token = res.bot_token;  
        slackConfig.app_token = res.app_token;
        slackConfig.app_id = res.app_id;
        slackConfig.client_id = res.client_id;
        slackConfig.client_secret = res.client_secret;
        slackConfig.signing_secret = res.signing_secret;
        slackConfig.verification_token = res.verification_token;
      } catch (e) {
        console.log(e);
        MessagePlugin.error('Get slack config error!');
      }
    }
});

// slack config
const slackConfig = reactive({
  bot_token: '',
  app_token: '',
  app_id: '',
  client_id: '',
  client_secret: '',
  signing_secret: '',
  verification_token: '',
});

const onInstall = () => {
    // generate slack install link.
    const slackInstall = `https://slack.com/oauth/v2/authorize?client_id=${slackConfig.client_id}&install_redirect=oauth&scope=app_mentions:read,chat:write,users:read`
    emit('close', false);
    window.open(slackInstall);
};

</script>

<template>
      <t-dialog 
        width="50%"
        destroyOnClose
        :visible="confirmVisible"
        :on-close="onCancel"
        :on-cancel="onCancel"
        confirm-btn="Install"
        cancel-btn="Cancel"
        @confirm="onInstall"
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
            <t-list v-if="selectPlatformType == PlatformType.Slack" :name="PlatformType.Slack" :split="true" class="t-list-wrapper">
                <t-list-item>
                    <t-list-item-meta  title="BotToken" description="Slack bot token" />
                    <template #action>
                        <t-input 
                            type="input"
                            class="kf-config-input"
                            placeholder="bot_token"
                            clearable
                            autoWidth
                            readonly
                            align="center"
                            v-model="slackConfig.bot_token" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="AppToken" description="Slack app token" />
                    <template #action>
                        <t-input 
                            type="text"
                            class="kf-config-input"
                            placeholder="app_token"
                            clearable
                            autoWidth
                            readonly
                            align="center"
                            v-model="slackConfig.app_token" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="AppID" description="Slack app id" />
                    <template #action>
                        <t-input 
                            type="text"
                            class="kf-config-input"
                            placeholder="app_id"
                            clearable
                            autoWidth
                            readonly
                            align="center"
                            v-model="slackConfig.app_id" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="ClientId" description="Slack client id" />
                    <template #action>
                        <t-input 
                            type="text"
                            class="kf-config-input"
                            placeholder="client_id"
                            clearable
                            autoWidth
                            readonly
                            align="center"
                            v-model="slackConfig.client_id" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="ClientSecret" description="Slack client secret" />
                    <template #action>
                        <t-input 
                            type="password"
                            class="kf-config-input"
                            placeholder="client_secret"
                            clearable
                            autoWidth
                            readonly
                            align="center"
                            v-model="slackConfig.client_secret" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="SigningSecret" description="Slack signing secret" />
                    <template #action>
                        <t-input 
                            type="password"
                            class="kf-config-input"
                            placeholder="signing_secret"
                            clearable
                            autoWidth
                            readonly
                            align="center"
                            v-model="slackConfig.signing_secret" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="VerificationToken" description="Slack verification token" />
                    <template #action>
                        <t-input 
                            type="text"
                            class="kf-config-input"
                            placeholder="verification_token"
                            clearable
                            autoWidth
                            readonly
                            align="center"
                            v-model="slackConfig.verification_token" />
                    </template>
                </t-list-item>
                </t-list>
        </template>
      </t-dialog>
</template>

<script lang="ts">
export default {
  name: 'PlatformDialog',
};
</script>

<style lang="less" scoped>
@import url('../index.less');
</style>
