<script lang="ts">
    import Button from "@/components/ui/button/button.svelte";
    import type { PageProps } from "./$types";
    import UserTable from "./components/user-table.svelte";
    import Icon from "@iconify/svelte";
    import { ROUTES } from "@/constants";

    let { data }: PageProps = $props();
</script>

<svelte:head>
    <title>Users</title>
</svelte:head>

<div class="flex w-full flex-col gap-4">
    <div class="flex w-full justify-between">
        <h1>User list</h1>

        <a href={ROUTES.createUser.path}>
            <Button variant="outline">
                <Icon
                    icon="material-symbols-light:tab-new-right-outline-rounded"
                    class="size-5"
                />
                <span> Create new user </span>
            </Button>
        </a>
    </div>

    <UserTable data={data.users} />
</div>
