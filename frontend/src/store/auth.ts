import { defineStore } from 'pinia';

import AuthService from '@/services/AuthService';
import UserService from '@/services/UserService';

import { useUserStore } from './user';

import type { AuthState, GoogleAuthCallbackResponseUser, User } from './types';
import type { Router } from 'vue-router';

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({ token: AuthService.getAccessToken() }),
  actions: {
    /**
     * google認証後のcallbackで、認証情報を保存する
     */
    async setAuth(user: GoogleAuthCallbackResponseUser): Promise<User> {
      // Set the token in the store
      await AuthService.saveToken(user.token);

      const loginUserInfo = await UserService.getInfo();
      const userStore = useUserStore();
      if (loginUserInfo.user == null) {
        throw Error('User not found');
      }

      userStore.setUser(loginUserInfo);

      return loginUserInfo.user;
    },

    /**
     * ログアウト
     *
     * ログアウトのボタンや、各種処理でログアウトをしたい時に使われる想定。
     */
    async logout(router: Router) {
      AuthService.logout();

      const userStore = useUserStore();
      userStore.$reset();

      await router.push('/login');
    },

    /**
     * ログインチェック
     */
    async isLogin() {
      if (this.token === '') {
        return false;
      }
      const userStore = useUserStore();
      if (userStore.user == null) {
        const loginUserInfo = await UserService.getInfo();
        if (loginUserInfo.user == null) {
          throw Error('内部的なエラー');
        }
        userStore.setUser(loginUserInfo);
      }

      return true;
    },
  },
});
