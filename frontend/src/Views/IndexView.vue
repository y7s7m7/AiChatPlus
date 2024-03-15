<script setup lang="ts">

import Sidebar from "../components/Sidebar.vue";
import ChatList from "../components/ChatList.vue";
</script>

<template>
  <div class="Box2">
      <router-view></router-view>
  </div>
</template>

<style lang="scss" scoped>
.Box2{
  display: flex;
  width: 100%;
}
</style>