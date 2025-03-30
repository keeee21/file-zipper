<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { downloadFile } from '@/composables/fileDownload'
import { useGetRoomValidity } from '@/composables/useIsValidUrl'
import { useGetFileNames } from '@/composables/useGetFileNames'
import { useGetSignedUrl } from '@/composables/useGetSignedUrl'

const route = useRoute()
const roomId = route.params.id as string

const fileNames = ref<string[]>([])
const password = ref<string>('')
const isValidUrl = ref<boolean>(false)

onMounted(async () => {
  // このroomが存在していて、かつ有効期限が切れていないかを確認する
  const isValidRoomRes = await useGetRoomValidity(roomId)
  isValidUrl.value = isValidRoomRes.isValid
  if (!isValidUrl.value) {
    alert('このURLは無効です。')
    // window.location.href = '/'
    // Note: いきなりリダイレクトは親切ではないので、レンダリング後にナビゲーションする
  }

  // ファイル名等を取得する
  const useGetFileRes = await useGetFileNames(roomId)
  if (useGetFileRes.data.fileNames) {
    fileNames.value = useGetFileRes.data.fileNames
  } else {
    alert('ファイル情報の取得に失敗しました。')
    window.location.reload()
  }
})

// ダウンロード
const handleDownload = async () => {
  const res = await useGetSignedUrl(roomId, password.value);
  const signedUrls = res.data;

  if (signedUrls.length > 0) {
    signedUrls.forEach((url, index) => {
      const name = fileNames.value[index] || `file${index + 1}`;
      downloadFile(name, url);
    });
  } else {
    alert('ダウンロードに失敗しました。パスワードが間違っている可能性があります。');
  }
}
</script>

<template>
  <div class="download-container">
    <h2 class="title">ファイルダウンロード</h2>
    <p class="filename">ファイル名: <strong>{{ fileNames }}</strong></p>

    <div class="password-row">
      <label for="password-input" class="password-label">パスワード:</label>
      <input
        id="password-input"
        v-model="password"
        type="password"
        placeholder="パスワードを入力"
        class="password-input"
      />
    </div>

    <button @click="handleDownload" class="download-button">
      ダウンロード
    </button>
  </div>
</template>

<style scoped>
.download-container {
  width: 500px;
  margin: 40px auto;
  padding: 24px;
  background-color: #fff;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.title {
  font-size: 20px;
  font-weight: bold;
  margin-bottom: 16px;
}

.filename {
  margin-bottom: 16px;
}

.password-row {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.password-label {
  min-width: 80px;
  margin-right: 8px;
}

.password-input {
  flex: 1;
  padding: 8px;
  box-sizing: border-box;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.download-button {
  padding: 10px 16px;
  background-color: #007bff;
  color: white;
  font-size: 14px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.download-button:hover {
  background-color: #0056b3;
}
</style>