<script lang="ts">
   import Icon from '@iconify/svelte';
   import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
   import * as Sidebar from '$lib/components/ui/sidebar/index.js';
   let { items }: { items: { name: string; url: string; icon: string }[] } =
      $props();
   const sidebar = Sidebar.useSidebar();
</script>

<Sidebar.Group class="group-data-[collapsible=icon]:hidden">
   <Sidebar.GroupLabel>Documents</Sidebar.GroupLabel>
   <Sidebar.Menu>
      {#each items as item (item.name)}
         <Sidebar.MenuItem>
            <Sidebar.MenuButton>
               {#snippet child({ props })}
                  <a {...props} href={item.url}>
                     <Icon icon={item.icon} width="24" height="24" />
                     <span>{item.name}</span>
                  </a>
               {/snippet}
            </Sidebar.MenuButton>
            <DropdownMenu.Root>
               <DropdownMenu.Trigger>
                  {#snippet child({ props })}
                     <Sidebar.MenuAction
                        {...props}
                        showOnHover
                        class="data-[state=open]:bg-accent rounded-sm"
                     >
                        <Icon
                           icon="material-symbols-light:home-max-dots-outline"
                           width="24"
                           height="24"
                           style="color: #0CB3B5"
                        />
                        <span class="sr-only">More</span>
                     </Sidebar.MenuAction>
                  {/snippet}
               </DropdownMenu.Trigger>
               <DropdownMenu.Content
                  class="w-24 rounded-lg"
                  side={sidebar.isMobile ? 'bottom' : 'right'}
                  align={sidebar.isMobile ? 'end' : 'start'}
               >
                  <DropdownMenu.Item>
                     <Icon
                        icon="material-symbols-light:folder-code"
                        width="24"
                        height="24"
                        style="color: #0CB3B5"
                     />
                     <span>Open</span>
                  </DropdownMenu.Item>
                  <DropdownMenu.Item>
                     <Icon
                        icon="material-symbols-light:inbox-text-share-outline-sharp"
                        width="24"
                        height="24"
                        style="color: #0CB3B5"
                     />
                     <span>Share</span>
                  </DropdownMenu.Item>
                  <DropdownMenu.Separator />
                  <DropdownMenu.Item variant="destructive">
                     <Icon
                        icon="material-symbols-light:delete-forever-outline"
                        width="24"
                        height="24"
                        style="color: #0CB3B5"
                     />
                     <span>Delete</span>
                  </DropdownMenu.Item>
               </DropdownMenu.Content>
            </DropdownMenu.Root>
         </Sidebar.MenuItem>
      {/each}
      <Sidebar.MenuItem>
         <Sidebar.MenuButton class="text-sidebar-foreground/70">
            <Icon
               icon="material-symbols-light:add"
               width="24"
               height="24"
               style="color: #0CB3B5"
            />
            <span>More</span>
         </Sidebar.MenuButton>
      </Sidebar.MenuItem>
   </Sidebar.Menu>
</Sidebar.Group>
