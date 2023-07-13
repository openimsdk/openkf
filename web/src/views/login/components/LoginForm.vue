<script setup lang="ts">
import type { FormRule, SubmitContext } from 'tdesign-vue-next';
import { MessagePlugin } from 'tdesign-vue-next';
import { accountLogin } from '@/api/index/login';
import { AccountLoginParam } from '@/api/request/userModel';
import { useRoute, useRouter } from 'vue-router';
import { OpenIM } from '@/api/openim';
import { OpenIMLoginConfig } from '@/constants';
import { IMLoginParam } from '@/api/request/openimModel';
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
            // Login to OpenKF
            let data: AccountLoginParam = {
                email: formData.value.email,
                password: formData.value.password,
            };
            accountLogin(data)
                .then(res => {
                    // Login to OpenIM
                    const config: IMLoginParam = {
                        userID: res.uuid,
                        token: res.im_token.token,
                        platformID: OpenIMLoginConfig.PlatformID,
                        apiAddress: OpenIMLoginConfig.APIAddress,
                        wsAddress: OpenIMLoginConfig.WSAddress,
                    };
                    // operationID is auto generated in login method.
                    OpenIM.login(config)
                        .then(res => {
                            MessagePlugin.success('Login success...');
                            const redirect = route.query.redirect as string;
                            const redirectUrl = redirect
                                ? decodeURIComponent(redirect)
                                : '/dashboard';
                            router.push(redirectUrl);
                        })
                        .catch(err => {
                            console.log(err);
                        });
                })
                .catch(err => {
                    console.log(err);
                    MessagePlugin.error('Login failed...', err.message);
                });
        } catch (e) {
            console.log(e);
            MessagePlugin.error(e ?? 'Login failed...');
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
            <t-input
                v-model="formData.email"
                size="large"
                clearable
                placeholder="Please input your email"
            >
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
                placeholder="Please input your password"
            >
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
                    <t-button block size="large" type="submit">
                        Login
                    </t-button>
                </t-form-item>
            </t-col>
        </t-row>
    </t-form>
</template>

<style lang="less" scoped>
@import url('../index.less');
</style>
