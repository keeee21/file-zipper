<script setup lang="ts">
import { useRouter } from 'vue-router';
import AuthService from '@/services/AuthService';
import { useAuthStore } from '@/store/auth';

const GOOGLE_CLIENT_ID = import.meta.env.VITE_GOOGLE_CLIENT_ID;
interface GoogleCredentialResponse {
  credential: string;
  select_by?: string;
  clientId?: string;
}

const router = useRouter();
const authStore = useAuthStore();

const handleCredentialResponse = async (response: GoogleCredentialResponse) => {
  const idToken = response.credential;

  try {
    // バックエンドにトークン送信
    const res = await fetch('api/auth/google', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ idToken }),
    });

    const data = await res.json();
    if (data.user) {
      await authStore.setAuth(data.user);
    }

    router.push({ name: 'index' });
  } catch (e) {
    AuthService.logout();
    console.error('Error handling Google credential response:', e);
  }
};

// @ts-expect-error: Google Identity callback のため、window に直接登録する
window.handleCredentialResponse = handleCredentialResponse;
</script>

<template>
  <div id="g_id_onload" :data-client_id="GOOGLE_CLIENT_ID" data-callback="handleCredentialResponse" data-auto_prompt="false"></div>
  <div class="g_id_signin" data-type="standard"></div>
</template>
