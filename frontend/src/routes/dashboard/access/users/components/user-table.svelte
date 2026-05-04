<script lang="ts">
    import type { ColumnDef } from "@tanstack/table-core";
    import { format } from "date-fns";
    import DataTableActions from "@/components/tables/data-table-actions.svelte";
    import DynamicTable from "@/components/tables/dynamic-table.svelte";
    import { renderComponent } from "@/components/ui/data-table";
    import { ROUTES } from "@/constants";
    import type { UserRow } from "../users.types";

    let { data }: { data: UserRow[] } = $props();

    const columns: ColumnDef<UserRow>[] = [
        {
            accessorKey: "createdAt",
            header: "Created At",
            cell: ({ row }) =>
                format(row.original.createdAt, "yyyy-MM-dd hh:mm"),
        },
        {
            accessorKey: "updatedAt",
            header: "Updated At",
            cell: ({ row }) =>
                format(row.original.updatedAt, "yyyy-MM-dd hh:mm"),
        },
        {
            accessorKey: "name",
            header: "Name",
            cell: ({ row }) => row.original.name,
        },
        {
            accessorKey: "email",
            header: "Email",
            cell: ({ row }) => row.original.email,
        },
        {
            accessorKey: "roleName",
            header: "Role",
            cell: ({ row }) => row.original.role.name,
        },
        {
            accessorKey: "actions",
            header: "Actions",
            cell: ({ row }) =>
                renderComponent(DataTableActions, {
                    editHref: ROUTES.editUser.path.replace(
                        "[id]",
                        row.original.id,
                    ),
                    deleteHref: ROUTES.deleteUser.path.replace(
                        "[id]",
                        row.original.id,
                    ),
                }),
        },
    ];
</script>

<div>
    <DynamicTable {data} {columns} />
</div>
