<script setup lang="ts">
import { onBeforeMount } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useAuthStore } from '@/store/auth';
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
  <h1>header nav</h1>
</template>
