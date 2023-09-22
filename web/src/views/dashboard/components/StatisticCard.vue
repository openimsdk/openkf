<script setup lang="ts">
import { ref, nextTick, computed, watch, onMounted } from 'vue';
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


echarts.use([GridComponent, TooltipComponent, LineChart, CanvasRenderer, LegendComponent]);
let lineContainer: HTMLElement;
let lineChart: echarts.ECharts;
const chartData = ref<GetUserStatisticsResponse[]>([]);

const updateContainer = () => {
    lineChart?.resize({
        width: lineContainer.clientWidth,
        height: lineContainer.clientHeight,
    });
};

onMounted(() => {
    // Support auto scale
    window.addEventListener('resize', updateContainer, false);
    nextTick(() => {
        initChart(tabValue.value as string);
    });

    fetchData();
});

const initChart = (chartId: string) => {
    lineContainer = document.getElementById(chartId)!;
    lineChart = echarts.init(lineContainer);
    // lineChart.setOption({});
};

const onLineChange = (value: DateRangeValue) => {
    startTime.value = value[0] as string
    endTime.value = value[1] as string

    fetchData();
};

watch(chartData, () => {
    updateOption();
})

const updateOption = () => {
    console.log("length", chartData.value.length)
    const chartOptions =  {
        grid: {
            top: '5%',
            right: '40px',
            left: '60px',
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
            data: chartData.value.map((item) => {
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
                data: chartData.value.map((item) => {
                    return item.data
                }),
                type: 'line',
                smooth: true
            }
        ]
    };

    lineChart.setOption(chartOptions);

    if (chartData.value.length > 0) {
        lineChart.hideLoading()
    } else {
        lineChart.showLoading({
            text: 'No data available!',
            showSpinner: false,
            textColor: 'black',
            maskColor: 'rgba(255, 255, 255, 1)',
            fontSize: '26px',
            fontWeight: 'bold'
        })
    }
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
        console.log("params", params)
        console.log("res", res)

        // Set to chartData
        chartData.value = []
        for (let i = 0; i < res.length; i++) {
            chartData.value.push({
                data: res[i].data,
                timestamp: res[i].timestamp,
            })
        }

        console.log("chartData", chartData.value)
        // Update Options data
        updateOption();
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
    console.log(value.toString())
    // initChart('online_time')
    nextTick(() => {
        initChart(tabValue.value as string);
    });

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
                    <div id="online_time" style="width: 100%; height: 328px" />
                </t-card>
            </t-tab-panel>
            <t-tab-panel value="message_count" label="Message Count">
                <t-card :bordered="false" class="card-padding-no" title="Message Count" describe="day">
                    <template #actions>
                        <t-date-range-picker class="card-date-picker-container" :default-value="[startTime, endTime]"
                            theme="primary" mode="date" @change="onLineChange" />
                    </template>
                    <div id="message_count" style="width: 100%; height: 328px" />
                </t-card>
            </t-tab-panel>
            <t-tab-panel value="message_length" label="Message Length">
                <t-card :bordered="false" class="card-padding-no" title="Message Length" describe="day">
                    <template #actions>
                        <t-date-range-picker class="card-date-picker-container" :default-value="[startTime, endTime]"
                            theme="primary" mode="date" @change="onLineChange" />
                    </template>
                    <div id="message_length" style="width: 100%; height: 328px" />
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