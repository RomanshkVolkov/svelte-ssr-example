<script lang="ts">
   import '../app.css';
   import { ModeWatcher } from 'mode-watcher';
   import { Toaster } from '$lib/components/ui/sonner/index';

   export const trailingSlash = 'always';
   export const prerender = true;

   const { children } = $props();
</script>

<ModeWatcher />
<Toaster />
{@render children()}
