<script lang="ts" setup>
import Sidebar from "./components/Sidebar.vue";
</script>

<template>
  <div class="Box">
    <div class="Title" style="--wails-draggable:drag"></div>
    <Sidebar/>
    <router-view></router-view>
  </div>
</template>

<style lang="scss" scoped>
.Title{
  height: 20px;
  position: absolute;
  top: 0; left: 0; right: 0;
  z-index: 99;
}
.Box {
  position: relative;
  display: flex;
  background: #0D1A1A;
  height: 100%;
  width: 100%;
  border-radius: 10px;
  overflow: hidden;
}
</style>
