<script setup lang="ts">
import type { FormRule, SubmitContext } from 'tdesign-vue-next';
import { MessagePlugin } from 'tdesign-vue-next';
import { sendEmailCode, registerAdmin } from '@/api/index/login';
import { RegisterAdminParam } from '@/api/request/userModel';
import { Counter } from '@/utils/common/time';
import { useRoute, useRouter } from 'vue-router';
import { ref } from 'vue';

const COMMUNITY_INITIAL_DATA = {
    email: '',
    name: '',
};

const ADMIN_INITIAL_DATA = {
    email: '',
    nickname: '',
    password: '',
    verifyCode: '',
    checked: false,
};

const formData = ref({
    community: COMMUNITY_INITIAL_DATA,
    admin: ADMIN_INITIAL_DATA,
});
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
    name: [
        {
            required: true,
            message: 'Please enter the correct name',
            type: 'warning',
            trigger: 'blur',
        },
    ],
    nickname: [
        {
            required: true,
            message: 'Please enter the correct nickname',
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
const showPsw = ref(false);
const showCommunity = ref(true);
const [countDown, handleCounter] = Counter();

const sendCode = () => {
    // Start a count down tool
    handleCounter();

    sendEmailCode(formData.value.admin.email)
        .then(res => {
            MessagePlugin.success('Send success...');
        })
        .catch(err => {
            MessagePlugin.error('Send failed...', err.message);
        });
};

const emit = defineEmits(['registerSuccess']);
const router = useRouter();
const route = useRoute();
const onSubmit = (ctx: SubmitContext) => {
    if (ctx.validateResult === true) {
        if (!formData.value.admin.checked) {
            MessagePlugin.error(
                'Please agree to the TDesign Service Agreement and TDesign Privacy Statement',
            );
            return;
        }

        // Define the data structure of the request body.
        let data: RegisterAdminParam = {
            code: formData.value.admin.verifyCode,
            community_info: {
                description: '',
                avatar: '',
                name: formData.value.community.name,
                email: formData.value.community.email,
            },
            user_info: {
                avatar: '',
                email: formData.value.admin.email,
                nickname: formData.value.admin.nickname,
                password: formData.value.admin.password,
            },
        };

        // Send request
        registerAdmin(data)
            .then(res => {
                MessagePlugin.success('Register success...');
                emit('registerSuccess');

                // Redirect to dashboard
                const redirect = route.query.redirect as string;
                const redirectUrl = redirect
                    ? decodeURIComponent(redirect)
                    : '/home/dashboard';
                router.push(redirectUrl);
            })
            .catch(err => {
                MessagePlugin.error('Register failed...', err.message);
            });
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
        :rules="rules"
        :requiredMark="false"
    >
        <t-form-item name="community.email" v-show="showCommunity">
            <t-input
                v-model="formData.community.email"
                size="large"
                clearable
                placeholder="Please input your community email"
            >
                <template #prefix-icon>
                    <t-icon name="mail" />
                </template>
            </t-input>
        </t-form-item>

        <t-form-item name="community.name" v-show="showCommunity">
            <t-input
                v-model="formData.community.name"
                size="large"
                clearable
                placeholder="Please input your community name"
            >
                <template #prefix-icon>
                    <t-icon name="cloud" />
                </template>
            </t-input>
        </t-form-item>

        <t-form-item name="admin.email" v-show="!showCommunity">
            <t-input
                v-model="formData.admin.email"
                size="large"
                clearable
                placeholder="Please input your email"
            >
                <template #prefix-icon>
                    <t-icon name="mail" />
                </template>
            </t-input>
        </t-form-item>

        <t-form-item name="admin.nickname" v-show="!showCommunity">
            <t-input
                v-model="formData.admin.nickname"
                size="large"
                clearable
                placeholder="Please input your nickname"
            >
                <template #prefix-icon>
                    <t-icon name="user" />
                </template>
            </t-input>
        </t-form-item>

        <t-form-item name="admin.password" v-show="!showCommunity">
            <t-input
                v-model="formData.admin.password"
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

        <t-form-item
            class="verification-code"
            name="admin_verifyCode"
            v-show="!showCommunity"
        >
            <t-input
                v-model="formData.admin.verifyCode"
                size="large"
                clearable
                placeholder="Please input your verify code"
            >
                <template #prefix-icon>
                    <t-icon name="discount" />
                </template>
            </t-input>
            <t-button
                variant="outline"
                :disabled="countDown > 0"
                @click="sendCode"
            >
                {{ countDown == 0 ? 'Send Code' : `Resend in ${countDown}s` }}
            </t-button>
        </t-form-item>

        <t-form-item class="check-container" name="admin_checked">
            <t-checkbox v-model="formData.admin.checked"></t-checkbox>
            <p>
                I have read and agree to OpenKF
                <span>Service Agreement </span> and
                <span> Privacy Policy. </span>
            </p>
        </t-form-item>

        <t-form-item v-show="!showCommunity">
            <t-button block size="large" type="submit"> Register </t-button>
        </t-form-item>

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

            <t-col :span="2">
                <t-button ghost>
                    <!-- <t-button block size="large" type="submit"> -->
                    <template #default>
                        <span
                            v-show="showCommunity"
                            @click="showCommunity = false"
                            >Next</span
                        >
                        <span
                            v-show="!showCommunity"
                            @click="showCommunity = true"
                            >Back</span
                        >
                    </template>
                </t-button>
            </t-col>
        </t-row>
    </t-form>
</template>

<style lang="less" scoped>
@import url('../index.less');
</style>
