import { ref } from 'vue';

import AuthService from '@/services/AuthService';

export function useFileUploader() {
  const previewFile = ref<string | null>(null);
  const fileData = ref<File | null>(null);
  const isUploading = ref(false);
  const fileInput = ref<HTMLInputElement | null>(null);
  const isDragging = ref(false);
  const errorMessage = ref<string | null>(null);

  /**
   * 画像選択ボタンをクリック
   */
  const selectFile = () => {
    if (!fileInput.value) {
      errorMessage.value = 'ファイルが見つかりません';
      return;
    }
    fileInput.value.click();
  };

  /**
   * ファイルを選択・ドラッグ&ドロップしたときに処理
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
      fileData.value = file;

      // 画像ファイルの場合はプレビューを生成
      if (file.type.startsWith('image/')) {
        const reader = new FileReader();
        reader.onload = () => {
          previewFile.value = reader.result as string;
        };
        reader.readAsDataURL(file);
      } else {
        previewFile.value = null;
      }
      errorMessage.value = null;
    }
  };

  /**
   * ファイルをアップロード
   */
  const uploadFile = async (password: string | null, expiration: number | null) => {
    const accessToken = AuthService.getAccessToken();

    if (!fileData.value) {
      errorMessage.value = 'ファイルが選択されていません';
      return false;
    }
    const actualFileName = fileData.value.name || 'unknown';

    isUploading.value = true;
    errorMessage.value = null;

    try {
      const formData = new FormData();
      formData.append('file', fileData.value);
      formData.append('fileName', actualFileName);
      if (password) {
        formData.append('password', password);
      }

      if (expiration !== null) {
        formData.append('expiration', expiration.toString());
      }

      const response = await fetch('api/file-upload', {
        method: 'POST',
        body: formData,
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });

      if (!response.ok) {
        throw new Error('File upload failed.');
      }

      isUploading.value = false;
      return response.json();
    } catch (error) {
      console.log('❌ アップロードエラー:', error);
      errorMessage.value = 'アップロードに失敗しました。再試行してください。';
      return false;
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
    fileData,
    isUploading,
    fileInput,
    isDragging,
    errorMessage,
    selectFile,
    pickFile,
    uploadFile,
    handleDragOver,
    handleDragLeave,
  };
}
