<script lang="ts">
    import type { ColumnDef } from "@tanstack/table-core";
    import { format } from "date-fns";
    import DataTableActions from "@/components/tables/data-table-actions.svelte";
    import DynamicTable from "@/components/tables/dynamic-table.svelte";
    import { renderComponent } from "@/components/ui/data-table";
    import { ROUTES } from "@/constants";
    import type { ProfileRow } from "../roles.types";

    let { data }: { data: ProfileRow[] } = $props();

    const columns: ColumnDef<ProfileRow>[] = [
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
            accessorKey: "actions",
            header: "Actions",
            cell: ({ row }) =>
                renderComponent(DataTableActions, {
                    editHref: ROUTES.editProfile.path.replace(
                        "[id]",
                        row.original.id,
                    ),
                    deleteHref: ROUTES.deleteProfile.path.replace(
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
