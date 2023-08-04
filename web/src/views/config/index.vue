<script setup lang="ts">
import  ConfigList from './components/ConfigList.vue'
import  ConfigForm from './components/ConfigForm.vue'

</script>

<template>
    <!-- <div> -->
<!-- <div class="dashboard-panel-detail"> -->
        <t-row class="config-container">
            <t-col :span="3">
                <config-list/>
            </t-col>
            <t-col :span="9">
                <config-form style="box-shadow: 0px 1px 2px rgba(0, 0, 0, 0.1);"/>
            </t-col>
        </t-row>
    <!-- </div> -->
</template>

<script lang="ts">
export default {
  name: 'ConfigIndex',
};
</script>

<style lang="less" scoped>
@import url('./index.less');
</style>
