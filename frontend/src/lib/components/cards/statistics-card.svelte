<script lang="ts">
   import { Badge } from '$lib/components/ui/badge/index.js';
   import Icon from '@iconify/svelte';
   import * as Card from '$lib/components/ui/card/index.js';
   import type { WithElementRef } from '@/utils';
   import type { HTMLAttributes } from 'svelte/elements';

   let {
      title,
      amount,
      trendingAmount,
      firstMessage,
      secondMessage,
      class: className,
      ...restProps
   }: WithElementRef<HTMLAttributes<HTMLDivElement>> & {
      title: string;
      amount: number | string;
      trendingAmount?: number | string;
      firstMessage?: string;
      secondMessage?: string;
   } = $props();

   const icon = String(trendingAmount).includes('+')
      ? 'fluent:arrow-trending-lines-20-filled'
      : 'fluent:arrow-trending-down-20-filled';
</script>

<Card.Root class="@container/card">
   <Card.Header>
      <Card.Description>{title}</Card.Description>
      <Card.Title
         class="@[250px]/card:text-3xl text-2xl font-semibold tabular-nums"
      >
         {amount}
      </Card.Title>
      <Card.Action>
         <Badge variant="outline">
            <Icon {icon} width="20" height="20" style="color: #0CB3B5" />
            {trendingAmount}
         </Badge>
      </Card.Action>
   </Card.Header>
   <Card.Footer class="flex-col items-start gap-1.5 text-sm">
      <div class="line-clamp-1 flex gap-2 font-medium">
         Trending up this month <Icon
            {icon}
            width="20"
            height="20"
            style="color: #0CB3B5"
         />
      </div>
      <div class="text-muted-foreground">Visitors for the last 6 months</div>
   </Card.Footer>
</Card.Root>
