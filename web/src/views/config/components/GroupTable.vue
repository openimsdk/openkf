<script setup lang="ts">
import { ref } from 'vue';
import { PrimaryTableCol, TableRowData, MessagePlugin } from 'tdesign-vue-next';
import { Icon } from 'tdesign-icons-vue-next';
import { getUserInfoListDefault } from '@/api/index/user';
import { updateStaffEnableStatus, deleteStaff } from '@/api/index/admin';
import { GetUserInfoResponse } from '@/api/response/userModel';
import { UpdateUserEnableStatusParam, DeleteUserParam } from '@/api/request/userModel';
import { onMounted, computed } from 'vue';
import { getAvatarString } from '@/utils/common/string';
import GroupDialog from './GroupDialog.vue';

const isDark = ref(false);
const data = ref([] as GetUserInfoResponse[]);
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
    const res = await getUserInfoListDefault(curr == undefined ? 1 : curr.current);
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
const banVisible = ref(false);
const chooseIdx = ref(-1);
const confirmBody = computed(() => {
  if (chooseIdx.value > -1) {
    const { nickname } = data.value[chooseIdx.value];
    return `After deletion, all information in ${nickname} staff will be cleared and cannot be recovered`;
  }
  return '';
});
const banBody = computed(() => {
  if (chooseIdx.value > -1) {
    const { nickname } = data.value[chooseIdx.value];
    return `${nickname} staff ban status will be changed`;
  }
  return '';
});

const handleClickBan = (row: { rowIndex: any }) => {
  chooseIdx.value = row.rowIndex;
  banVisible.value = true;
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
  banVisible.value = false;
  resetIdx();
};

const onConfirmDelete = async () => {
  const params: DeleteUserParam = {
    uuid: data.value[chooseIdx.value].uuid,
  };

  try {
    const res = await deleteStaff(params);
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
const onConfirmBan = async () => {
  // update
  data.value[chooseIdx.value].is_enable = !data.value[chooseIdx.value].is_enable;
  banVisible.value = false;

  const params: UpdateUserEnableStatusParam = {
    uuid: data.value[chooseIdx.value].uuid,
    is_enable: data.value[chooseIdx.value].is_enable,
  };

  try {
    const res = await updateStaffEnableStatus(params);
    MessagePlugin.success('Change ban status successfully');
  } catch (e) {
    // rollback
    data.value[chooseIdx.value].is_enable = !data.value[chooseIdx.value].is_enable;
    MessagePlugin.error('Change ban status failed!');
    console.log(e);
  } finally {
    dataLoading.value = false;
    resetIdx();
  }
  banVisible.value = false;
};


const groupDialogVisible = ref(false);
const AddStaff = () => {
  groupDialogVisible.value = true;
};
const handleGroupDialogClose = (data: boolean) => {
  groupDialogVisible.value = data;
};

const COLUMNS: PrimaryTableCol<TableRowData>[] = [
  {
    title: 'UUID',
    ellipsis: true,
    align: 'center',
    colKey: 'uuid',
  },
  {
    title: 'Avatar',
    ellipsis: true,
    align: 'center',
    colKey: 'avatar',
  },
  {
    title: 'Email',
    align: 'center',
    ellipsis: true,
    colKey: 'email',
  },
  {
    title: 'Nickname',
    align: 'center',
    ellipsis: true,
    colKey: 'nickname',
  },
  {
    title: 'Description',
    align: 'center',
    ellipsis: true,
    colKey: 'description',
  },
  {
    title: 'IsEnable',
    align: 'center',
    ellipsis: true,
    colKey: 'is_enable',
  },
  {
    title: 'Addition',
    align: 'center',
    ellipsis: true,
    colKey: 'is_admin',
  },
  {
    title: 'JoinTime',
    align: 'center',
    ellipsis: true,
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
          <t-list-item-meta  title="Group table" description="Check and add new staff" />
          <template #action>
              <t-button theme="default" shape="circle" variant="outline" @click="fetchData"><icon name="refresh" /></t-button> 
              <t-button theme="default" shape="circle" variant="outline" @click="AddStaff"><icon name="add" /></t-button> 
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
      table-content-width="100px"
      >
        <template #avatar="{ row }">
          <t-avatar v-if="row.avatar === ''">{{ getAvatarString(row.nickname) }}</t-avatar>
          <t-avatar v-else :alt="getAvatarString(row.nickname)" :hide-on-load-failed="true" :image="row.avatar" />
        </template>
        <template #is_enable="{ row }">
          <t-tag v-if="row.is_enable === true" theme="success" variant="light"> Enabled </t-tag>
          <t-tag v-if="row.is_enable === false" theme="danger" variant="light"> Disbaled </t-tag>
        </template>
        <template #is_admin="{ row }">
          <t-tag v-if="row.is_admin === true" theme="success" variant="outline"> Admin </t-tag>
          <t-tag v-if="row.is_admin === false" theme="warning" variant="outline"> Staff </t-tag>
        </template>
        <template #op="slotProps">
          <a class="t-button-link" @click="handleClickBan(slotProps)">Ban</a>
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
    <t-dialog
      v-model:visible="banVisible"
      header="Confirm update the currently information?"
      :body="banBody"
      :on-cancel="onCancel"
      confirm-btn="Change"
      cancel-btn="Cancel"
      @confirm="onConfirmBan"
    />
    <group-dialog :groupDialogVisible="groupDialogVisible" @close="handleGroupDialogClose" @create="fetchData" />
  </div>
</template>

<script lang="ts">
export default {
  name: 'GroupTable',
};
</script>

<style lang="less" scoped>
@import url('../index.less');
</style>
