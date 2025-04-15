<script setup lang="ts">
import { Button } from '@/components/ui/button';

defineProps<{
  expiration: number;
}>();

defineEmits<{
  (e: 'update:expiration', value: number): void;
}>();

const options = [
  { days: 1, label: '1日' },
  { days: 3, label: '3日' },
  { days: 5, label: '5日' },
  { days: 10, label: '10日' },
  { days: 30, label: '30日' },
];
</script>

<template>
  <div class="space-y-2">
    <h2 class="text-lg font-medium">ファイルの有効期限</h2>
    <div class="flex gap-2 justify-center">
      <Button
        v-for="option in options"
        :key="option.days"
        :variant="expiration === option.days ? 'default' : 'outline'"
        @click="$emit('update:expiration', option.days)"
      >
        {{ option.label }}
      </Button>
    </div>
  </div>
</template>
