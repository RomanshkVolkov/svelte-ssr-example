<script lang="ts">
    import {toast} from "svelte-sonner";
    import {superForm} from "sveltekit-superforms";
    import Button from "@/components/ui/button/button.svelte";
    import type {PageProps} from "./$types";

    let {data}: PageProps = $props();
    let isSubmitting = $state(false);

    const form = superForm(data.form, {
        onSubmit: () => {
            isSubmitting = true;
        },
        onUpdate: ({form, result}) => {
            isSubmitting = false;
            if (result.type === "failure") {
                toast.error("Internal server error");
                return;
            }

            toast.info(form.message);
        },
    });

    const {form: formData, enhance} = form;
</script>

<svelte:head>
    <title>Delete User | Guz-Studio Dashboard</title>
</svelte:head>

<form method="POST" use:enhance>
    <div class="flex h-svh justify-center">
        <input name="id" hidden type="text" defaultValue={$formData.id}/>
        <div class="flex flex-col items-center justify-center">
            <h1 class="text-2xl font-bold">
                Delete User "{data.user.name || "unknown"}"
            </h1>
            <p class="text-gray-500">
                Are you sure you want to delete this user?
            </p>
            <div class="w-full flex justify-center gap-4">
                <Button
                        class="btn btn-primary mt-4"
                        href="/dashboard/access/users"
                        isLoading={isSubmitting}
                >
                    Cancel
                </Button>

                <Button
                        class="btn btn-primary mt-4"
                        type="submit"
                        isLoading={isSubmitting}
                >
                    Confirm
                </Button>
            </div>
        </div>
    </div>
</form>
