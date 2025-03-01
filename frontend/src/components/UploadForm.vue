<script setup lang="ts">
import { ref } from 'vue';
import FileUpload from '@/components/FileUpload.vue';
import PasswordInput from '@/components/PasswordInput.vue';
import { useFileUploader } from '@/composables/useFileUploader';

const { uploadFile } = useFileUploader();
const fileUploadRef = ref();
const passwordInputRef = ref();

const handleUpload = () => {
  const fileName = fileUploadRef.value?.fileName;
  const isPasswordEnabled = passwordInputRef.value?.isPasswordEnabled;
  const password = passwordInputRef.value?.password;

  console.log("Uploading:", fileName);
  if (isPasswordEnabled) {
    console.log("Password:", password);
  }

  uploadFile();
};
</script>

<template>
  <div class="upload-container">
    <h1 class="title">File Zipper</h1>

    <FileUpload ref="fileUploadRef" />
    <PasswordInput ref="passwordInputRef" />

    <button @click="handleUpload" class="upload-button">
      Upload File
    </button>
  </div>
</template>

<style scoped>
.upload-container {
  max-width: 420px;
  margin: 50px auto;
  padding: 20px;
  background: white;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  border-radius: 12px;
  text-align: center;
}

.title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 15px;
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