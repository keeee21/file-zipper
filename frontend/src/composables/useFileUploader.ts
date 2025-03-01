import { ref } from 'vue';

export function useFileUploader() {
  const previewFile = ref<string | null>(null);
  const fileName = ref<string | null>(null);
  const isUploading = ref(false);
  const fileInput = ref<HTMLInputElement | null>(null);
  const isDragging = ref(false);

  /**
   * 画像選択ボタンをクリック
   */
  const selectFile = () => {
    fileInput.value?.click();
  };

  /**
   * ファイルを選択したときにプレビューする
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
      fileName.value = file.name; // ← ファイル名を取得
      const reader = new FileReader();
      reader.onload = () => {
        previewFile.value = reader.result as string;
      };
      reader.readAsDataURL(file);
    }
  };

  /**
   * 画像をアップロード
   */
  const uploadFile = async () => {
    if (!previewFile.value) return;

    isUploading.value = true;

    try {
      const response = await fetch('/api/file', {
        method: 'POST',
        body: JSON.stringify({ file: previewFile.value, fileName: fileName.value }),
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        alert('File uploaded successfully!');
      } else {
        alert('File upload failed. Please try again.');
      }
    } catch (error) {
      alert('An error occurred while uploading the file.');
    } finally {
      isUploading.value = false;
    }
  };

  /**
   * ドラッグイベント処理
   */
  const handleDragOver = (event: DragEvent) => {
    event.preventDefault();
    isDragging.value = true;
  };

  const handleDragLeave = () => {
    isDragging.value = false;
  };

  return {
    previewFile,
    fileName,
    isUploading,
    fileInput,
    isDragging,
    selectFile,
    pickFile,
    uploadFile,
    handleDragOver,
    handleDragLeave,
  };
}