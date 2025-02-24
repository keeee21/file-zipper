<script setup lang="ts">
import { useFileUploader } from '@/composables/useFileUploader';

const {
  previewImage,
  fileName,
  fileInput,
  isDragging,
  selectImage,
  pickFile,
  uploadImage,
  handleDragOver,
  handleDragLeave,
} = useFileUploader();
</script>

<template>
  <div class="container">
    <h1>File Zipper</h1>

    <div
      class="imagePreviewWrapper"
      :class="{ 'dragging': isDragging }"
      :style="{ 'background-image': previewImage ? `url(${previewImage})` : 'none' }"
      @click="selectImage"
      @dragover.prevent="handleDragOver"
      @dragleave="handleDragLeave"
      @drop.prevent="pickFile($event)"
    >
      <span v-if="!previewImage">Drag & Drop a file or Click to Upload</span>
    </div>

    <input
      ref="fileInput"
      type="file"
      @change="pickFile"
      accept="*"
      style="display: none"
    />

    <p v-if="fileName" class="file-name">{{ fileName }}</p>

    <button @click="uploadImage" :disabled="!previewImage">Upload Image</button>
  </div>
</template>

<style scoped>
.container {
  max-width: 600px;
  margin: 50px auto;
  text-align: center;
}

.imagePreviewWrapper {
  width: 400px;
  height: 400px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  margin: 20px auto;
  background-size: cover;
  background-position: center center;
  background-color: #f3f3f3;
  border: 4px dashed #007bff;
  font-size: 18px;
  font-weight: bold;
  color: #555;
  transition: all 0.3s ease;
  position: relative;
}

.imagePreviewWrapper.dragging {
  background-color: #e3f2fd;
  border-color: #0056b3;
}

.imagePreviewWrapper span {
  text-align: center;
  padding: 20px;
}

button {
  margin-top: 10px;
  padding: 15px 30px;
  font-size: 16px;
  background-color: #007bff;
  color: #fff;
  border: none;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}
</style>