import { createApp } from 'vue';
import { createPinia } from 'pinia';
import './style.css';
import router from '@/lib/router';
import App from './App.vue';
import setupAxiosInterceptors from './lib/axiosInterceptors';

setupAxiosInterceptors();

createApp(App).use(createPinia()).use(router).mount('#app');
