<script setup lang="ts">
import { LockKeyhole } from 'lucide-vue-next';
import { Package } from 'lucide-vue-next';
import { Hourglass } from 'lucide-vue-next';

import { Accordion, AccordionItem, AccordionTrigger, AccordionContent } from '@/components/ui/accordion';
import { Alert, AlertTitle, AlertDescription } from '@/components/ui/alert';
import { Badge } from '@/components/ui/badge';
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from '@/components/ui/card';
</script>

<template>
  <div class="flex items-center justify-center min-h-screen px-4">
    <Card class="w-full max-w-2xl space-y-6 p-6 my-5">
      <CardHeader>
        <CardTitle>File Zipperとは？</CardTitle>
        <CardDescription>ファイルを圧縮・共有できるシンプルなツールです。 </CardDescription>
      </CardHeader>

      <CardContent class="space-y-4">
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
