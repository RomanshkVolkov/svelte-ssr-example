<script lang="ts">
   import StatisticsCard from '@/components/cards/statistics-card.svelte';

   const data = [
      {
         title: 'Total Revenue',
         amount: '$1,000',
         trendingAmount: '+12.5%',
         firstMessage: 'Trending up this month ',
         secondMessage: 'Visitors for the last 6 months',
      },
      {
         title: 'New Customers',
         amount: '$500',
         trendingAmount: '-5.2%',
         firstMessage: 'Down 5.2% this period',
         secondMessage: 'Acquisition needs attention',
      },
      {
         title: 'Active accounts',
         amount: '8,529',
         trendingAmount: '+8.3%',
         firstMessage: 'Strong user retention',
         secondMessage: 'Engagement exceed targets',
      },
      {
         title: 'Growth Rate',
         amount: '$1,200',
         trendingAmount: '+15.0%',
         firstMessage: 'Steady performance increase',
         secondMessage: 'Meets growth projections',
      },
   ];
</script>

<div
   class="*:data-[slot=card]:from-primary/5 *:data-[slot=card]:to-card dark:*:data-[slot=card]:bg-card *:data-[slot=card]:shadow-xs @xl/main:grid-cols-2 @5xl/main:grid-cols-4 grid grid-cols-1 gap-4 *:data-[slot=card]:bg-gradient-to-t"
>
   {#each data as card}
      <StatisticsCard
         title={card.title}
         amount={card.amount}
         trendingAmount={card.trendingAmount}
         firstMessage={card.firstMessage}
         secondMessage={card.secondMessage}
      />
   {/each}
</div>
