import { createRouter, createWebHistory, type Router, type RouteRecordRaw } from 'vue-router';

import LayoutNoHeader from '@/pages/Layouts/NoHeader.vue';
import LayoutPublic from '@/pages/Layouts/Public.vue';
import LayoutPrivate from '@/pages/Layouts/Private.vue';
import NotFound from '@/pages/NotFound.vue';
import InternalServerError from '@/pages/InternalServerError.vue';
import Login from '@/pages/Login.vue';
import Index from '@/pages/Index.vue';
import Download from '@/pages/[id]/index.vue'

declare module 'vue-router' {
  interface RouteMeta {
    /**
     * 認証無しで表示可能かどうか
     *
     * isPublic = true で認証無し。
     * isPublic = false で認証情報が無い場合、login画面にlocation（または router.push()）されます。
     */
    isPublic?: boolean;
    /**
     * このページのタイトル
     *
     * 画面内のヘッダーエリアと、ブラウザタイトル、両方に影響します。（予定）
     */
    title?: string;
  }
}

// 全てのURLで認証を必要とする。isPublic: true の場合だけ認証不要。
const routes: Array<RouteRecordRaw> = [
  /*
   *
   *
   * NoHeaders
   */
  {
    path: '/__layoutNoHeader',
    name: 'LayoutNoHeader',
    component: LayoutNoHeader,
    meta: { isPublic: true, noHeader: true },
    children: [
        {
            component: Index,
            path: '/',
            name: 'index',
            meta: { isPublic: false }, // ログイン無しで来て、ログイン状態にしたいので isPublic: false にしている
        },
        {
            component: Login,
            path: '/login',
            name: 'login',
            meta: { isPublic: true, title: 'Login' },
        },
        {
            component: NotFound,
            path: '/404',
            name: 'notFound',
            meta: { isPublic: true, noHeader: true, title: '404' },
        },
        {
            component: InternalServerError,
            path: '/500',
            name: 'internalServerError',
            meta: { noHeader: true, isPublic: true, title: '500' },
        },
    ],
  },

  /*
    *
    *
    * Public pages - 認証不要なヘッダー
    */
  {
    path: '/__layoutPublic',
    name: 'LayoutPublic',
    component: LayoutPublic,
    meta: { isPublic: true },
    children: [
        // {
        //     component: sample,
        //     path: '/sample',
        //     name: 'sample',
        //     meta: { isPublic: true, title: 'title' },
        // },
    ],
  },

  /*
    *
    *
    * Private pages - 認証必要なヘッダー
    */
  {
    path: '/__layoutPrivate',
    name: 'LayoutPrivate',
    component: LayoutPrivate,
    meta: { isPublic: false },
    children: [
        {
            component: Download,
            path: '/:id',
            name: 'download',
            meta: { isPublic: false, title: 'ダウンロード' },
        },
    ],
  },

  // fallback route
  {
    path: '/:pathMatch(.*)',
    name: 'urlNotFound',
    component: NotFound,
    meta: { isPublic: true, noHeader: true, title: '404' },
  },
];

const router: Router = createRouter({
  history: createWebHistory(),
  routes,
});

router.afterEach((to) => {
  document.title = `${to.meta.title || 'タイトル未設定'} | File Zipper`;
});

export default router;