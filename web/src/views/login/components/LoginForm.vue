<script setup lang="ts">
import type { FormRule, SubmitContext } from 'tdesign-vue-next';
import { MessagePlugin } from 'tdesign-vue-next';
import { accountLogin } from '@/api/index/login';
import { AccountLoginParam } from '@/api/request/userModel';
import { useRoute, useRouter } from 'vue-router';
import { OpenIM } from '@/api/openim';
import { OpenIMLoginConfig } from '@/constants';
import { IMLoginParam } from '@/api/request/openimModel';
import { ref, reactive } from 'vue';
import useUserStore from '@/store/user';
import useMenuStore from '@/store/menu';
import { localCache } from '@/utils/common/cache';

const formData = reactive({ email: '', password: '' });
const showPsw = ref(false);
const form = ref(null);
const isRem = ref(false);
const rules: Record<string, FormRule[]> = {
    email: [
        {
            required: true,
            email: true,
            message: 'Please enter the correct e-mail',
            type: 'warning',
            trigger: 'blur',
        },
    ],
    password: [
        {
            required: true,
            message: 'Please enter the correct password',
            type: 'warning',
            trigger: 'blur',
        },
        {
            validator: val =>
                /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,18}$/.test(val),
            message:
                'Password must be 8-18 characters long and include at least one letter and one number',
            type: 'warning',
            trigger: 'blur',
        },
    ],
};
const router = useRouter();
const route = useRoute();

const onSubmit = async (ctx: SubmitContext) => {
    console.log(ctx);
    if (ctx.validateResult === true) {
        try {
            // Login to OpenKF
            let data: AccountLoginParam = {
                email: formData.email,
                password: formData.password,
            };
            accountLogin(data)
                .then(res => {
                    // Login to OpenIM
                    const config: IMLoginParam = {
                        userID: res.uuid,
                        token: res.im_token.token,
                        platformID: OpenIMLoginConfig.PlatformID,
                        apiAddr: OpenIMLoginConfig.APIAddress,
                        wsAddr: OpenIMLoginConfig.WSAddress,
                    };

                    let temp_data = res;
                    // operationID is auto generated in login method.
                    OpenIM.login(config)
                        .then(res => {
                            MessagePlugin.success('Login success...');

                            // Store data
                            const userStore = useUserStore();
                            userStore.StoreToken(temp_data);
                            userStore.StoreInfo();

                            console.log(userStore);

                            if (isRem.value) {
                                // localCache.setCache('token', config.token);
                            } else {
                                // localCache.removeCache('token');
                            }

                            // fetch menu info
                            const menuStore = useMenuStore();
                            menuStore.StoreMenu();

                            const redirect = route.query.redirect as string;
                            const redirectUrl = redirect
                                ? decodeURIComponent(redirect)
                                : '/home/dashboard';
                            router.push(redirectUrl);
                        })
                        .catch(err => {
                            console.log(err);
                        });
                })
                .catch(err => {
                    console.log(err);
                    MessagePlugin.error('Login failed...');
                });
        } catch (e) {
            console.log(e);
            MessagePlugin.error(
                'Please enter the correct email and password...',
            );
        }
    }
};
</script>

<template>
    <t-form
        ref="form"
        labelAlign="top"
        :rules="rules"
        :data="formData"
        @submit="onSubmit"
        :class="'item-container'"
        :requiredMark="false"
    >
        <t-form-item name="email">
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

        <t-form-item name="password">
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
            <t-checkbox v-model="isRem">Remember me</t-checkbox>
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
