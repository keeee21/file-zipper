<script setup lang="ts">
import { ref } from 'vue';
import FileUpload from '@/components/FileUpload.vue';
import PasswordInput from '@/components/PasswordInput.vue';
import { useFileUploader } from '@/composables/useFileUploader';
import { Button } from '@/components/ui/button';
import { Send } from 'lucide-vue-next';

const { uploadFile, errorMessage, fileData } = useFileUploader();
const password = ref<string>('');
const downloadLink = ref<string>('');

// アップロード処理
const handleUpload = async () => {
  if (!fileData.value) {
    alert('ファイルを選択してください');
    return;
  }

  const res = await uploadFile(password.value);
  if (res) {
    downloadLink.value = res.url;
  }
};
</script>

<template>
  <div class="max-w-xl w-[90%] mx-auto mt-12 p-10 bg-white shadow-md rounded-xl text-center space-y-6">
    <FileUpload @update:file-data="fileData = $event" />
    <PasswordInput v-model:password="password" />

    <p v-if="errorMessage" class="text-red-600 text-sm">{{ errorMessage }}</p>

    <p v-if="downloadLink" class="text-sm text-green-700 break-all">
      Download Link: <a :href="downloadLink" target="_blank" class="underline hover:text-green-900">{{ downloadLink }}</a>
    </p>

    <Button class="w-full flex items-center justify-center gap-2" @click="handleUpload">
      <Send class="w-4 h-4" />
      Upload File
    </Button>
  </div>
</template>
