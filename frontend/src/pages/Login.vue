<script setup lang="ts">
const GOOGLE_CLIENT_ID = import.meta.env.VITE_GOOGLE_CLIENT_ID;
const LOCAL_STORAGE_ACCESS_TOKEN_KEY = 'accessToken';

interface GoogleCredentialResponse {
  credential: string;
  select_by?: string;
  clientId?: string;
}

const handleCredentialResponse = async (response: GoogleCredentialResponse) => {
  const idToken = response.credential;

  // バックエンドにトークン送信
  const res = await fetch('api/auth/google', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ idToken }),
  });

  const data = await res.json();
  if (data.user?.token) {
    localStorage.setItem(LOCAL_STORAGE_ACCESS_TOKEN_KEY, data.user.token);
  }
};

// @ts-expect-error: Google Identity callback のため、window に直接登録する
window.handleCredentialResponse = handleCredentialResponse;
</script>

<template>
  <div id="g_id_onload" :data-client_id="GOOGLE_CLIENT_ID" data-callback="handleCredentialResponse" data-auto_prompt="false"></div>
  <div class="g_id_signin" data-type="standard"></div>
</template>
