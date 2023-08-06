<script setup lang="ts">
import { computed, onMounted } from 'vue';
</script>

<template>
  <router-view v-slot="{ Component }" class="content-container">
    <transition name="fade">
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </transition>
  </router-view>
</template>

<style lang="less" scoped>
.content-container {
  max-height: 100vh;
  overflow: auto;
}

.fade-leave-active,
.fade-enter-active {
//   transition: opacity @anim-duration-slow @anim-time-fn-easing;
//   transition: opacity;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
