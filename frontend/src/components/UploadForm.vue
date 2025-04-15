<script setup lang="ts">
import { Send } from 'lucide-vue-next';
import { ref } from 'vue';
import { toast } from 'vue-sonner';

import ExpirationSelect from '@/components/ExpirationSelect.vue';
import FileUpload from '@/components/FileUpload.vue';
import PasswordInput from '@/components/PasswordInput.vue';
import { Button } from '@/components/ui/button';
import { useFileUploader } from '@/composables/useFileUploader';

const { uploadFile, errorMessage, fileData } = useFileUploader();
const password = ref<string>('');
const downloadLink = ref<string>('');
const selected = ref(1);

// アップロード処理
const handleUpload = async () => {
  if (!fileData.value) {
    alert('ファイルを選択してください');
    return;
  }

  const res = await uploadFile(password.value, selected.value);
  if (res) {
    downloadLink.value = res.url;
  }
};

const copyToClipboard = async () => {
  if (downloadLink.value) {
    try {
      await navigator.clipboard.writeText(downloadLink.value);
      toast('リンクをコピーしました！');
    } catch (e) {
      toast('コピーに失敗しました');
    }
  }
};
</script>

<template>
  <div class="max-w-xl w-[90%] mx-auto p-2 bg-white rounded-xl text-center space-y-6">
    <ExpirationSelect v-model:expiration="selected" />
    <FileUpload @update:file-data="fileData = $event" />
    <PasswordInput v-model:password="password" />

    <p v-if="errorMessage" class="text-red-600 text-sm">{{ errorMessage }}</p>

    <div v-if="downloadLink" class="flex items-center justify-center gap-2">
      <Button variant="outline" size="sm" class="text-green-700 border-green-700 hover:bg-green-100 transition" @click="copyToClipboard">
        {{ downloadLink }}
      </Button>
    </div>

    <Button class="w-full flex items-center justify-center gap-2" :disabled="!!downloadLink" @click="handleUpload">
      <Send class="w-4 h-4" />
      {{ downloadLink ? 'Already Uploaded' : 'Upload File' }}
    </Button>
  </div>
</template>
