<script lang="ts">
    import Button from "@/components/ui/button/button.svelte";
    import type {PageProps} from "./$types";
    import RoleTable from "./components/role-table.svelte";
    import Icon from "@iconify/svelte";
    import {ROUTES} from "@/constants";

    let {data}: PageProps = $props();
</script>

<svelte:head>
    <title>Users</title>
</svelte:head>

<div class="flex w-full flex-col gap-4">
    <div class="flex w-full justify-between">
        <h1>Role list</h1>

        <a href={ROUTES.createRole.path}>
            <Button variant="outline">
                <Icon
                        icon="material-symbols-light:tab-new-right-outline-rounded"
                        class="size-5"
                />
                <span>Create new role</span>
            </Button>
        </a>
    </div>

    <RoleTable data={data.profiles}/>
</div>
