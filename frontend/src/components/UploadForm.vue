<script setup lang="ts">
import { ref } from 'vue';
import FileUpload from '@/components/FileUpload.vue';
import PasswordInput from '@/components/PasswordInput.vue';
import { useFileUploader } from '@/composables/useFileUploader';

const { uploadFile, errorMessage, fileData } = useFileUploader();
const isPasswordEnabled = ref(false);
const password = ref('');

// パスワードトグルの処理
const handlePasswordToggle = (enabled: boolean) => {
  if (enabled) {
    password.value = generateRandomPassword();
  } else {
    password.value = '';
  }
};

// ランダムパスワード生成
function generateRandomPassword() {
  return Math.random().toString(36).slice(-12);
}

// アップロード処理
const handleUpload = async () => {
  if (!fileData.value) {
    alert('ファイルを選択してください');
    return;
  }

  const success = await uploadFile(password.value);
  if (success) {
    alert('File uploaded successfully!');
  }
};
</script>

<template>
  <div class="upload-container">
    <h1 class="title">File Zipper</h1>

    <FileUpload @update:file-data="fileData = $event" />

    <PasswordInput v-model:is-password-enabled="isPasswordEnabled" v-model:password="password" @toggle="handlePasswordToggle" />

    <span v-if="errorMessage">{{ errorMessage }}</span>

    <button class="upload-button" @click="handleUpload">Upload File</button>
  </div>
</template>

<style scoped>
.upload-container {
  max-width: 600px;
  width: 90%;
  margin: 50px auto;
  padding: 40px;
  background: white;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  border-radius: 12px;
  text-align: center;
}

.upload-button {
  width: 100%;
  padding: 12px;
  font-size: 16px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s ease;
  margin-top: 15px;
}

.upload-button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}
</style>
