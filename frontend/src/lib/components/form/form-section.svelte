<script lang="ts">
   import { cn } from '@/utils';
   import Icon from '@iconify/svelte';
   import type { Snippet } from 'svelte';

   let {
      title,
      description,
      icon,
      className,
      children,
   }: {
      title: string;
      description: string;
      icon: string;
      className?: string;
      children: Snippet;
   } = $props();
</script>

<div class={cn('space-y-6', className)}>
   <div class="flex items-start gap-4 pb-4 border-b">
      <div
         class="flex-shrink-0 w-10 h-10 bg-primary/10 rounded-lg flex items-center justify-center"
      >
         <Icon {icon} class="size-6 text-primary" />
      </div>
      <div class="flex-1 min-w-0">
         <h3 class="text-lg font-semibold text-foreground">{title}</h3>
         <p class="text-sm text-muted-foreground mt-1">{description}</p>
      </div>
   </div>
   <div class="space-y-4">{@render children()}</div>
</div>
