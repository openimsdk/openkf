<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { getUserInfoListDefault } from '@/api/index/user';
import { GetUserInfoResponse } from '@/api/response/userModel';
import { toUpperCase, getAvatarString } from '@/utils/common/string';
import { LAST_7_DAYS } from '@/constants/date';
import type { DateRangeValue } from 'tdesign-vue-next';

const dataLoading = ref(false);


// onMounted(() => {
// });

const onLineChange = (value: DateRangeValue) => {
//   lineChart.setOption(
//     getFolderLineDataSet({
//       dateTime: value as string[],
//       ...chartColors.value,
//     }),
//   );c
};
</script>

<template>
    <t-card class="statistics-container" :bordered="false">
        <t-tabs value="second">
            <t-tab-panel value="first" label="Dialog">
                <t-card
                    :bordered="false"
                    class="card-padding-no"
                    title="Dialog service statistics.(TODO...)"
                    describe="day"
                >
                    <template #actions>
                        <t-date-range-picker
                            class="card-date-picker-container"
                            :default-value="LAST_7_DAYS"
                            theme="primary"
                            mode="date"
                            @change="onLineChange"
                        />
                    </template>
                    <div
                        id="lineContainer"
                        style="width: 100%; height: 328px"
                    />
                </t-card>
            </t-tab-panel>
            <t-tab-panel value="second" label="Rate">
                <t-card
                    :bordered="false"
                    class="card-padding-no"
                    title="Custom rate for your service.(TODO...)"
                    describe="day"
                >
                    <template #actions>
                        <t-date-range-picker
                            class="card-date-picker-container"
                            :default-value="LAST_7_DAYS"
                            theme="primary"
                            mode="date"
                            @change="onLineChange"
                        />
                    </template>
                    <div
                        id="lineContainer"
                        style="width: 100%; height: 328px"
                    />
                </t-card>
            </t-tab-panel>
        </t-tabs>
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