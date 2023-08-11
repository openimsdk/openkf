<script setup lang="ts">
import { ref } from 'vue';

const emit = defineEmits(['change'])
const isDark = ref(false);
const selectedConfigOrder = ref<number>(1);

interface ConfigCard {
    order: number;
    title: string;
    description: string;
}
const cardList:ConfigCard[] = [
    {
        order: 1,
        title: "System",
        description: "Customize your OpenKF system."
    },
    {
        order: 2,
        title: "Personal",
        description: "Update your personal information."
    },
    {
        order: 3,
        title: "Community",
        description: "Update your community information."
    },
    {
        order: 4,
        title: "Group",
        description: "Manage your customer service staff."
    },
    {
        order: 5,
        title: "Bot",
        description: "Manage LLM Bots."
    },
    {
        order: 6,
        title: "KnowledgeBase",
        description: "Manage local knowledge base."
    },
]

const changeConfig = (config: ConfigCard) => {
    // change bind value
    selectedConfigOrder.value = config.order;
    console.log("Change configCard", config.order)
    emit('change', config)
}
</script>

<template>
    <div :class="[isDark ? 'dark' : 'light', 'config-list-wrapper']">
        <div class="config-title">
            <h2 class="welcome">OpenKF</h2>
            <h1 class="sub-title">
                Manage your own custom service config.
            </h1>
        </div>
        <div 
            v-for="item in cardList"
            :key="item.title" 
            :class="[selectedConfigOrder === item.order ? 'kf-selected-card' : '', 'kf-config-list-card']">
            <t-card :title="item.title" hover-shadow @click="changeConfig(item)">
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
