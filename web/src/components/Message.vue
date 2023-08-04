<template>
    <div :class="[isLeft ? 'left' : 'right']">
        <t-avatar :image="avatarSrc" v-if="isLeft" size="large" />
        <div>
            <PopMessage :position="position" :content="content" />
            <span :class="['time', { right: !isLeft }]">{{ time }}</span>
        </div>
        <t-avatar :image="avatarSrc" v-if="!isLeft" size="large" />
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import PopMessage from '@/components/PopMessage';
const props = defineProps<{
    content: string;
    time: string;
    position: string;
    avatarSrc: string;
}>();
const isLeft = computed(() => props.position === 'left');
</script>

<style lang="less" scoped>
.left {
    display: flex;
    justify-content: flex-start;
    align-items: flex-end;
    margin-bottom: 10px;
    .t-avatar {
        margin-right: 10px;
    }
}
.right {
    display: flex;
    justify-content: flex-end;
    align-items: flex-end;
    margin-bottom: 10px;
    .t-avatar {
        margin-left: 10px;
    }
}
.time {
    font-size: 12px;
    font-family: Arial-Regular, Arial;
    font-weight: 400;
    color: #858585;
}
</style>
