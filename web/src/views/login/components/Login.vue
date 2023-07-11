<script setup lang="ts">
import type { FormRule, SubmitContext } from 'tdesign-vue-next';
import { MessagePlugin } from 'tdesign-vue-next';
import { useRoute, useRouter } from 'vue-router';
import { ref } from 'vue';

const USER_INITIAL_DATA = {
    email: '',
    password: '',
};

const formData = ref({ ...USER_INITIAL_DATA });
const showPsw = ref(false);

const router = useRouter();
const route = useRoute();

const onSubmit = async (ctx: SubmitContext) => {
  if (ctx.validateResult === true) {
    try {

      MessagePlugin.success('登陆成功');
      const redirect = route.query.redirect as string;
      const redirectUrl = redirect ? decodeURIComponent(redirect) : '/dashboard';
      router.push(redirectUrl);
    } catch (e) {
      console.log(e);
      MessagePlugin.error(e.message);
    }
  }
};
</script>

<template>
    <t-form
        ref="form"
        labelAlign="top"
        :data="formData"
        @submit="onSubmit"
        :class="'item-container'"
    >
        <t-form-item name="admin_email">
            <t-input v-model="formData.email" size="large" clearable placeholder="Please input your email" >
                <template #prefix-icon>
                    <t-icon name="mail" />
                </template>
            </t-input>
        </t-form-item>

        <t-form-item name="admin_password">
            <t-input
                v-model="formData.password"
                size="large"
                :type="showPsw ? 'text' : 'password'"
                clearable
                placeholder="Please input your password" >
                <template #prefix-icon>
                    <t-icon name="lock-on" />
                </template>
            </t-input>
        </t-form-item>

        <div class="check-container-login remember-pwd">
            <!-- TODO: Add to localstorage -->
            <t-checkbox>Remember me</t-checkbox>
            <!-- TODO: Redirect to find my password page -->
            <span class="tip">Find my password</span>
        </div>

        <t-row justify="space-between">
            <t-col :span="4">
                <!-- TODO: Disable unused method -->
                <t-space size="1">
                    <t-button ghost shape="circle" theme="primary">
                        <template #icon><t-icon name="mail" /></template>
                    </t-button>
                    <t-button ghost shape="circle" theme="primary">
                        <template #icon><t-icon name="logo-github" /></template>
                    </t-button>
                    <t-button ghost shape="circle" theme="primary">
                        <template #icon><t-icon name="logo-wecom" /></template>
                    </t-button>

                    <template #separator>
                        <t-divider layout="vertical" />
                    </template>
                </t-space>
            </t-col>

            <t-col :span="6">
                <t-form-item>
                    <t-button block size="large" type="submit"> Login </t-button>
                </t-form-item>
            </t-col>
        </t-row>

    </t-form>
</template>

<style lang="less" scoped>
@import url('../index.less');
</style>