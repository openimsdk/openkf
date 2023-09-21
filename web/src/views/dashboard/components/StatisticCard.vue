<script setup lang="ts">
import { ref, nextTick, computed, onMounted } from 'vue';
import { getUserStatistics } from '@/api/index/statistics';
import { GetUserStatisticsResponse } from '@/api/response/userModel';
import { LAST_7_DAYS } from '@/constants/date';
import type { DateRangeValue, TabValue } from 'tdesign-vue-next';
import { LineChart } from 'echarts/charts';
import * as echarts from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { GridComponent, LegendComponent, TooltipComponent } from 'echarts/components';
import { GetUserStatisticsParam } from '@/api/request/userModel';

const dataLoading = ref(false);

const chartData = ref<GetUserStatisticsResponse[]>([]);

echarts.use([GridComponent, TooltipComponent, LineChart, CanvasRenderer, LegendComponent]);
let lineContainer: HTMLElement;
let lineChart: echarts.ECharts;

const updateContainer = () => {
    lineChart?.resize({
        width: lineContainer.clientWidth,
        height: lineContainer.clientHeight,
    });
};

onMounted(() => {
    fetchData();

    // Support auto scale
    window.addEventListener('resize', updateContainer, false);
    nextTick(() => {
        initChart('onlineTimeContainer');
    });
});

const initChart = (chartId: string) => {
    lineContainer = document.getElementById(chartId)!;
    lineChart = echarts.init(lineContainer);
    lineChart.setOption(fetchOption());
};

const onLineChange = (value: DateRangeValue) => {
    startTime.value = value[0] as string
    endTime.value = value[1] as string

    fetchData();
    console.log("onLineChange", value)

    lineChart.setOption(fetchOption());
};

const fetchOption = () => {
    return {
        grid: {
            top: '5%',
            right: '10px',
            left: '30px',
            bottom: '60px',
        },
        legend: {
            left: 'center',
            bottom: '0',
            orient: 'horizontal', // legend 横向布局。
            data: ['data'],
            textStyle: {
                fontSize: 12,
            },
        },
        xAxis: {
            type: 'category',
            data: chartData.value.forEach((item) => {
                const date = new Date(item.timestamp * 1000)
                const year = date.getFullYear();
                const month = date.getMonth() + 1;
                const day = date.getDate();
                const formattedDate = `${year}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}`;

                return formattedDate
            }),
            boundaryGap: false,
        },
        yAxis: {
            type: 'value',
        },
        tooltip: {
            trigger: 'item',
        },
        series: [
            {
                data: chartData.value.forEach((item) => {
                    return item.data
                }),
                type: 'line',
                smooth: true
            }
        ]
    };
}

const fetchData = async () => {
    dataLoading.value = true;
    try {
        const params: GetUserStatisticsParam = {
            type: tabValue.value as string,
            start_timestamp: (new Date(startTime.value)).getTime() / 1000, // to second
            end_timestamp: (new Date(endTime.value)).getTime() / 1000,
        }

        const res = await getUserStatistics(params);

        // Set to chartData
        res.forEach((item) => {
            chartData.value = []
            chartData.value.push({
                data: item.data,
                timestamp: item.timestamp,
            })
        })

        console.log("chartData", chartData.value)
    } catch (e) {
        console.log(e);
    } finally {
        dataLoading.value = false;
    }
};

const startTime = ref(LAST_7_DAYS[0])
const endTime = ref(LAST_7_DAYS[1])
const tabValue = ref<TabValue>('online_time');
const handlerTabsChange = (value: TabValue) => {
    console.log(value)
    initChart(value.toString())

    // fetch data
    fetchData();
};
</script>

<template>
    <t-card class="statistics-container" :bordered="false">
        <t-tabs v-model="tabValue" @change="handlerTabsChange">
            <t-tab-panel value="online_time" label="Online Time">
                <t-card :bordered="false" class="card-padding-no" title="Online Time" describe="day">
                    <template #actions>
                        <t-date-range-picker class="card-date-picker-container" :default-value="[startTime, endTime]"
                            theme="primary" mode="date" @change="onLineChange" />
                    </template>
                    <div id="onlineTimeContainer" style="width: 100%; height: 328px" />
                </t-card>
            </t-tab-panel>
            <t-tab-panel value="message_count" label="Message Count">
                <t-card :bordered="false" class="card-padding-no" title="Message Count" describe="day">
                    <template #actions>
                        <t-date-range-picker class="card-date-picker-container" :default-value="[startTime, endTime]"
                            theme="primary" mode="date" @change="onLineChange" />
                    </template>
                    <div id="messageCountContainer" style="width: 100%; height: 328px" />
                </t-card>
            </t-tab-panel>
            <t-tab-panel value="message_length" label="Message Length">
                <t-card :bordered="false" class="card-padding-no" title="Message Length" describe="day">
                    <template #actions>
                        <t-date-range-picker class="card-date-picker-container" :default-value="[startTime, endTime]"
                            theme="primary" mode="date" @change="onLineChange" />
                    </template>
                    <div id="messageLengthContainer" style="width: 100%; height: 328px" />
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