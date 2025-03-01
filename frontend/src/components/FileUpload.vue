<script setup lang="ts">
import { useFileUploader } from '@/composables/useFileUploader';

const {
  previewFile,
  fileName,
  fileInput,
  isDragging,
  selectFile,
  pickFile,
  handleDragOver,
  handleDragLeave,
} = useFileUploader();

defineExpose({
  fileName,
  previewFile,
  pickFile
});
</script>

<template>
  <div class="upload-box"
      :class="{ 'dragging': isDragging }"
      @click="selectFile"
      @dragover.prevent="handleDragOver"
      @dragleave="handleDragLeave"
      @drop.prevent="pickFile($event)"
    >
    <span v-if="!previewFile" class="upload-text">ファイルを選択 または ドラッグ&ドロップ</span>
    <img v-else :src="previewFile" class="preview-image" />
  </div>

  <input
    ref="fileInput"
    type="file"
    @change="pickFile"
    accept="*"
    style="display: none"
  />

  <p v-if="fileName" class="file-name">{{ fileName }}</p>
</template>

<style scoped>
.upload-box {
  width: 300px;
  height: 180px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border-radius: 8px;
  border: 2px dashed #007bff;
  background: #f8f9fa;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.upload-text {
  font-size: 16px;
  color: #555;
}

.preview-image {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
  border-radius: 8px;
}

.file-name {
  margin-top: 10px;
  font-size: 14px;
  color: #333;
}
</style>