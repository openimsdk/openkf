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

<script setup lang="ts">
import { MoreIcon, ShopIcon } from 'tdesign-icons-vue-next';
import type { Platform } from '@/types/interface';
import { PropType } from 'vue';

// eslint-disable-next-line
const props = defineProps({
    platform: {
      type: Object as PropType<Platform>,
    }
});
const emit = defineEmits(['manage-platform']);
const handleClickManage = (name: string) => {
  emit('manage-platform', name);
};
</script>

<template>
    <t-card theme="poster2" :bordered="false">
      <template #avatar>
        <t-avatar size="56px">
          <template #icon>
            <t-image
              shape="circle"
              :src="platform!.avatar"
              :style="{ width: '120px', height: '120px' }"
            />
          </template>
        </t-avatar>
      </template>
      <template #status>
        <t-tag :theme="platform!.is_enable ? 'success' : 'default'" :disabled="!platform!.is_enable">{{
          platform!.is_enable ? 'On' : 'Off'
        }}</t-tag>
      </template>
      <template #content>
        <p class="list-card-item_detail--name">{{ platform!.name }}</p>
        <p class="list-card-item_detail--desc">{{ platform!.description }}</p>
      </template>
      <template #footer>
        <t-tag v-for="item in platform!.tags" :key="item" shape="mark" theme="primary" style="margin-left: 5px;">{{ item }}</t-tag>
      </template>
      <template #actions>
        <t-dropdown
          :disabled="!platform!.is_enable"
          trigger="click"
          :options="[
            {
              content: 'Manage',
              value: 'manage',
              onClick: () => handleClickManage(platform!.name),
            },
          ]"
        >
          <t-button theme="default" :disabled="!platform?.is_enable" shape="square" variant="text">
            <more-icon />
          </t-button>
        </t-dropdown>
      </template>
    </t-card>
  </template>

<script lang="ts">
export default {
  name: 'PlatformCard',
};
</script>
  
<style lang="less" scoped>
.list-card {
  height: 100%;

  &-operation {
    display: flex;
    justify-content: space-between;
    margin-bottom: var(--td-comp-margin-xxl);

    .search-input {
      width: 360px;
    }
  }

  &-item {
    padding: var(--td-comp-paddingTB-xl) var(--td-comp-paddingTB-xl);

    :deep(.t-card__header) {
      padding: 0;
    }

    :deep(.t-card__body) {
      padding: 0;
      margin-top: var(--td-comp-margin-xxl);
      margin-bottom: var(--td-comp-margin-xxl);
    }

    :deep(.t-card__footer) {
      padding: 0;
    }
  }

  &-pagination {
    padding: var(--td-comp-paddingTB-xl) var(--td-comp-paddingTB-xl);
  }

  &-loading {
    height: 100%;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
  }
}
</style>
  