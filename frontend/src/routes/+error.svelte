<!-- src/routes/+error.svelte -->
<script>
   import { page } from '$app/stores';
   import { goto } from '$app/navigation';
   import * as Card from '@/components/ui/card/index.js';
   import Button from '@/components/ui/button/button.svelte';

   // Función para reportar error
   async function reportError() {
      // Enviar error a tu sistema de logging
      console.log('Error reported:', $page.error);
   }

   // Función para buscar ayuda
   function getHelp() {
      window.open('https://support.example.com', '_blank');
   }
</script>

<div class="flex h-[100vh] w-vw items-center justify-center">
   <Card.Root>
      <Card.Header class="flex w-full justify-center gap-4">
         <div class="flex gap-4">
            <p class="text-5xl mb-4">⚠️</p>
            <h1 class="text-5xl">{$page.status || 404}</h1>
         </div>
      </Card.Header>

      <Card.Content>
         {#if $page.status === 404}
            <h2>Page Not Found</h2>
            <p>The page you're looking for doesn't exist.</p>
         {:else}
            <h2>Something went wrong</h2>
            <p>{$page.error?.message || 'An unexpected error occurred'}</p>
         {/if}
      </Card.Content>

      <Card.Footer>
         <div class="flex w-full justify-around">
            <Button class="btn primary" onclick={() => goto('/dashboard')}>
               Go Home</Button
            >

            <Button class="btn secondary" onclick={() => history.back()}>
               Go Back
            </Button>
         </div>

         {#if $page.status !== 404}
            <Button class="btn tertiary" onclick={reportError}>
               Report Error
            </Button>
         {/if}
      </Card.Footer>
   </Card.Root>
</div>
