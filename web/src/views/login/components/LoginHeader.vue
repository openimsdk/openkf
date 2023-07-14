<script setup lang="ts">
import { ref } from 'vue';
import OpenKFLogo from '@/assets/openkf-logo.svg?component';

const navToGitHub = () => {
    window.open('https://github.com/OpenIMSDK/OpenKF');
};

const navToDoc = () => {
    window.open('https://openkf.github.io/website/');
};

const darkMode = ref(false);
const emit = defineEmits<{
    (e: 'changeMode', value: boolean): void;
}>();

const changeModeAction = () => {
    emit('changeMode', darkMode.value);
};
</script>

<template>
    <header class="login-header">
        <div class="openkf-header">
            <open-k-f-logo class="logo" />
            <span class="title">OpenKF</span>
        </div>
        <div class="operations-container">
            <t-button
                theme="default"
                shape="square"
                variant="text"
                @click="navToGitHub"
            >
                <t-icon name="logo-github-filled" class="icon" />
            </t-button>
            <t-button
                theme="default"
                shape="square"
                variant="text"
                @click="navToDoc"
            >
                <t-icon name="help-circle-filled" class="icon" />
            </t-button>
            <!-- Add theme tabs -->
            <t-switch
                v-model="darkMode"
                size="small"
                @change="changeModeAction"
            ></t-switch>
        </div>
    </header>
</template>

<style lang="less" scoped>
.login-header {
    padding: 0 var(--td-comp-paddingLR-xl);
    display: flex;
    justify-content: space-between;
    align-items: center;
    backdrop-filter: blur(20px);
    color: var(--td-text-color-primary);
    height: var(--td-comp-size-xxxl);

    .openkf-header {
        display: flex;
        align-items: center;
        height: var(--td-comp-size-xxxxxl);

        .logo {
            width: var(--td-comp-size-xxxxxl);
            height: var(--td-comp-size-xxxxxl);
        }

        .title {
            font: var(--td-font-title-large);
            font-weight: 600;
            color: var(--td-text-color-primary);
            margin-top: var(--td-comp-margin-xs);
            margin-left: var(--td-comp-margin-s);

            &.margin-no {
                margin-top: 0;
            }
        }
    }

    .operations-container {
        display: flex;
        align-items: center;
        .t-button,
        .t-switch {
            margin-left: var(--td-comp-margin-l);
        }
    }
}
</style>
