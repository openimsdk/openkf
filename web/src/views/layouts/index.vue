<script setup lang="ts">
import LayoutContent from './components/LayoutContent.vue';
import LayoutSideNav from './components/LayoutSideNav.vue';
</script>

<template>
    <t-layout class="base">
        <t-aside><layout-side-nav /></t-aside>
        <t-layout>
            <t-content><layout-content /></t-content>
        </t-layout>
    </t-layout>
</template>

<style lang="less" scoped>
.base {
    min-height: 100vh;
}
</style>
