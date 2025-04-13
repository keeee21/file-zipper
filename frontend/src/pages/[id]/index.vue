<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import { Button } from '@/components/ui/button';
import { Card, CardHeader, CardTitle, CardContent, CardFooter } from '@/components/ui/card';
import { downloadFile } from '@/composables/fileDownload';
import { useGetFileNames } from '@/composables/useGetFileNames';
import { useGetSignedUrl } from '@/composables/useGetSignedUrl';
import { useGetRoomValidity } from '@/composables/useIsValidUrl';

const route = useRoute();
const router = useRouter();

const roomId = route.params.id as string;

const fileNames = ref<string[]>([]);
const password = ref<string>('');
const isValidUrl = ref<boolean>(false);

onMounted(async () => {
  const isValidRoomRes = await useGetRoomValidity(roomId);
  isValidUrl.value = isValidRoomRes.isValid;
  if (!isValidUrl.value) {
    alert('このURLは無効です。');
    router.push('/');
  }

  const useGetFileRes = await useGetFileNames(roomId);
  if (useGetFileRes.data.fileNames) {
    fileNames.value = useGetFileRes.data.fileNames;
  } else {
    alert('ファイル情報の取得に失敗しました。');
    window.location.reload();
  }
});

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
};
</script>

<template>
  <div class="flex items-center justify-center min-h-screen px-4">
    <Card class="w-full max-w-2xl text-center">
      <CardHeader>
        <CardTitle>ファイルダウンロード</CardTitle>
      </CardHeader>

      <CardContent class="space-y-6">
        <p class="text-sm text-gray-700 break-all">
          ファイル名: <strong>{{ fileNames.join(', ') }}</strong>
        </p>

        <div class="flex items-center gap-3 text-left">
          <label for="password-input" class="min-w-[80px] text-sm text-gray-700">パスワード:</label>
          <input
            id="password-input"
            v-model="password"
            type="password"
            placeholder="パスワードを入力"
            class="flex-1 px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
      </CardContent>

      <CardFooter>
        <Button class="w-full justify-center" @click="handleDownload">ダウンロード</Button>
      </CardFooter>
    </Card>
  </div>
</template>
