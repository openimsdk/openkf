<template>
    <div :class="['pop', isLeft ? 'left' : 'right']">
        <p style="word-wrap: break-word; width: 200px">{{ content }}</p>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
const props = defineProps<{
    position: string;
    content: string;
}>();
const isLeft = computed(() => {
    return props.position === 'left';
});
</script>

<style lang="less" scoped>
.pop {
    max-width: 232px;
    box-shadow: -1px 1px 3px 1px rgba(0, 0, 0, 0.16);
    border-radius: 13px 13px 13px 13px;
    opacity: 1;
    padding: 8px;
    font-weight: 400;
    line-height: 18px;
    letter-spacing: 1px;
    font-family: Arial-Regular, Arial;
    font-size: 14px;
    position: relative;
    &.left {
        background: #529f54;
        color: #ffffff;
        // &:before {
        //     content: '';
        //     position: absolute;
        //     width: 5px;
        //     height: 5px;
        //     top: calc(100%);
        //     left: -30px;
        //     border-left: 20px solid transparent;
        //     border-right: 20px solid #529f54;
        //     border-top: 5px solid transparent;
        //     border-bottom: 5px solid transparent;
        //     transform: rotateZ(-45deg);
        // }
    }
    &.right {
        background: #ffffff;
        opacity: 0.85;
        color: #474747;
        // &:before {
        //     content: '';
        //     position: absolute;
        //     width: 5px;
        //     height: 5px;
        //     top: calc(100%);
        //     right: -30px;
        //     border-left: 20px solid #ffffff;
        //     border-right: 20px solid transparent;
        //     border-top: 5px solid transparent;
        //     border-bottom: 5px solid transparent;
        //     transform: rotateZ(45deg);
        // }
    }
}
</style>
