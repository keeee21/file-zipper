<script setup lang="ts">
import SideBar from '@/components/layout/SideBar.vue';
</script>

<template>
  <div class="flex h-screen">
    <div class="h-full">
      <SideBar />
    </div>

    <main class="flex-1 h-full overflow-auto">
      <RouterView />
    </main>
  </div>
</template>
