import { ref } from 'vue';

export function useFileUploader() {
  const previewImage = ref<string | null>(null);
  const fileName = ref<string | null>(null);
  const isUploading = ref(false);
  const fileInput = ref<HTMLInputElement | null>(null);
  const isDragging = ref(false);

  /**
   * 画像選択ボタンをクリック
   */
  const selectImage = () => {
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
        previewImage.value = reader.result as string;
      };
      reader.readAsDataURL(file);
    }
  };

  /**
   * 画像をアップロード
   */
  const uploadImage = async () => {
    if (!previewImage.value) return;

    isUploading.value = true;

    try {
      const response = await fetch('/api/image', {
        method: 'POST',
        body: JSON.stringify({ image: previewImage.value, fileName: fileName.value }),
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        alert('Image uploaded successfully!');
      } else {
        alert('Image upload failed. Please try again.');
      }
    } catch (error) {
      alert('An error occurred while uploading the image.');
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
    previewImage,
    fileName,
    isUploading,
    fileInput,
    isDragging,
    selectImage,
    pickFile,
    uploadImage,
    handleDragOver,
    handleDragLeave,
  };
}