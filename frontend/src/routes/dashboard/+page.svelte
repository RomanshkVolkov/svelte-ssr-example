<script lang="ts">
   import ExampleTable from '@/components/tables/example/example-table.svelte';
   import ChartAreaInteractive from './components/chart-area-interactive.svelte';
   import SectionCards from './components/section-cards.svelte';
   import data from '@/components/tables/example/data';
   import * as m from '$lib/paraglide/messages'
   import { browser } from '$app/environment';
</script>

<svelte:head>
   <title>G Studios | Dashboard</title>
</svelte:head>

<div class="flex w-full flex-col gap-4">
   <h1 class="text-2xl font-bold mb-4">{m.menu_dashboard()}</h1>
   <SectionCards />
   <ChartAreaInteractive />
   <ExampleTable {data} />
</div>
