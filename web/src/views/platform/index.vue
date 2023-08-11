<script setup lang="ts">
import { ref } from 'vue';
import { Icon, RefreshIcon } from 'tdesign-icons-vue-next';
import PlatformCard from './components/PlatformCard.vue';
import PlatformDialog from './components/PlatformDialog.vue';
import { PLATFORMS } from '@/constants'
import { MessagePlugin } from 'tdesign-vue-next';
import { PlatformType } from '@/constants';

const isDark = ref(false);
const platforms = ref(PLATFORMS);
const confirmVisible = ref(false);
const selectPlatformType = ref(PlatformType.Slack);

const checkUpdate = () => {
  console.log('check update');
  MessagePlugin.info('TODO...');
};

const handleManagePlatform = (name: PlatformType) => {
  confirmVisible.value = true;
  selectPlatformType.value = name;
};
const onCancel = () => {
  confirmVisible.value = false;
};

</script>

<template>
  <div :class="[isDark ? 'dark' : 'light', 'platform-wrapper']">
    <div class="platform-header ">
        <t-row align="middle">
            <t-col :span="10">
                <h2 class="welcome">Platform</h2>
                <h1 class="sub-title">
                    Check and use various platform access.
                </h1>
            </t-col>
        
            <t-col :span="2" align="center">
                <!-- TODO: check OpenKF version and update. -->
                <t-tooltip content="Check OpenKF Version [TODO...]" >
                  <t-button theme="default" variant="outline" @click="checkUpdate">
                    <template #icon><RefreshIcon /></template>
                    Check Update
                  </t-button>
                </t-tooltip>
            </t-col>
        </t-row>
    </div>

    <div class="list-card-item">
      <t-row :gutter="[16, 16]">
        <t-col
          v-for="item in platforms"
          :key="item.order"
          :lg="4"
          :xs="6"
          :xl="3"
        >
          <platform-card
            class="list-card-item"
            :platform="item"
            @manage-platform="handleManagePlatform(item.name)"
          />
        </t-col>
      </t-row>
    </div>

    <!-- Do not support delete function -->
    <platform-dialog 
      v-model:visible="confirmVisible"
      :selectPlatformType="selectPlatformType"
      @close="onCancel" />
  </div>
</template>

<script lang="ts">
export default {
  name: 'PlatformIndex',
};
</script>

<style lang="less" scoped>
@import url('./index.less');

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
