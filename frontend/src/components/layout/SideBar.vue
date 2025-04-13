<script setup lang="ts">
import { onBeforeMount } from 'vue';
import { useRouter, useRoute } from 'vue-router';

import { SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar';
import { useAuthStore } from '@/store/auth';

import AppSidebar from '../AppSidebar.vue';

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();

onBeforeMount(async () => {
  if (!route.meta.isPublic) {
    if (!(await authStore.isLogin())) {
      await router.push('/login');
    }
  }
});
</script>

<template>
  <SidebarProvider>
    <AppSidebar />
    <main>
      <SidebarTrigger />
      <slot />
    </main>
  </SidebarProvider>
</template>
