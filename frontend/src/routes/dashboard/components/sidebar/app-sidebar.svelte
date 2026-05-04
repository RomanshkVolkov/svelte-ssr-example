<script lang="ts" module>
   import AudioWaveformIcon from '@lucide/svelte/icons/audio-waveform';
   import CommandIcon from '@lucide/svelte/icons/command';
   import GalleryVerticalEndIcon from '@lucide/svelte/icons/gallery-vertical-end';
   import {ROUTES} from "@/constants";

   // This is sample data.
   const data = {
      user: {
         name: 'Sys admin',
         email: 'm@example.com',
         avatar: '/avatars/shadcn.jpg',
      },
      teams: [
         {
            name: 'Acme Inc',
            logo: GalleryVerticalEndIcon,
            plan: 'Enterprise',
         },
         {
            name: 'Acme Corp.',
            logo: AudioWaveformIcon,
            plan: 'Startup',
         },
         {
            name: 'Evil Corp.',
            logo: CommandIcon,
            plan: 'Free',
         },
      ],
      navMain: [
         {
            title: 'menu_dashboard',
            url: '/dashboard',
            icon: 'line-md:home',
         },
         {
            title: 'menu_access',
            url: '#',
            icon: 'line-md:account',
            isActive: true,
            items: [
               {
                  title: 'menu_access_user',
                  url: ROUTES.userList.path,
               },
               {
                  title: 'menu_access_roles',
                  url: ROUTES.roleList.path,
               },
            ],
         },
         {
            title: 'menu_management',
            url: '#',
            icon: 'line-md:folder-check',
            items: [
               {
                  title: 'menu_management_categories',
                  url: '#',
               },
               {
                  title: 'menu_management_products',
                  url: '#',
               },
            ],
         },
      ],
      projects: [
         {
            name: 'Design Engineering',
            url: '#',
            icon: 'line-md:home',
         },
         {
            name: 'Sales & Marketing',
            url: '#',
            icon: 'line-md:home',
         },
         {
            name: 'Travel',
            url: '#',
            icon: 'line-md:home',
         },
      ],
   };
</script>

<script lang="ts">
   import NavMain from './nav-main.svelte';
   import NavUser from '../sidebar/nav-user.svelte';
   import TeamSwitcher from './team-switcher.svelte';
   import * as Sidebar from '$lib/components/ui/sidebar/index.js';
   import { type ComponentProps } from 'svelte';
   import { browser } from '$app/environment';
   import Skeleton from '@/components/ui/skeleton/skeleton.svelte';
   let {
      ref = $bindable(null),
      collapsible = 'icon',
      ...restProps
   }: ComponentProps<typeof Sidebar.Root> = $props();
</script>

{#if browser}
   <Sidebar.Root {collapsible} {...restProps}>
      <Sidebar.Header>
         <TeamSwitcher teams={data.teams} />
      </Sidebar.Header>
      <Sidebar.Content>
         <NavMain items={data.navMain} />
      </Sidebar.Content>
      <Sidebar.Footer>
         <NavUser user={data.user} />
      </Sidebar.Footer>
      <Sidebar.Rail />
   </Sidebar.Root>
{:else}
   <div class="flex h-[100vh] w-60 flex-col justify-between p-4">
      <div class="flex flex-col gap-4">
         <Skeleton class="h-20" />
         <Skeleton class=" h-7" />
         <Skeleton class=" h-7" />
         <Skeleton class=" h-7" />
         <Skeleton class=" h-7" />
      </div>
      <div class="w-full">
         <Skeleton class="h-14" />
      </div>
   </div>
{/if}
