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
import { ref, computed, onMounted } from 'vue';
import { getUserInfoListDefault } from '@/api/index/user';
import { GetUserInfoResponse } from '@/api/response/userModel';
import { toUpperCase, getAvatarString } from '@/utils/common/string';

const data = ref([] as GetUserInfoResponse[]);
const pagination = ref({ current: 1, pageSize: 0, total: 0 });
const dataLoading = ref(false);

const fetchData = async (curr: any) => {
  dataLoading.value = true;
  try {
    const res = await getUserInfoListDefault(curr == undefined ? 1 : curr.current);
    console.log(res);
    data.value = res.list;
    pagination.value = {
      ...pagination.value,
      total: res.total,
      pageSize: res.page_size,
      current: res.page,
    };
  } catch (e) {
    console.log(e);
  } finally {
    dataLoading.value = false;
  }
};

const onPageSizeChange = (size: number) => {
  pagination.value.pageSize = size;
  pagination.value.current = 1;
};
const onCurrentChange = (current: number) => {
  fetchData({ current });
  pagination.value.current = current;
};


onMounted(() => {
  fetchData(undefined);
});
</script>

<template>
    <t-card title="Member" class="user-team" :bordered="false">
        <template #actions>
            <t-button
                theme="default"
                shape="square"
                variant="text">
                <t-icon name="ellipsis" />
            </t-button>
        </template>
        <t-list :split="false">
            <t-list-item v-for="(item, index) in data" :key="index">
                <t-list-item-meta
                    :title="item.nickname"
                    :description="item.description"
                >
                  <template #image>
                    <t-avatar class="kf-list-item-avatar-size" v-if="item.avatar === ''">{{ getAvatarString(item.nickname) }}</t-avatar>
                    <t-avatar class="kf-list-item-avatar-size" v-else :alt="getAvatarString(item.nickname)" :hide-on-load-failed="true" :image="item.avatar" />
                  </template>
                </t-list-item-meta>
            </t-list-item>
        </t-list>
        <t-pagination
            class="pagination-container"
            v-model="pagination.current"
          v-model:page-size="pagination.pageSize"
          :total="pagination.total"
            showPageNumber
            :showPageSize="false"
            showPreviousAndNextBtn
            :totalContent="false"
            @page-size-change="onPageSizeChange"
            @current-change="onCurrentChange"
            size="small"
        />
    </t-card>
</template>

<script lang="ts">
export default {
    name: 'MemberCard',
};
</script>

<style scoped lang="less">
@import url('../index.less');

.kf-list-item-avatar-size {
  width: var(--td-comp-size-xxxl);
  height: var(--td-comp-size-xxxl);
}
</style>