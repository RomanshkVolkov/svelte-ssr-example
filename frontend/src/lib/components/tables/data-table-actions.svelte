<script lang="ts">
    import Icon from "@iconify/svelte";
    import { buttonVariants } from "$lib/components/ui/button/index.js";

    let { editHref, deleteHref }: { editHref?: string; deleteHref?: string } =
        $props();
</script>

<div class="flex w-full justify-center gap-2">
    {#if editHref}
        <a href={editHref} class={buttonVariants({ variant: "ghost" })}>
            <Icon icon="akar-icons:edit" class="size-6" />
        </a>
    {/if}
    {#if deleteHref}
        <a href={deleteHref} class={buttonVariants({ variant: "ghost" })}>
            <Icon icon="akar-icons:trash-can" class="size-6" />
        </a>
    {/if}
</div>
