<script setup lang="ts">
import { ref } from 'vue';

import { Checkbox } from '@/components/ui/checkbox';

const password = defineModel<string>('password');
const isEnabled = ref(false);

// ランダムパスワード生成
function generateRandomPassword() {
  return Math.random().toString(36).slice(-12);
}

const handleToggle = (checked: boolean | 'indeterminate') => {
  if (checked === true) {
    password.value = generateRandomPassword();
  } else {
    password.value = '';
  }
};
</script>

<template>
  <div class="password-toggle">
    <Checkbox id="passowrd-check" v-model="isEnabled" @update:model-value="handleToggle" />
    <label for="passowrd-check" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
      パスワードを設定する（任意）
    </label>
  </div>

  <div v-if="isEnabled" class="password-field">
    <input v-model="password" type="text" class="password-input" placeholder="パスワードを入力" />
  </div>
</template>

<style scoped>
.password-toggle {
  margin: 15px 0;
  font-size: 14px;
}

.password-field {
  display: flex;
  justify-content: center;
  margin-bottom: 10px;
}

.password-input {
  width: 80%;
  padding: 10px;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 6px;
  text-align: center;
}
</style>
