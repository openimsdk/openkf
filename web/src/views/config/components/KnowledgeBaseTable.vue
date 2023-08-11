<script setup lang="ts">
import { ref, reactive } from 'vue';
import { PrimaryTableCol, TableRowData, MessagePlugin } from 'tdesign-vue-next';
import { Icon } from 'tdesign-icons-vue-next';
import { getBotInfoListDefault, getBotList } from '@/api/index/bot';
import { updateBot, deleteBot } from '@/api/index/bot';
import { GetBotInfoResponse } from '@/api/response/botModel';
import { UpdateBotParam, DeleteBotParam } from '@/api/request/botModel';
import { onMounted, computed } from 'vue';
import { getAvatarString } from '@/utils/common/string';
import BotDialog from './BotDialog.vue';

const isDark = ref(false);
const data = ref([] as GetBotInfoResponse[]);
const pagination = ref({
  showPageSize: false,
  totalContent: false,
  defaultPageSize: 5,
  total: 10,
  defaultCurrent: 1,
});
const dataLoading = ref(false);
// TODO: page change bugs here.
const fetchData = async (curr: any) => {
  // console.log(curr);
  dataLoading.value = true;
  try {
    const res = await getBotInfoListDefault(curr == undefined ? 1 : curr.current);
    data.value = res.list;
    pagination.value = {
      ...pagination.value,
      total: res.total,
      defaultPageSize: res.page_size,
      defaultCurrent: res.page,
    };
  } catch (e) {
    console.log(e);
  } finally {
    dataLoading.value = false;
    resetIdx();
  }
};

onMounted(() => {
  fetchData(undefined);
});

const confirmVisible = ref(false);
const chooseIdx = ref(-1);
const confirmBody = computed(() => {
  if (chooseIdx.value > -1) {
    const { nickname } = data.value[chooseIdx.value];
    return `After deletion, all information in ${nickname} bot will be cleared and cannot be recovered`;
  }
  return '';
});

const currBotValue: GetBotInfoResponse = reactive({
  uuid: '',
  avatar: '',
  bot_addr: '',
  bot_port: 0,
  bot_token: '',
  description: '',
  nickname: '',
});
const handleClickEdit = (row: { rowIndex: any }) => {
  chooseIdx.value = row.rowIndex;
  console.log(data.value[chooseIdx.value]);
  currBotValue.uuid = data.value[chooseIdx.value].uuid;
  currBotValue.avatar = data.value[chooseIdx.value].avatar;
  currBotValue.bot_addr = data.value[chooseIdx.value].bot_addr;
  currBotValue.bot_port = data.value[chooseIdx.value].bot_port;
  currBotValue.bot_token = data.value[chooseIdx.value].bot_token;
  currBotValue.description = data.value[chooseIdx.value].description;
  currBotValue.nickname = data.value[chooseIdx.value].nickname;

  botDialogVisible.value = true;
};
const handleClickDelete = (row: { rowIndex: any }) => {
  chooseIdx.value = row.rowIndex;
  confirmVisible.value = true;
};

const resetIdx = () => {
  chooseIdx.value = -1;
};

const onCancel = () => {
  confirmVisible.value = false;
  resetIdx();
};

const onConfirmDelete = async () => {
  const params: DeleteBotParam = {
    uuid: data.value[chooseIdx.value].uuid,
  };

  try {
    const res = await deleteBot(params);
    MessagePlugin.success('Delete successfully');
  } catch (e) {
    MessagePlugin.error('Delete failed!');
    console.log(e);
  } finally {
    dataLoading.value = false;
    resetIdx();
  }
  fetchData(undefined);
  confirmVisible.value = false;
};

const botDialogVisible = ref(false);
const AddBot = () => {
  botDialogVisible.value = true;
  createBotValue.value = true;
};

const createBotValue = ref(false)
const handleGroupDialogClose = (data: boolean) => {
  botDialogVisible.value = data;
  createBotValue.value = false;
};

const COLUMNS: PrimaryTableCol<TableRowData>[] = [
  {
    title: 'UUID',
    ellipsis: true,
    align: 'center',
    colKey: 'uuid',
  },
  {
    title: 'Name',
    ellipsis: true,
    align: 'center',
    colKey: 'name',
  },
  {
    title: 'Description',
    ellipsis: true,
    align: 'center',
    colKey: 'description',
  },
  {
    title: 'Infomation',
    ellipsis: true,
    align: 'center',
    colKey: 'info',
  },
  {
    title: 'CreatedAt',
    align: 'center',
    colKey: 'created_at',
  },
  {
    title: 'Operation',
    align: 'center',
    colKey: 'op',
    fixed: 'right',
    minWidth: 150,
  },
];

</script>

<template>
  <div :class="[isDark ? 'dark' : 'light', 'config-form-wrapper', 'kf-table-container']">
      <t-list-item>
          <t-list-item-meta  title="KnowledgeBase table (DEBUG...)" description="Manage and add new local knowledge base" />
          <template #action>
              <t-button theme="default" shape="circle" variant="outline" @click="fetchData"><icon name="refresh" /></t-button> 
              <t-button theme="default" shape="circle" variant="outline" @click="AddBot"><icon name="add" /></t-button> 
          </template>
      </t-list-item>
      <t-base-table
      row-key="index"
      :data="data"
      :columns="COLUMNS"
      :pagination="pagination"
      @page-change="fetchData"
      cellEmptyContent="-"
      tableLayout="auto"
      >
        <template #avatar="{ row }">
          <t-avatar v-if="row.avatar === ''">{{ getAvatarString(row.nickname) }}</t-avatar>
          <t-avatar v-else :alt="getAvatarString(row.nickname)" :hide-on-load-failed="true" :image="row.avatar" />
        </template>
        <template #op="slotProps">
          <a class="t-button-link" @click="handleClickEdit(slotProps)">Edit</a>
          <a class="t-button-link" @click="handleClickDelete(slotProps)">Delete</a>
        </template>
    </t-base-table>

    <t-dialog
      v-model:visible="confirmVisible"
      header="Confirm deleting the currently information?"
      :body="confirmBody"
      :on-cancel="onCancel"
      confirm-btn="Confirm"
      cancel-btn="Cancel"
      @confirm="onConfirmDelete"
    />
    <bot-dialog :botDialogVisible="botDialogVisible" :createBotValue="createBotValue" :currBotValue="currBotValue" @close="handleGroupDialogClose" @create="fetchData" />
  </div>
</template>

<script lang="ts">
export default {
  name: 'KnowledgeBaseTable',
};
</script>

<style lang="less" scoped>
@import url('../index.less');
</style>
