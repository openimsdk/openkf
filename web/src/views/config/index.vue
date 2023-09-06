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
