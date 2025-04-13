<script setup lang="ts">
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from '@/components/ui/card';
import { Accordion, AccordionItem, AccordionTrigger, AccordionContent } from '@/components/ui/accordion';
import { Alert, AlertTitle, AlertDescription } from '@/components/ui/alert';
import { Badge } from '@/components/ui/badge';
import { LockKeyhole } from 'lucide-vue-next';
import { Package } from 'lucide-vue-next';
import { Hourglass } from 'lucide-vue-next';
import { Separator } from '@/components/ui/separator';
import { Upload, Share, Download } from 'lucide-vue-next';
</script>

<template>
  <div class="flex items-center justify-center min-h-screen px-4">
    <Card class="w-full max-w-2xl space-y-6 p-6 my-5">
      <CardHeader>
        <CardTitle>File Zipperとは？</CardTitle>
        <CardDescription>ファイルを圧縮・共有できるシンプルなツールです。 </CardDescription>
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

        <Separator />

        <h3 class="text-lg font-bold mb-4">File Zipperの特徴</h3>
        <Accordion type="single" collapsible class="w-full">
          <AccordionItem value="feature1">
            <AccordionTrigger>
              <LockKeyhole class="!rotate-0 !transform-none" />
              パスワード保護
            </AccordionTrigger>
            <AccordionContent> アップロード時にパスワードを設定し、ダウンロード時に入力を求めることで、セキュリティを強化します。 </AccordionContent>
          </AccordionItem>

          <AccordionItem value="feature2">
            <AccordionTrigger>
              <Package class="!rotate-0 !transoform-none" />
              ZIP形式での圧縮
            </AccordionTrigger>
            <AccordionContent> 複数のファイルをZIP形式でまとめてダウンロードできます。 </AccordionContent>
          </AccordionItem>

          <AccordionItem value="feature3">
            <AccordionTrigger>
              <Hourglass class="!rotate-0 !transoform-none" />
              有効期限付きリンク
            </AccordionTrigger>
            <AccordionContent> ダウンロードリンクには有効期限が設定されており、期限が過ぎるとアクセスできなくなります。 </AccordionContent>
          </AccordionItem>
        </Accordion>
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
