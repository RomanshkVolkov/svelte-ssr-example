<script lang="ts">
    import {toast} from "svelte-sonner";
    import {
        type Infer,
        type SuperValidated,
        superForm,
    } from "sveltekit-superforms";
    import FormSection from "@/components/form/form-section.svelte";
    import CustomCombobox from "@/components/form/inputs/custom-combobox.svelte";
    import CustomInput from "@/components/form/inputs/custom-input.svelte";
    import Button from "@/components/ui/button/button.svelte";
    import FormFieldErrors from "@/components/ui/form/form-field-errors.svelte";
    import {ROUTES} from "@/constants";
    import type {CommonOptionTypes} from "@/types/base-model";
    import {browser} from "$app/environment";
    import * as Form from "$lib/components/ui/form/index.js";
    import {EditUserSchema} from "./edit-user.schema";

    let {
        data,
    }: {
        data: {
            form: SuperValidated<Infer<typeof EditUserSchema>>;
            roles: CommonOptionTypes[];
        };
    } = $props();
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

            toast.success(form.message);
        },
    });

    const {form: formData, enhance, reset} = form;
</script>

<svelte:head>
    <title>Edit user</title>
</svelte:head>

<form method="POST" use:enhance>
    <div class="flex flex-col gap-4">
        <FormSection
                title="General data"
                description="Insert user general data"
                icon="material-symbols-light:account-circle"
        >
            <div class="grid gridl-cols-1 md:grid-cols-3 gap-4">
                <input hidden readonly name="id" value={$formData.id}/>
                <Form.Field {form} name="name">
                    <Form.Control>
                        {#snippet children({props})}
                            <Form.Label>Name</Form.Label>
                            <CustomInput
                                    {...props}
                                    name="name"
                                    wrapperClassName="w-full"
                                    value={$formData.name}
                            />
                        {/snippet}
                    </Form.Control>
                    <FormFieldErrors/>
                </Form.Field>

                <Form.Field {form} name="roleID">
                    <Form.Control>
                        {#snippet children({props})}
                            <Form.Label>Role</Form.Label>
                            <CustomCombobox
                                    name="roleID"
                                    defaultValue={$formData.roleID}
                                    options={data.roles}
                            />
                        {/snippet}
                    </Form.Control>
                    <FormFieldErrors/>
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
                        {#snippet children({props})}
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
                    <Form.FieldErrors/>
                </Form.Field>
            </div>
        </FormSection>

        <div class="flex w-full justify-end gap-4">
            <div class="flex justify-center gap-4">
                <Button href="/dashboard/access/users" variant="secondary">
                    Cancel
                </Button>
                <Button type="submit" isLoading={isSubmitting}
                >Save data
                </Button>
            </div>

        </div>
    </div>
</form>
