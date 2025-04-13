<script setup lang="ts">
import { Upload, Share, Download } from 'lucide-vue-next';

import { Alert, AlertTitle, AlertDescription } from '@/components/ui/alert';
import { Badge } from '@/components/ui/badge';
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from '@/components/ui/card';
import { Separator } from '@/components/ui/separator';
</script>

<template>
  <div class="flex items-center justify-center min-h-screen px-4">
    <Card class="w-full max-w-2xl space-y-6 p-6 my-5">
      <CardHeader>
        <CardTitle>File Zipperの使い方</CardTitle>
        <CardDescription>ファイルをアップデートしてURLを共有するだけ </CardDescription>
      </CardHeader>

      <CardContent class="space-y-4">
        <!-- ファイルの送り方 -->
        <div>
          <h3 class="text-lg font-bold mb-4">ファイルの送り方</h3>

          <div class="space-y-4">
            <Card class="p-4 flex gap-4 items-start">
              <Upload class="text-blue-500 mt-1" />
              <div>
                <p class="font-semibold">Step 1: ファイルを選択</p>
                <p class="text-sm text-muted-foreground">「Home」画面からアップロードしたいファイルを選択してください。</p>
              </div>
            </Card>

            <Card class="p-4 flex gap-4 items-start">
              <Share class="text-green-500 mt-1" />
              <div>
                <p class="font-semibold">Step 2: URLとパスワードを共有</p>
                <p class="text-sm text-muted-foreground">アップロード後に表示されたURLとパスワードを相手に共有してください。</p>
              </div>
            </Card>
          </div>
        </div>

        <Separator />

        <!-- ファイルの受け取り方 -->
        <div>
          <h3 class="text-lg font-bold mb-4">ファイルの受け取り方</h3>

          <div class="space-y-4">
            <Card class="p-4 flex gap-4 items-start">
              <Download class="text-purple-600 mt-1" />
              <div>
                <p class="font-semibold">Step 1: URLにアクセス</p>
                <p class="text-sm text-muted-foreground">送られてきたURLにアクセスします。</p>
              </div>
            </Card>

            <Card class="p-4 flex gap-4 items-start">
              <Download class="text-purple-600 mt-1" />
              <div>
                <p class="font-semibold">Step 2: ダウンロード</p>
                <p class="text-sm text-muted-foreground">パスワードを入力して、ダウンロードボタンを押すとファイルを受け取れます。</p>
              </div>
            </Card>
          </div>
        </div>
      </CardContent>

      <CardFooter class="flex flex-col items-start space-y-4">
        <Alert variant="default">
          <AlertTitle>注意事項</AlertTitle>
          <AlertDescription> ファイルはアップロード後、24時間以内にダウンロードしてください。それ以降は自動的に削除されます。 </AlertDescription>
        </Alert>

        <div class="flex flex-wrap gap-2">
          <Badge variant="secondary">無料</Badge>
          <Badge variant="secondary">モバイル対応</Badge>
        </div>
      </CardFooter>
    </Card>
  </div>
</template>
