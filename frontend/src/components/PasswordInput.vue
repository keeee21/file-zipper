<script setup lang="ts">
import { ref, watch } from 'vue';

const isPasswordEnabled = ref(false);
const password = ref(generateRandomPassword());

defineExpose({ isPasswordEnabled, password });

function generateRandomPassword() {
  return Math.random().toString(36).slice(-12);
}

watch(isPasswordEnabled, (enabled) => {
  password.value = enabled ? generateRandomPassword() : '';
});
</script>

<template>
  <div class="password-toggle">
    <label>
      <input type="checkbox" v-model="isPasswordEnabled" />
      パスワードを設定する
    </label>
  </div>

  <div v-if="isPasswordEnabled" class="password-field">
    <input type="text" v-model="password" class="password-input" />
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