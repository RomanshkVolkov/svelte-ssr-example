<script lang="ts">
   import * as Sidebar from '$lib/components/ui/sidebar/index.js';
   import Skeleton from '@/components/ui/skeleton/skeleton.svelte';
   import AppSidebar from './components/sidebar/app-sidebar.svelte';
   import SiteHeader from './components/sidebar/site-header.svelte';

   let { children } = $props();
</script>

<Sidebar.Provider>
   <AppSidebar variant="inset" />
   <Sidebar.Inset>
      <SiteHeader />
      <div class="flex flex-1 flex-col">
         <div class="@container/main flex flex-1 flex-col gap-2">
            <div class="flex flex-col gap-4 p-4 md:gap-6 md:p-6">
               {@render children?.()}
            </div>
         </div>
      </div>
   </Sidebar.Inset>
</Sidebar.Provider>
