<script setup lang="ts">
import { ref, reactive } from 'vue';
import { CreateUserParam } from '@/api/request/userModel';
import { createStaff } from '@/api/index/admin';
import { getAvatarString } from '@/utils/common/string';
import { MessagePlugin } from 'tdesign-vue-next';

const emit = defineEmits(['close', 'create'])
defineProps({
    groupDialogVisible: Boolean,
});
const onCancel = () => {
    console.log('cancel');
    emit('close', false);
};
const onSubmit = async () => {
    const param: CreateUserParam = {
        user_info: {
            avatar: staffInfo.avatar,
            email: staffInfo.email,
            nickname: staffInfo.nickname,
            password: staffInfo.password,
        },
    };
    try {
        const res = await createStaff(param);
        MessagePlugin.success('Create a new custom service staff successfully!');
        emit('close', false);
        emit('create', true);
    } catch (e) {
        console.log(e);
        MessagePlugin.error('Create error!');
    } finally {
        onCancel();
    }
};

const staffInfo = reactive({
    avatar: '',
    email: '',
    nickname: '',
    password: '',
});
</script>

<template>
      <t-dialog 
        width="50%"
        destroyOnClose
        :visible="groupDialogVisible"
        :on-close="onCancel"
        :on-cancel="onCancel"
        confirm-btn="Submit"
        cancel-btn="Cancel"
        @confirm="onSubmit"
        >
        <template #header>
            <div class="config-title" style="padding: 0px;">
                <h2 class="welcome">Group</h2>
                <h1 class="sub-title">
                    Add a new custom service staff, message will be sent to the staff's email.
                </h1>
            </div>
        </template>
        <template #body>
            <t-list :name="'Group'" :split="true" class="t-list-wrapper">
                <t-list-item>
                    <t-list-item-meta  title="Avatar" description="Custom service staff's avatar" />
                    <template #action>
                        <t-avatar> {{ getAvatarString(staffInfo.nickname) }}</t-avatar>
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="Nickname" description="*Custom service staff's nickname" />
                    <template #action>
                        <t-input 
                            type="input"
                            class="kf-config-input"
                            placeholder="nickname"
                            clearable
                            autoWidth
                            align="center"
                            v-model="staffInfo.nickname" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="Email" description="*Custom service staff's email address" />
                    <template #action>
                        <t-input 
                            type="input"
                            class="kf-config-input"
                            placeholder="email"
                            clearable
                            autoWidth
                            align="center"
                            v-model="staffInfo.email" />
                    </template>
                </t-list-item>
                <t-list-item>
                    <t-list-item-meta  title="Password" description="*Custom service staff's password" />
                    <template #action>
                        <t-input 
                            type="password"
                            class="kf-config-input"
                            placeholder="password"
                            clearable
                            autoWidth
                            align="center"
                            v-model="staffInfo.password" />
                    </template>
                </t-list-item>
                </t-list>
        </template>
      </t-dialog>
</template>

<script lang="ts">
export default {
  name: 'GroupDialog',
};
</script>

<style lang="less" scoped>
@import url('../index.less');
</style>
