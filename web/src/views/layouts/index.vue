<script setup lang="ts">
import LayoutContent from './components/LayoutContent.vue';
import LayoutSideNav from './components/LayoutSideNav.vue';
</script>

<template>
    <div class="base">
        <t-layout>
            <t-aside><layout-side-nav /></t-aside>
            <t-layout>
                <t-content><layout-content /></t-content>
            </t-layout>
        </t-layout>
    </div>
</template>

<style lang="less" scoped>
.base {
    height: 100vh;
}
</style>
