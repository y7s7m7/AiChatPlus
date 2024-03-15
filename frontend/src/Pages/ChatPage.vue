<script lang="ts" setup>
import ChatList from "../components/ChatList.vue";
import ChatViewPage from "../Views/ChatViewPage.vue";
</script>

<template>
  <div class="Box3">
    <ChatList/>
    <ChatViewPage/>
  </div>
</template>

<style lang="scss" scoped>
.Box3 {
  height: 100%;
  width: 100%;
  display: flex;
}
</style>