import { defineStore } from 'pinia';
import type { UserState, InfoUserResponse } from './types';

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    user: null,
  }),
  getters: {
    getUser(state) {
      return state.user;
    },
  },
  actions: {
    setUser(data: InfoUserResponse) {
      this.$reset();
      this.user = data.user;

      if (data.user == null) {
        return;
      }
    },
  },
});
