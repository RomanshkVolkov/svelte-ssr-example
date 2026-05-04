<script lang="ts">
    import {Button} from '$lib/components/ui/button/index.js';
    import * as Sidebar from '$lib/components/ui/sidebar/index.js';
    import {Separator} from '$lib/components/ui/separator/index.js';
    import ThemeToggle from '../../../components/theme-toggle.svelte';
    import LanguageSwitch from '../../../components/language-switch.svelte';
</script>

<header
        class="h-(--header-height) group-has-data-[collapsible=icon]/sidebar-wrapper:h-(--header-height) flex shrink-0 Eitems-center gap-2 border-b transition-[width,height] ease-linear"
>
    <div class="flex w-full items-center gap-1 px-4 lg:gap-2 lg:px-6">
        <Sidebar.Trigger class="-ml-1"/>
        <Separator
                orientation="vertical"
                class="mx-2 data-[orientation=vertical]:h-4"
        />
        <div class="ml-auto flex items-center gap-2">
            <div class="flex gap-2 py-1">
                <ThemeToggle/>
                <LanguageSwitch/>
            </div>
            <Button
                    href="https://github.com/shadcn-ui/ui/tree/main/apps/v4/app/(examples)/dashboard"
                    variant="ghost"
                    size="sm"
                    class="dark:text-foreground hidden sm:flex"
                    target="_blank"
                    rel="noopener noreferrer"
            >
                GitHub
            </Button>
        </div>
    </div>
</header>
