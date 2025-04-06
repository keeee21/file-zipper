<script setup lang="ts">
import { computed } from 'vue';
import UploadForm from '@/components/UploadForm.vue';
import { useUserStore } from '@/store/user';

const userStore = useUserStore();
const user = computed(() => userStore.user);
</script>

<template>
  <div v-if="user">ようこそ、{{ user.name }} さん！</div>
  <UploadForm />
</template>
