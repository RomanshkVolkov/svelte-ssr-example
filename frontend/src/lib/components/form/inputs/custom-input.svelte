<script lang="ts">
   import Button from '@/components/ui/button/button.svelte';
   import Input from '@/components/ui/input/input.svelte';
   import Label from '@/components/ui/label/label.svelte';
   import { cn } from '@/utils';
   import Icon from '@iconify/svelte';
   import type {
      HTMLInputAttributes,
      HTMLInputTypeAttribute,
   } from 'svelte/elements';

   type InputType = Exclude<HTMLInputTypeAttribute, 'file'>;
   const {
      wrapperClassName,
      ...inputProps
   }: {
      wrapperClassName?: string;
   } & Omit<HTMLInputAttributes, 'type'> &
      (
         | { type: 'file' | 'password'; files?: FileList }
         | { type?: InputType; files?: undefined }
      ) = $props();

   let { id, name, type, ...restProps } = inputProps;
   const isPasswordType = type === 'password';

   let inputRef: HTMLInputElement | null = null;
   let showPassword = $state(false);

   function toggleShowPassword() {
      showPassword = !showPassword;
   }
</script>

<div class={cn('relative flex w-fit flex-col gap-2', wrapperClassName)}>
   {#if isPasswordType}
      <Button
         variant="ghost"
         size="icon"
         class="absolute right-0 bottom-0 cursor-pointer"
         onclick={toggleShowPassword}
      >
         {#if showPassword}
            <Icon icon="ic:baseline-remove-red-eye" />
         {:else}
            <Icon icon="iconamoon:eye-off-light" width="24" height="24" />
         {/if}
      </Button>
   {/if}
   <Input
      {id}
      {name}
      {...restProps}
      type={isPasswordType && showPassword ? 'text' : (type as any)}
   />
</div>
