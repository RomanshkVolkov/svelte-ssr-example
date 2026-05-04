<script lang="ts">

    import {locales, setLocale} from '$lib/paraglide/runtime';
    import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
    import {Button} from "$lib/components/ui/button/index.js";
    import {Languages} from '@lucide/svelte'
    import {browser} from "$app/environment";

    const languages = {
        es: "Español",
        en: "English",
    }

    function handleChangeLocale(locale: (typeof locales)[number]) {
        setLocale(locale);
    }

</script>

<!-- is client use dropdown -->
{#if browser}
    <div>
        <DropdownMenu.Root>
            <DropdownMenu.Trigger>
                {#snippet child({props})}
                    <Button {...props} variant="outline">
                        <Languages/>
                    </Button>
                {/snippet}
            </DropdownMenu.Trigger>
            <DropdownMenu.Content class="w-56" align="start">
                {#each locales as locale}
                    <DropdownMenu.Item>
                        {#snippet child({props, action})}
                            <button
                                    {...props}
                                    on:click={() => handleChangeLocale(locale)}
                                    class="w-full flex justify-between text-left p-2 hover:bg-accent rounded"
                                    use:action
                            >
                                <span>{languages[locale]}</span>
                                <span>{locale}</span>
                            </button>
                        {/snippet}
                    </DropdownMenu.Item>
                {/each}
            </DropdownMenu.Content>
        </DropdownMenu.Root>
    </div>
    {:else}

    <Button variant="outline" isLoading>
        <Languages/>
    </Button>
{/if}