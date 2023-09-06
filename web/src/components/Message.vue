<!--
 Copyright © 2023 OpenKF & OpenIM open source community. All rights reserved.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

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
