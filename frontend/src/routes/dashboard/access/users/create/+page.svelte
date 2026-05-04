<script lang="ts">
    import { toast } from "svelte-sonner";
    import { superForm } from "sveltekit-superforms";
    import FormSection from "@/components/form/form-section.svelte";
    import CustomCombobox from "@/components/form/inputs/custom-combobox.svelte";
    import CustomInput from "@/components/form/inputs/custom-input.svelte";
    import Button from "@/components/ui/button/button.svelte";
    import FormFieldErrors from "@/components/ui/form/form-field-errors.svelte";
    import { ROUTES } from "@/constants";
    import { goto } from "$app/navigation";
    import * as Form from "$lib/components/ui/form/index.js";
    import type { PageData } from "./$types";
    import { browser } from "$app/environment";
    import { redirect } from "@sveltejs/kit";

    let { data }: { data: PageData } = $props();
    let isSubmitting = $state(false);

    const form = superForm(data.form, {
        onSubmit: () => {
            isSubmitting = true;
        },
        onUpdate: ({ form, result }) => {
            isSubmitting = false;
            if (result.type === "failure") {
                return;
            }

            toast.success(form.message);

            if (browser) window.location.href = ROUTES.usersList.path;
        },
    });

    const { form: formData, enhance } = form;
</script>

<svelte:head>
    <title>Create user</title>
</svelte:head>

<form method="POST" use:enhance>
    <div class="flex flex-col gap-4">
        <FormSection
            title="General data"
            description="Insert user general data"
            icon="material-symbols-light:account-circle"
        >
            <div class="grid gridl-cols-1 md:grid-cols-3 gap-4">
                <Form.Field {form} name="name">
                    <Form.Control>
                        {#snippet children({ props })}
                            <Form.Label>Name</Form.Label>
                            <CustomInput
                                {...props}
                                name="name"
                                wrapperClassName="w-full"
                                value={$formData.name}
                            />
                        {/snippet}
                    </Form.Control>
                    <FormFieldErrors />
                </Form.Field>

                <Form.Field {form} name="roleID">
                    <Form.Control>
                        {#snippet children({ props })}
                            <Form.Label>Role</Form.Label>
                            <CustomCombobox
                                name="roleID"
                                defaultValue={$formData.roleID}
                                options={data.roles}
                            />
                        {/snippet}
                    </Form.Control>
                    <FormFieldErrors />
                </Form.Field>
            </div>
        </FormSection>

        <FormSection
            title="Credentials"
            description="Insert your access credentials"
            icon="material-symbols-light:lock"
        >
            <div class="grid gridl-cols-1 md:grid-cols-3 gap-4">
                <Form.Field {form} name="email">
                    <Form.Control>
                        {#snippet children({ props })}
                            <Form.Label>Email</Form.Label>
                            <CustomInput
                                {...props}
                                name="email"
                                value={$formData.email}
                                type="email"
                                wrapperClassName="w-full"
                            />
                        {/snippet}
                    </Form.Control>
                    <Form.FieldErrors />
                </Form.Field>

                <Form.Field {form} name="password">
                    <Form.Control>
                        {#snippet children({ props })}
                            <Form.Label>Password</Form.Label>
                            <CustomInput
                                {...props}
                                name="password"
                                value={$formData.password}
                                type="password"
                                wrapperClassName="w-full"
                            />
                        {/snippet}
                    </Form.Control>
                    <Form.FieldErrors />
                </Form.Field>
            </div>
        </FormSection>

        <div class="flex w-full justify-end">
            <Button type="submit" variant="secondary" isLoading={isSubmitting}
                >Save data</Button
            >
        </div>
    </div>
</form>
