<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';

const props = defineProps<{
  isPasswordEnabled: boolean;
  password: string;
}>();

const emit = defineEmits(['update:isPasswordEnabled', 'update:password', 'toggle']);

const togglePassword = () => {
  emit('update:isPasswordEnabled', !props.isPasswordEnabled);
  emit('toggle', !props.isPasswordEnabled);
};

const updatePassword = (event: Event) => {
  const target = event.target as HTMLInputElement;
  emit('update:password', target.value);
};
</script>

<template>
  <div class="password-toggle">
    <label>
      <input type="checkbox" :checked="isPasswordEnabled" @change="togglePassword" />
      パスワードを設定する
    </label>
  </div>

  <div v-if="isPasswordEnabled" class="password-field">
    <input type="text" :value="password" @input="updatePassword" class="password-input" />
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