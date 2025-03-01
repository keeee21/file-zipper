<script setup lang="ts">
import { defineProps, defineEmits, ref } from 'vue';

const props = defineProps<{ fileData: File | null }>();
const emit = defineEmits(['update:fileData']);

const fileName = ref<string | null>(null);
const previewFile = ref<string | null>(null);
const fileInput = ref<HTMLInputElement | null>(null);
const isDragging = ref(false);

/**
 * クリックでファイル選択
 */
const selectFile = () => {
  fileInput.value?.click();
};

/**
 * ファイルを選択またはドラッグ&ドロップしたときの処理
 */
const pickFile = (event?: Event | DragEvent) => {
  let file: File | null = null;

  if (event instanceof DragEvent) {
    event.preventDefault();
    isDragging.value = false;
    if (event.dataTransfer?.files.length) {
      file = event.dataTransfer.files[0];
    }
  } else if (event instanceof Event) {
    const input = fileInput.value;
    if (!input || !input.files || input.files.length === 0) return;
    file = input.files[0];
  }

  if (file) {
    fileName.value = file.name;
    emit('update:fileData', file); // ファイルデータを親に渡す

    if (file.type.startsWith('image/')) {
      const reader = new FileReader();
      reader.onload = () => {
        previewFile.value = reader.result as string;
      };
      reader.readAsDataURL(file);
    } else {
      previewFile.value = null; // 画像でない場合、プレビューなし
    }
  }
};

/**
 * ドラッグイベント
 */
const handleDragOver = (event: DragEvent) => {
  event.preventDefault();
  isDragging.value = true;
};

const handleDragLeave = () => {
  isDragging.value = false;
};
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