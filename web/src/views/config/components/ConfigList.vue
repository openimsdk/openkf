<script setup lang="ts">
import { ref } from 'vue';

const isDark = ref(false);
const cardList = [
    {
        title: "System",
        description: "Customize your OpenKF system."
    },
    {
        title: "Personal",
        description: "Update your personal information."
    },
    {
        title: "Community",
        description: "Update your community information."
    },
    {
        title: "Group",
        description: "Manage your customer service staff."
    },
    {
        title: "Bot",
        description: "Manage LLM Bots."
    },
    {
        title: "KnowledgeBase",
        description: "Manage local knowledge base."
    },
]
</script>

<template>
    <div :class="[isDark ? 'dark' : 'light', 'config-list-wrapper']">
        <div class="config-title">
            <h2 class="welcome">OpenKF</h2>
            <h1 class="sub-title">
                Manage your own custom service config.
            </h1>
        </div>
        <div v-for="item in cardList" :key="item.title" class="kf-config-list-card">
            <t-card :title="item.title" hover-shadow >
                {{ item.description }}
                <template #actions>
                    <t-tag
                        v-show="item.title === 'Bot' || item.title === 'KnowledgeBase'"
                        theme="warning"
                        shape="round"
                        variant="light">New</t-tag>
                </template>
            </t-card>
        </div>
    </div>
</template>

<script lang="ts">
export default {
  name: 'Configlist',
};
</script>

<style lang="less" scoped>
@import url('../index.less');
</style>
