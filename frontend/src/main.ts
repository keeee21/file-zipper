import { createApp } from 'vue'
import './style.css'
import router from '@/lib/router'
import App from './App.vue'

createApp(App).use(router).mount('#app')
