<script setup lang="ts">
import { computed } from 'vue';

import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import UploadForm from '@/components/UploadForm.vue';
import { useUserStore } from '@/store/user';

const userStore = useUserStore();
const user = computed(() => userStore.user);
</script>

<template>
  <div class="flex items-center justify-center min-h-screen px-4">
    <Card class="w-full max-w-2xl text-center">
      <CardHeader v-if="user">
        <CardTitle>ようこそ、{{ user.name }} さん！</CardTitle>
      </CardHeader>

      <CardContent>
        <UploadForm />
      </CardContent>
    </Card>
  </div>
</template>
