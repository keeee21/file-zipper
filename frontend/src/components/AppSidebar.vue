<script setup lang="ts">
import { Home, BookOpen, GraduationCap, Settings } from 'lucide-vue-next';
import { useRouter } from 'vue-router';

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '@/components/ui/sidebar';

const router = useRouter();

// Menu items.
const items = [
  {
    title: 'Home',
    url: '/',
    icon: Home,
  },
  {
    title: 'What is File Zipper?',
    url: '/what-is',
    icon: BookOpen,
  },
  {
    title: 'How to use',
    url: '/how-to-use',
    icon: GraduationCap,
  },
  {
    title: 'Settings',
    url: '/settings',
    icon: Settings,
  },
];

// ページ遷移処理
const navigate = (url: string) => {
  router.push(url);
};
</script>

<template>
  <Sidebar>
    <SidebarContent>
      <SidebarGroup>
        <SidebarGroupLabel>File Zipper</SidebarGroupLabel>
        <SidebarGroupContent>
          <SidebarMenu>
            <SidebarMenuItem v-for="item in items" :key="item.title">
              <SidebarMenuButton as-child @click="navigate(item.url)">
                <div class="flex items-center gap-2 cursor-pointer">
                  <component :is="item.icon" />
                  <span>{{ item.title }}</span>
                </div>
              </SidebarMenuButton>
            </SidebarMenuItem>
          </SidebarMenu>
        </SidebarGroupContent>
      </SidebarGroup>
    </SidebarContent>
  </Sidebar>
</template>
