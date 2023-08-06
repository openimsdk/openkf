<script setup lang="ts">
import { ref } from 'vue';

const isDark = ref(false);
const cardList = [
    {
        title: 'System',
        description: 'Customize your OpenKF system.',
    },
    {
        title: 'System',
        description: 'Customize your OpenKF system.',
    },
    {
        title: 'System',
        description: 'Customize your OpenKF system.',
    },
    {
        title: 'System',
        description: 'Customize your OpenKF system.',
    },
    {
        title: 'System',
        description: 'Customize your OpenKF system.',
    },
    {
        title: 'Personal',
        description: 'Update your personal information.',
    },
    {
        title: 'Community',
        description: 'Update your community information.',
    },
    {
        title: 'Group',
        description: 'Manage your customer service staff.',
    },
    {
        title: 'Bot',
        description: 'Manage LLM Bots.',
    },
    {
        title: 'KnowledgeBase',
        description: 'Manage local knowledge base.',
    },
];
</script>

<template>
    <div :class="[isDark ? 'dark' : 'light', 'session-list-wrapper']">
        <div class="session-title">
            <h2 class="welcome">OpenKF session</h2>
            <h3 class="sub-title">Start your custom service session.</h3>
        </div>
        <div
            v-for="item in cardList"
            :key="item.title"
            class="kf-session-list-card"
        >
            <t-card :title="item.title" hover-shadow>
                {{ item.description }}
                <template #actions>
                    <t-button
                        shape="circle"
                        theme="default"
                        size="small"
                        variant="outline"
                        ghost
                        ><t-icon name="close"
                    /></t-button>
                </template>
            </t-card>
        </div>
    </div>
</template>

<script lang="ts">
export default {
    name: 'SessionList',
};
</script>

<style lang="less" scoped>
@import url('../index.less');
</style>
