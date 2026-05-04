<script lang="ts">
   import * as Avatar from '$lib/components/ui/avatar/index.js';
   import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
   import * as Sidebar from '$lib/components/ui/sidebar/index.js';
   import Icon from '@iconify/svelte';

   let { user }: { user: { name: string; email: string; avatar: string } } =
      $props();
   const sidebar = Sidebar.useSidebar();
</script>

<Sidebar.Menu>
   <Sidebar.MenuItem>
      <DropdownMenu.Root>
         <DropdownMenu.Trigger>
            {#snippet child({ props })}
               <Sidebar.MenuButton
                  {...props}
                  size="lg"
                  class="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground"
               >
                  <Avatar.Root class="size-8 rounded-lg grayscale">
                     <Avatar.Image src={user.avatar} alt={user.name} />
                     <Avatar.Fallback class="rounded-lg">CN</Avatar.Fallback>
                  </Avatar.Root>
                  <div class="grid flex-1 text-left text-sm leading-tight">
                     <span class="truncate font-medium">{user.name}</span>
                     <span class="text-muted-foreground truncate text-xs">
                        {user.email}
                     </span>
                  </div>
                  <Icon
                     icon="mdi:dots-vertical"
                     width="24"
                     height="24"
                     style="color: #0CB3B5"
                  />
               </Sidebar.MenuButton>
            {/snippet}
         </DropdownMenu.Trigger>
         <DropdownMenu.Content
            class="w-(--bits-dropdown-menu-anchor-width) min-w-56 rounded-lg"
            side={sidebar.isMobile ? 'bottom' : 'right'}
            align="end"
            sideOffset={4}
         >
            <DropdownMenu.Label class="p-0 font-normal">
               <div
                  class="flex items-center gap-2 px-1 py-1.5 text-left text-sm"
               >
                  <Avatar.Root class="size-8 rounded-lg">
                     <Avatar.Image src={user.avatar} alt={user.name} />
                     <Avatar.Fallback class="rounded-lg">CN</Avatar.Fallback>
                  </Avatar.Root>
                  <div class="grid flex-1 text-left text-sm leading-tight">
                     <span class="truncate font-medium">{user.name}</span>
                     <span class="text-muted-foreground truncate text-xs">
                        {user.email}
                     </span>
                  </div>
               </div>
            </DropdownMenu.Label>
            <DropdownMenu.Separator />
            <DropdownMenu.Group>
               <DropdownMenu.Item>
                  <Icon
                     icon="material-symbols-light:account-circle-outline"
                     width="24"
                     height="24"
                     style="color: #0CB3B5"
                  />
                  Account
               </DropdownMenu.Item>
               <DropdownMenu.Item>
                  <Icon
                     icon="material-symbols-light:credit-score-outline"
                     width="24"
                     height="24"
                     style="color: #0CB3B5"
                  />
                  Billing
               </DropdownMenu.Item>
               <DropdownMenu.Item>
                  <Icon
                     icon="material-symbols-light:notifications-active-outline"
                     width="24"
                     height="24"
                     style="color: #0CB3B5"
                  />
                  Notifications
               </DropdownMenu.Item>
            </DropdownMenu.Group>
            <DropdownMenu.Separator />
            <DropdownMenu.Item>
               <Icon
                  icon="material-symbols-light:logout"
                  width="24"
                  height="24"
                  style="color: #0CB3B5"
               />
               Log out
            </DropdownMenu.Item>
         </DropdownMenu.Content>
      </DropdownMenu.Root>
   </Sidebar.MenuItem>
</Sidebar.Menu>
