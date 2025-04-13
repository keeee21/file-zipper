<script setup lang="ts">
import { computed } from 'vue';
import UploadForm from '@/components/UploadForm.vue';
import { useUserStore } from '@/store/user';

const userStore = useUserStore();
const user = computed(() => userStore.user);
</script>

<template>
  <div class="flex items-center justify-center h-full w-full">
    <div class="w-full max-w-2xl text-center space-y-4">
      <div v-if="user" class="text-xl font-semibold">ようこそ、{{ user.name }} さん！</div>
      <UploadForm />
    </div>
  </div>
</template>
