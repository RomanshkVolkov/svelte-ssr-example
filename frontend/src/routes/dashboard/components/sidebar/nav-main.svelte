<script lang="ts">
    import * as m from '$lib/paraglide/messages'
    import Icon from '@iconify/svelte';
    import * as Sidebar from '$lib/components/ui/sidebar/index.js';
    import * as Collapsible from '@/components/ui/collapsible/index.js';


    let {
        items,
    }: {
        items: {
            title: string;
            url: string;
            // this should be `Component` after @lucide/svelte updates types
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            icon?: string;
            isActive?: boolean;
            items?: {
                title: string;
                url: string;
            }[];
        }[];
    } = $props();
</script>

<Sidebar.Group>
    <Sidebar.GroupLabel>
        {m.sidebar_platform()}
    </Sidebar.GroupLabel>

    <Sidebar.Menu>
        {#each items as item (item.title)}
            {#if !!item?.items?.length}
                <Collapsible.Root open={item.isActive} class="group/collapsible">
                    {#snippet child({props})}
                        <Sidebar.MenuItem {...props}>
                            <Collapsible.Trigger>
                                {#snippet child({props})}
                                    <Sidebar.MenuButton
                                            {...props}
                                            tooltipContent={item.title}
                                    >
                                        {#if item.icon}
                                            <Icon
                                                    icon={item.icon}
                                                    width="24"
                                                    height="24"
                                                    class="shrink-0"
                                            />
                                        {/if}
                                        <span>{m[item.title as keyof typeof m]?.()}</span>
                                        <Icon
                                                icon="formkit:right"
                                                width="24"
                                                height="24"
                                                class="ml-auto ransition-transform duration-200 group-data-[state=open]/collapsible:rotate-90"
                                        />
                                    </Sidebar.MenuButton>
                                {/snippet}
                            </Collapsible.Trigger>
                            <Collapsible.Content>
                                <Sidebar.MenuSub>
                                    {#each item.items ?? [] as subItem (subItem.title)}
                                        <Sidebar.MenuSubItem>
                                            <Sidebar.MenuSubButton>
                                                {#snippet child({props})}
                                                    <a
                                                            data-sveltekit-preload-data="tap"
                                                            href={subItem.url}
                                                            {...props}
                                                    >
                                                        <span>{m[subItem.title as keyof typeof m]?.()}</span>
                                                    </a>
                                                {/snippet}
                                            </Sidebar.MenuSubButton>
                                        </Sidebar.MenuSubItem>
                                    {/each}
                                </Sidebar.MenuSub>
                            </Collapsible.Content>
                        </Sidebar.MenuItem>
                    {/snippet}
                </Collapsible.Root>
            {:else}
                <Sidebar.MenuItem>
                    <Sidebar.MenuButton>
                        {#snippet child({props})}
                            <a
                                    data-sveltekit-preload-data="tap"
                                    href={item.url}
                                    {...props}
                            >
                                <span>{m[item.title as keyof typeof m]?.()}</span>
                            </a>
                        {/snippet}
                    </Sidebar.MenuButton>
                </Sidebar.MenuItem>
            {/if}
        {/each}
    </Sidebar.Menu>
</Sidebar.Group>
